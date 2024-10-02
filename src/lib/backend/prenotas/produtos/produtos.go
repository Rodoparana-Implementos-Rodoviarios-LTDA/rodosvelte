package produtos

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
	produtosCache    []Produto  // Cache dos dados dos produtos
	cacheTimestamp   time.Time  // Timestamp do cache
	mutex            sync.Mutex // Controle de concorrencia para o cache
	cacheInitialized bool       // Flag para garantir que o cache foi inicializado ao menos uma vez
	cacheUpdating    bool       // Flag para indicar se o cache esta sendo atualizado
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type Produto struct {
	Codigo      string // Código do produto
	Descricao   string // Descrição do produto
	CodDesc     string // Código e Descrição combinados
	Origem      string // Origem do produto
	UnidadeMed  string // Unidade de medida
	Localizacao string // Localização do produto
	Grupo       string // Grupo do produto
	Tipo        string // Tipo do produto
	GrupoTrib   string // Grupo de tributação
	NCM         string // NCM (Posição IPI)
}

type RawProduto struct {
	Codigo      string `json:"B1_COD"`    // Código
	Descricao   string `json:"B1_DESC"`   // Descrição
	CodDesc     string `json:"DESC"`      // Código + Descrição
	Origem      string `json:"B1_ORIGEM"` // Origem
	UnidadeMed  string `json:"B1_UM"`     // Unidade de medida
	Localizacao string `json:"B1_LOCPAD"` // Localização
	Grupo       string `json:"B1_GRUPO"`  // Grupo
	Tipo        string `json:"B1_TIPO"`   // Tipo
	GrupoTrib   string `json:"B1_GRTRIB"` // Grupo Tributário
	NCM         string `json:"B1_POSIPI"` // NCM (Posição IPI)
}

// ===================
// REQUISICAO E CACHE
// ===================

// InitializeCache inicializa o cache de produtos com todos os produtos do sistema
func InitializeCache() {
	log.Println("Carregando dados iniciais de produtos...")
	fetchProdutosData()

	// Inicia o polling para buscar novos produtos incrementais
	go startPolling()
}

// Funcao para buscar os produtos iniciais
func fetchProdutosData() {
	mutex.Lock()
	defer mutex.Unlock()

	// Faz a requisicao para carregar todos os produtos
	produtos, err := utils.FetchFromEndpoint[RawProduto]("http://172.16.99.174:8400/rest/reidoapsdu/consultar/likeprod", nil)
	if err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return
	}

	// Processa os produtos e armazena no cache
	produtosCache = processProdutos(produtos)
	cacheTimestamp = time.Now()
	cacheInitialized = true

	log.Println("Cache de produtos carregado com sucesso!")
}

// Polling para verificar se há novos produtos a cada minuto
func startPolling() {
	for {
		time.Sleep(1 * time.Minute)
		fetchProdutosIncremental()
	}
}

// Funcao para buscar produtos incrementais
func fetchProdutosIncremental() {
	mutex.Lock()
	defer mutex.Unlock()

	if !cacheInitialized {
		log.Println("Cache nao inicializado ainda. Polling abortado.")
		return
	}

	// Faz a requisicao para buscar produtos mais recentes
	produtosNovos, err := utils.FetchFromEndpoint[RawProduto]("http://172.16.99.174:8400/rest/reidoapsdu/consultar/likeprod", nil)
	if err != nil {
		log.Printf("Erro ao buscar produtos incrementais: %v", err)
		return
	}

	// Adiciona novos produtos ao cache
	if len(produtosNovos) > len(produtosCache) {
		novaQuantidade := len(produtosNovos) - len(produtosCache)
		produtosCache = append(produtosCache, processProdutos(produtosNovos[len(produtosNovos)-novaQuantidade:]...)...)
		log.Printf("Novos produtos adicionados ao cache: %d novos registros.", novaQuantidade)
	} else {
		log.Println("Nenhum novo produto encontrado.")
	}
}

// ===================
// HANDLER HTTP
// ===================

// GetProdutos responde com os dados filtrados dos produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	// Verifica se o cache foi inicializado
	if !cacheInitialized || len(produtosCache) == 0 {
		http.Error(w, "Nenhum produto disponivel.", http.StatusNotFound)
		return
	}

	// Aplica os filtros
	filterableColumns := []string{"Codigo", "Descricao", "CodDesc", "Origem", "UnidadeMed", "Localizacao", "Grupo", "Tipo", "GrupoTrib", "NCM"}
	filteredProdutos := applyFilters(produtosCache, r, filterableColumns)

	// Retorna os dados filtrados
	err := json.NewEncoder(w).Encode(filteredProdutos)
	if err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

// Funcao para processar os produtos brutos
func processProdutos(rawProdutos ...RawProduto) []Produto {
	var processed []Produto
	for _, raw := range rawProdutos {
		produto := Produto{
			Codigo:      utils.TrimString(raw.Codigo),
			Descricao:   utils.TrimString(raw.Descricao),
			CodDesc:     utils.TrimString(raw.CodDesc),
			Origem:      utils.TrimString(raw.Origem),
			UnidadeMed:  utils.TrimString(raw.UnidadeMed),
			Localizacao: utils.TrimString(raw.Localizacao),
			Grupo:       utils.TrimString(raw.Grupo),
			Tipo:        utils.TrimString(raw.Tipo),
			GrupoTrib:   utils.TrimString(raw.GrupoTrib),
			NCM:         utils.TrimString(raw.NCM),
		}
		processed = append(processed, produto)
	}
	return processed
}

// Funcao para aplicar filtros nos dados de produtos com base nos headers
func applyFilters(produtos []Produto, r *http.Request, filterableColumns []string) []Produto {
	return utils.FilterData(produtos, r, filterableColumns)
}
