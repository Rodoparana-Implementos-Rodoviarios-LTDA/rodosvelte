package prenotas

import (
	"backend/utils" // Pacote utils com as funções genéricas
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// ===================
// VARIÁVEIS
// ===================

// preNotasCache armazena em cache os dados das pre_notas na memória.
// Isso evita a necessidade de buscar os dados do endpoint a cada requisição HTTP,
// melhorando a performance ao permitir o acesso rápido aos dados já processados.
// O cache é periodicamente atualizado pela função StartFetchingPreNotas.
var preNotasCache []PreNota

// mutex é usado para garantir exclusão mútua (mutex - mutual exclusion)
// no acesso ao preNotasCache. Como múltiplas goroutines podem tentar ler
// ou modificar o cache ao mesmo tempo, o mutex garante que apenas uma goroutine
// por vez possa acessar/modificar o cache, evitando condições de corrida (race conditions).
var mutex sync.Mutex

// ===================
// ESTRUTURAS
// ===================

// Headers finais para a tabela
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

// Headers originais da requisição
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

// ===================
// REQUISIÇÃO HTTP
// ===================

// Função para iniciar a busca periódica das pre_notas
func StartFetchingPreNotas() {
	go func() {
		for {
			log.Println("Buscando pre_notas...")
			preNotas, err := utils.FetchFromEndpoint[RawPreNota]("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=99999", nil)
			if err != nil {
				log.Printf("Erro ao buscar pre_notas: %v", err)
			} else {
				processedPreNotas := processPreNotas(preNotas)
				mutex.Lock()
				preNotasCache = processedPreNotas // Atualiza o cache
				mutex.Unlock()
				log.Println("Importação de pre_notas concluída com sucesso!")
			}
			time.Sleep(1 * time.Minute) // Intervalo de 1 minuto
		}
	}()
}

func GetPreNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pega os parâmetros de paginação e ordenação
	page, pageSize := utils.GetPaginationParams(r)
	filters := utils.GetFilters(r)
	sortBy, sortOrder := utils.GetSortingParams(r, map[string]bool{
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
	})

	mutex.Lock()
	defer mutex.Unlock()

	if len(preNotasCache) == 0 {
		http.Error(w, "Nenhuma pre_nota disponível.", http.StatusNotFound)
		return
	}

	// Chama as funções auxiliares para filtrar, ordenar e paginar
	filteredPreNotas := filterPreNotas(preNotasCache, filters)
	sortedPreNotas := sortPreNotas(filteredPreNotas, sortBy, sortOrder)
	paginatedPreNotas := paginatePreNotas(sortedPreNotas, page, pageSize)

	// Retorna os dados paginados
	json.NewEncoder(w).Encode(paginatedPreNotas)
}

// Função para processar e ajustar os dados das pre_notas
func processPreNotas(rawPreNotas []RawPreNota) []PreNota {
	var processed []PreNota
	for _, raw := range rawPreNotas {
		status := utils.DetermineStatus(raw.Status, raw.Rev)

		preNota := PreNota{
			Filial:     utils.TrimString(raw.Filial),
			NF:         utils.TrimString(raw.Doc) + " - " + utils.TrimString(raw.Serie),
			Status:     status,
			Fornecedor: utils.TrimString(raw.Fornece),
			Emissao:    utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
			Inclusao:   utils.FormatDate(utils.TrimString(raw.Inclusao), "20060102", "02/01/2006"),
			Vencimento: utils.FormatDate(utils.TrimString(raw.Vencimento), "20060102", "02/01/2006"),
			Valor:      utils.FormatCurrency(utils.TrimString(raw.ValBrut)),
			Tipo:       utils.TrimString(raw.Tipo),
			Prioridade: utils.TrimString(raw.Prior),
			Usuario:    utils.TrimString(raw.UsrRa),
			Obs:        utils.TrimString(raw.Obs),
			ObsRev:     utils.TrimString(raw.ObsRev),
			Revisao:    utils.TrimString(raw.Rev),
			Rec:        utils.TrimString(raw.Rec),
		}
		processed = append(processed, preNota)
	}
	return processed
}

// ===================
// PARÂMETROS DA REQUISIÇÃO
// ===================

// Função para filtrar pre_notas
func filterPreNotas(preNotas []PreNota, filters map[string]string) []PreNota {
	return utils.FilterData(preNotas, filters, func(item PreNota, key, value string) bool {
		switch key {
		case "Filial":
			return item.Filial == value
		case "Status":
			return item.Status == value
		case "Fornecedor":
			return item.Fornecedor == value
		case "Emissao":
			return item.Emissao == value
		case "Inclusao":
			return item.Inclusao == value
		case "Vencimento":
			return item.Vencimento == value
		case "Valor":
			return item.Valor == value
		case "Tipo":
			return item.Tipo == value
		case "Prioridade":
			return item.Prioridade == value
		case "Usuario":
			return item.Usuario == value
		case "Obs":
			return item.Obs == value
		default:
			return true
		}
	})
}

// Função para ordenar pre_notas
func sortPreNotas(preNotas []PreNota, sortBy, sortOrder string) []PreNota {
	return utils.SortData(preNotas, sortBy, sortOrder, func(i, j PreNota) bool {
		switch sortBy {
		case "Filial":
			return i.Filial < j.Filial
		case "NF":
			return i.NF < j.NF
		case "Fornecedor":
			return i.Fornecedor < j.Fornecedor
		case "Emissao":
			return i.Emissao < j.Emissao
		case "Inclusao":
			return i.Inclusao < j.Inclusao
		case "Vencimento":
			return i.Vencimento < j.Vencimento
		case "Valor":
			return i.Valor < j.Valor
		case "Tipo":
			return i.Tipo < j.Tipo
		case "Prioridade":
			return i.Prioridade < j.Prioridade
		case "Usuario":
			return i.Usuario < j.Usuario
		default:
			return true
		}
	})
}

// Função para paginar pre_notas
func paginatePreNotas(preNotas []PreNota, page, pageSize int) []PreNota {
	return utils.Paginate(preNotas, page, pageSize)
}