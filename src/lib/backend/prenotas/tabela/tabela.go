package tabela

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// ===================
// VARIÁVEIS GLOBAIS
// ===================

// preNotasCache mantém em memória os dados das pre_notas, evitando buscas frequentes no endpoint.
var preNotasCache []PreNota
var mutex sync.Mutex // Controla o acesso concorrente ao cache.

// ===================
// ESTRUTURAS DE DADOS
// ===================

// PreNota representa a estrutura final da nota.
type PreNota struct {
	Filial, NF, Status, Fornecedor, Emissao, Inclusao, Vencimento, Valor,
	Tipo, Prioridade, Usuario, Obs, ObsRev, Revisao, Rec string
}

// RawPreNota mapeia os dados originais do endpoint.
type RawPreNota struct {
	Filial     string `json:"F1_FILIAL"`
	Doc        string `json:"F1_DOC"`
	Serie      string `json:"F1_SERIE"`
	Status     string `json:"F1_STATUS"`
	Fornecedor string `json:"FORNECE"`
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
// REQUISIÇÃO E CACHE
// ===================

// StartFetchingPreNotas inicia a atualização periódica do cache de pre_notas.
func StartFetchingPreNotas() {
	go func() {
		for {
			log.Println("Buscando pre_notas...")
			preNotas, err := utils.FetchFromEndpoint[RawPreNota]("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=99", nil)
			if err != nil {
				log.Printf("Erro ao buscar pre_notas: %v", err)
				continue
			}

			mutex.Lock()
			preNotasCache = processPreNotas(preNotas)
			mutex.Unlock()
			log.Println("Cache de pre_notas atualizado com sucesso!")
			time.Sleep(1 * time.Minute) // Atualiza o cache a cada minuto
		}
	}()
}

// ===================
// HANDLER PRINCIPAL
// ===================

// GetPreNotas responde com os dados de pre_notas filtrados, classificados e paginados.
func GetPreNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	// Verifica se há dados no cache
	if len(preNotasCache) == 0 {
		http.Error(w, "Nenhuma pre_nota disponível.", http.StatusNotFound)
		return
	}

	// Aplicar Filtros
	filterableColumns := []string{"Filial", "NF", "Fornecedor", "Emissao", "Inclusao", "Vencimento", "Status"}
	filteredPreNotas := applyFilters(preNotasCache, r, filterableColumns)

	// Aplicar Classificação
	sortedPreNotas := applySorting(filteredPreNotas, r)

	// Aplicar Paginação
	paginatedPreNotas := utils.Paginate(sortedPreNotas, r)

	// Retornar os dados filtrados, classificados e paginados
	json.NewEncoder(w).Encode(paginatedPreNotas)
}

// ===================
// FUNÇÕES AUXILIARES
// ===================

// Função para aplicar os filtros nos dados com base nos headers.
func applyFilters(preNotas []PreNota, r *http.Request, filterableColumns []string) []PreNota {
	// Chama a função de filtro genérica do utils
	return utils.FilterData(preNotas, r, filterableColumns)
}

// Função para aplicar a classificação com base nos headers.
func applySorting(preNotas []PreNota, r *http.Request) []PreNota {
	// Headers que determinam a coluna e ordem de classificação
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	// Define valores padrão se os headers não forem fornecidos
	if sortBy == "" {
		sortBy = "Inclusao" // Valor padrão
	}
	if sortOrder == "" {
		sortOrder = "desc" // Valor padrão
	}

	// Chama a função de classificação genérica do utils
	return utils.SortByColumn(preNotas, sortBy, sortOrder)
}

// ===================
// PROCESSAMENTO
// ===================

// processPreNotas transforma os dados brutos em pre_notas processadas.
func processPreNotas(rawPreNotas []RawPreNota) []PreNota {
	var processed []PreNota
	for _, raw := range rawPreNotas {
		// Concatena NF e Série, mas garante que não tenha valores " - " se os dois campos estiverem vazios
		nf := utils.TrimString(raw.Doc)
		serie := utils.TrimString(raw.Serie)
		nfCompleta := nf
		if serie != "" {
			nfCompleta += " - " + serie
		}

		// Verifica e formata as datas
		emissao := utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006")
		inclusao := utils.FormatDate(utils.TrimString(raw.Inclusao), "20060102", "02/01/2006")
		vencimento := utils.FormatDate(utils.TrimString(raw.Vencimento), "20060102", "02/01/2006")

		preNota := PreNota{
			Filial:     utils.TrimString(raw.Filial),
			NF:         nfCompleta,
			Status:     utils.DetermineStatus(raw.Status, raw.Rev),
			Fornecedor: utils.TrimString(raw.Fornecedor),
			Emissao:    emissao,
			Inclusao:   inclusao,
			Vencimento: vencimento,
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
