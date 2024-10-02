package historico

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// ===================
// VARIAVEIS GLOBAIS
// ===================
var (
	historicoCache   []Historico     // Cache dos dados de historico
	cacheTimestamp   time.Time       // Timestamp do cache
	cacheDuration    = 1 * time.Hour // Duracao do cache ajustada para 1 hora
	mutex            sync.Mutex      // Controle de concorrencia para o cache
	cacheInitialized bool            // Flag para indicar se o cache foi inicializado
	cacheUpdating    bool            // Flag para indicar se o cache esta sendo atualizado
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type Historico struct {
	Filial, NF, Vendedor, Cliente, Produto, DataHora, Responsavel, Placa, Observacao string
	Saldo, Rec                                                                       int
}

type RawHistorico struct {
	Filial      string `json:"Z08_FILIAL"`
	Origem      string `json:"Z08_ORIGEM"`
	Doc         string `json:"Z08_DOC"`
	Serie       string `json:"Z08_SERIE"`
	CodCli      string `json:"Z08_CLIFOR"`
	LojaCli     string `json:"Z08_LOJA"`
	NomeCliente string `json:"A1_NOME"`
	CodProd     string `json:"Z08_COD"`
	ItemProd    string `json:"Z08_ITEM"`
	NomeProd    string `json:"B1_DESC"`
	TipMov      string `json:"Z08_TIPMOV"`
	Data        string `json:"Z08_DATA"`
	Hora        string `json:"Z08_HORA"`
	RetirPor    string `json:"Z08_RETPOR"`
	Respons     string `json:"Z08_RESPON"`
	Placa       string `json:"Z08_PLACA"`
	Obs         string `json:"Z08_OBSERV"`
	CodUsuario  string `json:"Z08_CODUSR"`
	Conferido   string `json:"Z08_CONFER"`
	DataConf    string `json:"Z08_DTCONF"`
	HoraConf    string `json:"Z08_HRCONF"`
	Rec         int    `json:"Z08R_E_C_N_O_"`
	Saldo       int    `json:"SALDO"`
}

// ===================
// REQUISICAO E CACHE
// ===================

func InitializeCache() {
	log.Println("Carregando dados iniciais de historico...")
	fetchHistoricoData()
}

func fetchHistoricoData() {
	mutex.Lock()

	// Verifica se o cache ja esta sendo atualizado
	if cacheUpdating {
		log.Println("Atualizacao do cache de historico ja esta em andamento, aguardando...")
		mutex.Unlock()
		return
	}

	cacheUpdating = true // Marca que o cache esta sendo atualizado
	mutex.Unlock()

	// Realiza a requisicao para o endpoint
	rawResponse, err := utils.FetchRawFromEndpoint("http://protheus-vm:9010/rest/MovPortaria/CarregamentoSaida", nil)
	if err != nil {
		log.Printf("Erro ao buscar historico: %v", err)
		mutex.Lock()
		cacheUpdating = false // Libera a flag de atualizacao em caso de erro
		mutex.Unlock()
		return
	}

	cleanedResponse := cleanJSON(rawResponse)

	var historico []RawHistorico
	if err := json.Unmarshal([]byte(cleanedResponse), &historico); err != nil {
		log.Printf("Erro ao decodificar JSON do historico: %v", err)
		mutex.Lock()
		cacheUpdating = false // Libera a flag de atualizacao em caso de erro
		mutex.Unlock()
		return
	}

	// Atualiza o cache
	mutex.Lock()
	historicoCache = processHistoricos(historico)
	cacheTimestamp = time.Now()
	cacheInitialized = true
	cacheUpdating = false // Libera a flag de atualizacao
	mutex.Unlock()

	log.Println("Cache de historico atualizado com sucesso!")
}

// ===================
// GET
// ===================
func GetHistorico(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()

	// Verifica se o cache esta expirado ou vazio
	if !cacheInitialized || len(historicoCache) == 0 || time.Since(cacheTimestamp) > cacheDuration {
		log.Println("Cache de historico expirado ou vazio. Carregando novos dados em background...")

		// Atualiza o cache em segundo plano
		go fetchHistoricoData()
	}

	// Aplica filtros aos dados do cache
	filterableColumns := []string{"Filial", "NF", "Vendedor", "Cliente", "Produto", "DataHora", "Responsavel", "Placa", "Observacao", "Saldo", "Rec"}
	filteredHistorico := applyFilters(historicoCache, r, filterableColumns)

	// Se nao houver resultados no cache, tenta buscar novos dados
	if len(filteredHistorico) == 0 {
		log.Println("Dados nao encontrados no cache, buscando no webservice...")

		// Libera o mutex enquanto busca no webservice
		mutex.Unlock()

		// Requisita novos dados
		fetchHistoricoData()

		// Bloqueia novamente apos a requisicao
		mutex.Lock()

		// Aplica o filtro novamente apos atualizar o cache
		filteredHistorico = applyFilters(historicoCache, r, filterableColumns)

		// Se ainda assim nao encontrar, retorna erro
		if len(filteredHistorico) == 0 {
			mutex.Unlock()
			http.Error(w, "Nenhum historico disponivel.", http.StatusNotFound)
			return
		}
	}

	// Aplica classificacao e paginacao
	sortedHistorico := applySorting(filteredHistorico, r)
	paginatedHistorico := utils.Paginate(sortedHistorico, r)

	mutex.Unlock()

	// Retorna os dados filtrados, classificados e paginados
	if err := json.NewEncoder(w).Encode(paginatedHistorico); err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

func cleanJSON(raw string) string {
	return strings.ReplaceAll(strings.ReplaceAll(raw, "\n", ""), "\t", "")
}

func applyFilters(historico []Historico, r *http.Request, filterableColumns []string) []Historico {
	return utils.FilterData(historico, r, filterableColumns)
}

func applySorting(historico []Historico, r *http.Request) []Historico {
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	if sortBy == "" {
		sortBy = "DataHora"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	return utils.SortByColumn(historico, sortBy, sortOrder)
}

// ===================
// PROCESSAMENTO
// ===================

func processHistoricos(rawHistoricos []RawHistorico) []Historico {
	var processed []Historico
	for _, raw := range rawHistoricos {
		nf := utils.TrimString(raw.Doc)
		serie := utils.TrimString(raw.Serie)
		nfCompleta := nf
		if serie != "" {
			nfCompleta += " - " + serie
		}

		Data := utils.FormatDate(utils.TrimString(raw.Data), "20060102", "02/01/2006")

		historicoItem := Historico{
			Filial:      utils.TrimString(raw.Filial),
			NF:          nfCompleta,
			Vendedor:    utils.TrimString(raw.CodUsuario) + " - " + utils.TrimString(raw.Origem),
			Cliente:     utils.TrimString(raw.CodCli) + " - " + utils.TrimString(raw.LojaCli) + " - " + utils.TrimString(raw.NomeCliente),
			Produto:     utils.TrimString(raw.CodProd) + " - " + utils.TrimString(raw.ItemProd) + " - " + utils.TrimString(raw.NomeProd),
			Saldo:       raw.Saldo,
			Rec:         raw.Rec,
			DataHora:    Data + " - " + utils.TrimString(raw.Hora),
			Responsavel: utils.TrimString(raw.Respons),
			Placa:       utils.TrimString(raw.Placa),
			Observacao:  utils.TrimString(raw.Obs),
		}

		processed = append(processed, historicoItem)
	}
	return processed
}
