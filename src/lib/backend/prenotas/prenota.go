package prenotas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// PreNota e RawPreNota são as estruturas que serão usadas para representar	 os dados
type PreNota struct {
	Filial     string `json:"Filial"`
	NF         string `json:"NF"`
	Status     string `json:"Status"`
	Fornecedor string `json:"Fornecedor"`
	Emissao    string `json:"Emissao"`
	Inclusao   string `json:"Inclusao"`
	Vencimento string `json:"Vencimento"`
	Valor      string `json:"Valor"`
	Tipo       string `json:"Tipo"`
	Prioridade string `json:"Prioridade"`
	Usuario    string `json:"Usuario"`
	Obs        string `json:"Obs"`
	ObsRev     string `json:"ObsRev"`
	Revisao    string `json:"Revisao"`
	Rec        string `json:"Rec"`
}

type RawPreNota struct {
	Filial     string `json:"F1_FILIAL"`
	Doc        string `json:"F1_DOC"`
	Serie      string `json:"F1_SERIE"`
	Status     string `json:"F1_STATUS"`
	Fornece    string `json:"FORNECE"`
	Emissao    string `json:"F1_EMISSAO"`
	Inclusao   string `json:"F1_DTDIGIT"`
	Vencimento string `json:"VENCIMENTO"`
	ValBrut    string `json:"F1_VALBRUT"`
	Tipo       string `json:"F1_XTIPO"`
	Prior      string `json:"F1_XPRIOR"`
	UsrRa      string `json:"F1_XUSRRA"`
	Obs        string `json:"F1_XOBS"`
	ObsRev     string `json:"F1_ZOBSREV"`
	Rev        string `json:"F1_XREV"`
	Rec        string `json:"REC"`
}

var preNotasCache []PreNota
var mutex sync.Mutex

// Função para iniciar a busca periódica das pre_notas
func StartFetchingPreNotas() {
	go func() {
		for {
			log.Println("Buscando pre_notas...")
			preNotas, err := fetchPreNotas()
			if err != nil {
				log.Printf("Erro ao buscar pre_notas: %v", err)
			} else {
				processedPreNotas := processPreNotas(preNotas)
				mutex.Lock()
				preNotasCache = processedPreNotas // Atualiza o cache
				mutex.Unlock()
				log.Println("Importação de pre_notas concluída com sucesso!")
			}
			time.Sleep(1 * time.Minute) // Intervalo de 10 minutos
		}
	}()
}

// Função para buscar pre_notas de um endpoint externo (como o Protheus)
func fetchPreNotas() ([]RawPreNota, error) {
	url := "http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=99999"
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada: %d %s", resp.StatusCode, resp.Status)
	}

	var rawPreNotas []RawPreNota
	err = json.NewDecoder(resp.Body).Decode(&rawPreNotas)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer o parsing do JSON: %v", err)
	}

	return rawPreNotas, nil
}

// Função para processar e ajustar os dados das pre_notas
func processPreNotas(rawPreNotas []RawPreNota) []PreNota {
	var processed []PreNota
	for _, raw := range rawPreNotas {
		// Verifica o status com base em Revisao e Status
		status := determineStatus(strings.TrimSpace(raw.Status), strings.TrimSpace(raw.Rev))

		preNota := PreNota{
			Filial:     strings.TrimSpace(raw.Filial),
			NF:         strings.TrimSpace(raw.Doc) + " - " + strings.TrimSpace(raw.Serie),
			Status:     status,
			Fornecedor: strings.TrimSpace(raw.Fornece),
			Emissao:    formatDate(strings.TrimSpace(raw.Emissao)),
			Inclusao:   formatDate(strings.TrimSpace(raw.Inclusao)),
			Vencimento: formatDate(strings.TrimSpace(raw.Vencimento)),
			Valor:      formatCurrency(strings.TrimSpace(raw.ValBrut)),
			Tipo:       strings.TrimSpace(raw.Tipo),
			Prioridade: strings.TrimSpace(raw.Prior),
			Usuario:    strings.TrimSpace(raw.UsrRa),
			Obs:        strings.TrimSpace(raw.Obs),
			ObsRev:     strings.TrimSpace(raw.ObsRev),
			Revisao:    strings.TrimSpace(raw.Rev),
			Rec:        strings.TrimSpace(raw.Rec),
		}
		processed = append(processed, preNota)
	}
	return processed
}

// Função para determinar o status da pre_nota com base em Revisao e Status
func determineStatus(status, revisao string) string {
	if strings.TrimSpace(revisao) != "" {
		return "Revisar"
	}
	if strings.TrimSpace(status) != "" {
		return "Classificado"
	}
	return "Pendente"
}

// Função para formatar a data no formato DDMMYYYY para DD/MM/YYYY
func formatDate(dateStr string) string {
	if len(dateStr) == 8 {
		return dateStr[6:8] + "-" + dateStr[4:6] + "-" + dateStr[0:4]
	}
	return dateStr
}

// Função para formatar valores monetários (Real)
func formatCurrency(valueStr string) string {
	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		log.Printf("Erro ao formatar valor: %v", err)
		return valueStr
	}
	return fmt.Sprintf("%.2f", value)
}

// Handler HTTP para buscar pre_notas com paginação, filtro e ordenação
func GetPreNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pegar os parâmetros de paginação
	page, pageSize := getPaginationParams(r)

	// Pegar parâmetros de filtro e ordenação
	filters := getFilters(r)
	sortBy, sortOrder := getSortingParams(r)

	mutex.Lock()
	defer mutex.Unlock()

	if len(preNotasCache) == 0 {
		http.Error(w, "Nenhuma pre_nota disponível.", http.StatusNotFound)
		return
	}

	// Filtrar e ordenar os dados
	filteredPreNotas := filterPreNotas(preNotasCache, filters)
	sortedPreNotas := sortPreNotas(filteredPreNotas, sortBy, sortOrder)

	// Aplicar paginação
	paginatedPreNotas := paginate(sortedPreNotas, page, pageSize)

	json.NewEncoder(w).Encode(paginatedPreNotas)
}

// Função para pegar parâmetros de paginação
func getPaginationParams(r *http.Request) (int, int) {
	page := 1
	pageSize := 10

	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("pageSize")

	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}
	if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 {
		pageSize = ps
	}

	return page, pageSize
}

// Função para pegar parâmetros de filtro da URL e aplicar dinamicamente
func getFilters(r *http.Request) map[string]string {
	filters := make(map[string]string)

	// Captura todos os parâmetros de query
	for key, values := range r.URL.Query() {
		filters[key] = values[0] // Pega apenas o primeiro valor
	}

	return filters
}

// Função para pegar parâmetros de ordenação
func getSortingParams(r *http.Request) (string, string) {
	// Lista de campos válidos para ordenação (você pode adicionar mais conforme necessário)
	validSortFields := map[string]bool{
		"Filial":     true,
		"NF":         true,
		"Fornecedor": true,
		"Emissao":    true,
		"Inclusao":   true,
		"Vencimento": true,
		"Valor":      true,
		"Tipo":       true,
		"Prioridade": true,
		"Status":     true,
		"Usuario":    true,
	}

	// Pega o parâmetro de ordenação do request
	sortBy := r.URL.Query().Get("sortBy")
	if _, valid := validSortFields[sortBy]; !valid {
		sortBy = "Filial" // Campo padrão se o campo enviado não for válido
	}

	// Pega a direção de ordenação, que pode ser 'asc' ou 'desc'
	sortOrder := r.URL.Query().Get("sortOrder")
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc" // Ordenação padrão
	}

	return sortBy, sortOrder
}

// Função para aplicar filtros dinâmicos, incluindo faixas de datas para Emissao, Inclusao e Vencimento
func filterPreNotas(preNotas []PreNota, filters map[string]string) []PreNota {
	var filtered []PreNota

	// Funções auxiliares para fazer parsing de data
	parseDate := func(dateStr string) (time.Time, error) {
		return time.Parse("2006-01-02", dateStr) // Formato de data YYYY-MM-DD
	}

	checkDateRange := func(date time.Time, startDate, endDate time.Time) bool {
		if !startDate.IsZero() && date.Before(startDate) {
			return false
		}
		if !endDate.IsZero() && date.After(endDate) {
			return false
		}
		return true
	}

	// Converte as faixas de data para as três colunas
	var emissaoStart, emissaoEnd, inclusaoStart, inclusaoEnd, vencimentoStart, vencimentoEnd time.Time

	if start, ok := filters["emissaoStart"]; ok {
		emissaoStart, _ = parseDate(start) // Ignora erro para simplificar
	}
	if end, ok := filters["emissaoEnd"]; ok {
		emissaoEnd, _ = parseDate(end) // Ignora erro para simplificar
	}

	if start, ok := filters["inclusaoStart"]; ok {
		inclusaoStart, _ = parseDate(start)
	}
	if end, ok := filters["inclusaoEnd"]; ok {
		inclusaoEnd, _ = parseDate(end)
	}

	if start, ok := filters["vencimentoStart"]; ok {
		vencimentoStart, _ = parseDate(start)
	}
	if end, ok := filters["vencimentoEnd"]; ok {
		vencimentoEnd, _ = parseDate(end)
	}

	for _, preNota := range preNotas {
		match := true

		// Verifica se cada filtro corresponde ao valor correto da pre_nota
		for key, value := range filters {
			switch key {
			case "Filial":
				if preNota.Filial != value {
					match = false
				}
			case "Status":
				if preNota.Status != value {
					match = false
				}
			case "Fornecedor":
				if preNota.Fornecedor != value {
					match = false
				}
			case "Emissao":
				// Converte a data de Emissao da pre_nota
				emissaoDate, err := time.Parse("02-01-2006", preNota.Emissao)
				if err != nil {
					log.Printf("Erro ao parsear data de Emissao: %v", err)
					match = false
				} else {
					// Verifica se a data está dentro da faixa
					if !checkDateRange(emissaoDate, emissaoStart, emissaoEnd) {
						match = false
					}
				}
			case "Inclusao":
				// Converte a data de Inclusao da pre_nota
				inclusaoDate, err := time.Parse("02-01-2006", preNota.Inclusao)
				if err != nil {
					log.Printf("Erro ao parsear data de Inclusao: %v", err)
					match = false
				} else {
					// Verifica se a data está dentro da faixa
					if !checkDateRange(inclusaoDate, inclusaoStart, inclusaoEnd) {
						match = false
					}
				}
			case "Vencimento":
				// Converte a data de Vencimento da pre_nota
				vencimentoDate, err := time.Parse("02-01-2006", preNota.Vencimento)
				if err != nil {
					log.Printf("Erro ao parsear data de Vencimento: %v", err)
					match = false
				} else {
					// Verifica se a data está dentro da faixa
					if !checkDateRange(vencimentoDate, vencimentoStart, vencimentoEnd) {
						match = false
					}
				}
			case "Valor":
				if preNota.Valor != value {
					match = false
				}
			case "Tipo":
				if preNota.Tipo != value {
					match = false
				}
			case "Prioridade":
				if preNota.Prioridade != value {
					match = false
				}
			case "Usuario":
				if preNota.Usuario != value {
					match = false
				}
			case "Obs":
				if preNota.Obs != value {
					match = false
				}
			// Adicione mais colunas conforme necessário
			default:
				// Coluna desconhecida, ignora
			}
		}

		// Se todos os filtros corresponderem, adiciona a pre_nota ao resultado
		if match {
			filtered = append(filtered, preNota)
		}
	}

	return filtered
}

// Função para aplicar a ordenação nas pre_notas
func sortPreNotas(preNotas []PreNota, sortBy, sortOrder string) []PreNota {
	sort.SliceStable(preNotas, func(i, j int) bool {
		switch sortBy {
		case "Filial":
			if sortOrder == "asc" {
				return preNotas[i].Filial < preNotas[j].Filial
			}
			return preNotas[i].Filial > preNotas[j].Filial
		case "NF":
			if sortOrder == "asc" {
				return preNotas[i].NF < preNotas[j].NF
			}
			return preNotas[i].NF > preNotas[j].NF
		default:
			return true
		}
	})
	return preNotas
}

// Função para aplicar paginação nas pre_notas
func paginate(preNotas []PreNota, page, pageSize int) []PreNota {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(preNotas) {
		return []PreNota{}
	}
	if end > len(preNotas) {
		end = len(preNotas)
	}

	return preNotas[start:end]
}
