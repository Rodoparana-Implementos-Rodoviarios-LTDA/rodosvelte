package portaria

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)


// ===================
// VARIAVEIS GLOBAIS
// ===================
var (
	movimentosPortariaCache []MovimentoPortaria
	cacheTimestamp          time.Time
	cacheDuration           = 1 * time.Hour
	mutex                   sync.Mutex
	cacheInitialized        bool
	cacheUpdating           bool
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type MovimentoPortaria struct {
	Filial, NF, Cliente, Produto, TipoMov, DataHora, Responsavel, Placa, Observacao string
	Saldo, Rec                                                                      int
}

type RawMovimentoPortaria struct {
	Filial        string `json:"Z08_FILIAL"`
	Origem        string `json:"Z08_ORIGEM"`
	Documento     string `json:"Z08_DOC"`
	Serie         string `json:"Z08_SERIE"`
	Cliente       string `json:"Z08_CLIFOR"`
	Loja          string `json:"Z08_LOJA"`
	ClienteNome   string `json:"A1_NOME"`
	Codigo        string `json:"Z08_COD"`
	Item          string `json:"Z08_ITEM"`
	Descricao     string `json:"B1_DESC"`
	TipoMovimento string `json:"Z08_TIPMOV"`
	Data          string `json:"Z08_DATA"`
	Hora          string `json:"Z08_HORA"`
	Responsavel   string `json:"Z08_RESPON"`
	Placa         string `json:"Z08_PLACA"`
	Observacao    string `json:"Z08_OBSERV"`
	Saldo         int    `json:"SALDO"`
	Seq           int    `json:"Z08"`
}

// ===================
// REQUISICAO E CACHE
// ===================

func InitializeCache() {
	log.Println("Carregando dados iniciais de portaria...")
	fetchMovimentosPortariaData()
}

func fetchMovimentosPortariaData() {
	mutex.Lock()

	if cacheUpdating {
		log.Println("Atualizacao do cache de portaria ja esta em andamento, aguardando...")
		mutex.Unlock()
		return
	}

	cacheUpdating = true
	mutex.Unlock()

	movimentos, err := utils.FetchFromEndpoint[RawMovimentoPortaria]("http://protheus-vm:9010/rest/MovPortaria/MovimentosPorNF", nil)
	if err != nil {
		log.Printf("Erro ao buscar movimentos da portaria: %v", err)
		mutex.Lock()
		cacheUpdating = false
		mutex.Unlock()
		return
	}

	mutex.Lock()
	movimentosPortariaCache = processMovimentosPortaria(movimentos)
	cacheTimestamp = time.Now()
	cacheInitialized = true
	cacheUpdating = false
	mutex.Unlock()

	log.Println("Cache de movimentos da portaria atualizado com sucesso!")
}

// ===================
// HANDLER HTTP
// ===================

func GetMovimentosPortaria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()

	if !cacheInitialized || len(movimentosPortariaCache) == 0 || time.Since(cacheTimestamp) > cacheDuration {
		log.Println("Cache de portaria expirado ou vazio. Carregando novos dados em background...")

		go fetchMovimentosPortariaData()
	}

	filterableColumns := []string{"Filial", "NF", "Cliente", "Produto", "DataHora", "TipoMov"}
	filteredMovimentos := applyFilters(movimentosPortariaCache, r, filterableColumns)

	if len(filteredMovimentos) == 0 {
		log.Println("Dados nao encontrados no cache, buscando no webservice...")
		mutex.Unlock()
		fetchMovimentosPortariaData()
		mutex.Lock()
		filteredMovimentos = applyFilters(movimentosPortariaCache, r, filterableColumns)
		if len(filteredMovimentos) == 0 {
			mutex.Unlock()
			http.Error(w, "Nenhum movimento disponivel.", http.StatusNotFound)
			return
		}
	}

	sortedMovimentos := applySorting(filteredMovimentos, r)
	paginatedMovimentos := utils.Paginate(sortedMovimentos, r)
	mutex.Unlock()
	if err := json.NewEncoder(w).Encode(paginatedMovimentos); err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

func applyFilters(movimentos []MovimentoPortaria, r *http.Request, filterableColumns []string) []MovimentoPortaria {
	return utils.FilterData(movimentos, r, filterableColumns)
}

func applySorting(movimentos []MovimentoPortaria, r *http.Request) []MovimentoPortaria {

	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	if sortBy == "" {
		sortBy = "DataHora"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	return utils.SortByColumn(movimentos, sortBy, sortOrder)
}

// ===================
// PROCESSAMENTO DOS DADOS
// ===================

func processMovimentosPortaria(rawMovimentos []RawMovimentoPortaria) []MovimentoPortaria {
	var processed []MovimentoPortaria
	for _, raw := range rawMovimentos {
		movimento := MovimentoPortaria{
			Filial:      utils.TrimString(raw.Filial),
			NF:          utils.TrimString(raw.Documento) + " - " + utils.TrimString(raw.Serie),
			Cliente:     utils.TrimString(raw.Cliente) + " - " + utils.TrimString(raw.Loja) + " - " + utils.TrimString(raw.ClienteNome),
			Produto:     utils.TrimString(raw.Codigo) + " - " + utils.TrimString(raw.Item) + " - " + utils.TrimString(raw.Descricao),
			TipoMov:     utils.MapTipoMovimento(raw.TipoMovimento),
			DataHora:    utils.FormatDate(utils.TrimString(raw.Data), "20060102", "02/01/2006") + " " + utils.TrimString(raw.Hora),
			Responsavel: utils.TrimString(raw.Responsavel),
			Placa:       utils.TrimString(raw.Placa),
			Observacao:  utils.TrimString(raw.Observacao),
			Rec:         raw.Rec,
			Saldo:       raw.Saldo,
		}
		processed = append(processed, movimento)
	}
	return processed
}
