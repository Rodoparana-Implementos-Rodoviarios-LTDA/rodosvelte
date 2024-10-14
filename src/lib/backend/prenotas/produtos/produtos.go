package produtos

import (
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// ===================
// VARIAVEIS GLOBAIS
// ===================
var (
	produtosCache    []Produto  // Cache dos dados dos produtos
	mutex            sync.Mutex // Controle de concorrência para o cache
	cacheInitialized bool       // Flag para garantir que o cache foi inicializado ao menos uma vez
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type Produto struct {
	Codigo, Descricao, Origem, Unidade, Grupo, Tipo, NCM string
}

type RawProduto struct {
	Codigo    string `json:"B1_COD"`
	Descricao string `json:"B1_DESC"`
	Origem    string `json:"B1_ORIGEM"`
	Unidade   string `json:"B1_UM"`
	Grupo     string `json:"B1_GRUPO"`
	Tipo      string `json:"B1_TIPO"`
	NCM       string `json:"B1_POSIPI"`
}

// ===================
// REQUISICAO E CACHE
// ===================

// InitializeCache faz o carregamento inicial de todos os registros de uma vez
func InitializeCache() {
	log.Println("Carregando dados de todos os produtos...")

	mutex.Lock()
	defer mutex.Unlock()

	// Faz a requisição para buscar todos os registros de uma vez
	produtos, err := fetchProdutosFromAPI("", "")
	if err != nil {
		log.Printf("Erro ao buscar produtos: %v", err)
		return
	}

	// Processa os dados e atualiza o cache
	produtosCache = processProdutos(produtos)
	cacheInitialized = true
	log.Println("Cache de produtos carregado com sucesso!")
}

// ===================
// HANDLER PRINCIPAL
// ===================

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	// Verifica se o cache foi inicializado
	if !cacheInitialized {
		http.Error(w, "Cache não inicializado ainda.", http.StatusInternalServerError)
		return
	}

	// Captura o valor do header 'tipo'
	tipo := r.Header.Get("tipo")

	var grupos, ngrupos string

	// Define o grupo ou ngrupo com base no tipo de produto
	switch tipo {
	case "Despesa/Imobilizado":
		grupos = "'0100','0500'"
	case "Collection":
		grupos = "'0228'"
	case "Materia Prima":
		grupos = " 	"
	case "Revenda":
		ngrupos = "'0100','0500','0276','0270','0228','0001','0003','0006','0039','0100','0105','0260','0301','0306','0311','0500','0501','0999','MAQU','VEIC'"
	default:
		grupos = "" // Se não for nenhum dos tipos, busca todos os produtos
	}

	// Busca os produtos diretamente da API com base no 'GRUPO' ou 'NGRUPO'
	produtos, err := fetchProdutosFromAPI(grupos, ngrupos)
	if err != nil {
		http.Error(w, "Erro ao buscar produtos da API.", http.StatusInternalServerError)
		return
	}

	// Retorna os produtos diretamente, sem cache
	err = json.NewEncoder(w).Encode(produtos)
	if err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

// Faz a requisição à API usando 'GRUPO' ou 'NGRUPO'
func fetchProdutosFromAPI(grupos, ngrupos string) ([]RawProduto, error) {
	url := "http://172.16.99.174:8400/rest/reidoapsdu/consultar/likeprod"
	headers := make(map[string]string)

	if grupos != "" {
		headers["GRUPO"] = grupos
	}
	if ngrupos != "" {
		headers["NGRUPO"] = ngrupos
	}

	produtos, err := utils.FetchFromEndpoint[RawProduto](url, headers)
	if err != nil {
		return nil, err
	}

	return produtos, nil
}

// Função que processa uma lista de produtos e converte para a estrutura final
func processProdutos(rawProdutos []RawProduto) []Produto {
	var processed []Produto
	for _, raw := range rawProdutos {
		processed = append(processed, processProduto(raw))
	}
	return processed
}

// Função que processa um único produto bruto e o converte para o formato final
func processProduto(raw RawProduto) Produto {
	return Produto{
		Codigo:    utils.TrimString(raw.Codigo),
		Descricao: utils.TrimString(raw.Descricao),
		Origem:    utils.TrimString(raw.Origem),
		Unidade:   utils.TrimString(raw.Unidade),
		Grupo:     utils.TrimString(raw.Grupo),
		Tipo:      utils.TrimString(raw.Tipo),
		NCM:       utils.TrimString(raw.NCM),
	}
}
