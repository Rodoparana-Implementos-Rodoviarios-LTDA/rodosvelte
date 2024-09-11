package borracharia

import (
	"backend/utils" // Pacote utils com as funções genéricas
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// ===================
// VARIÁVEIS
// ===================

// borrachariaCache armazena em cache os dados de borracharia na memória.
var borrachariaCache []Borracharia

// mutex é usado para garantir a segurança de acesso ao cache em concorrência.
var mutex sync.Mutex

// ===================
// ESTRUTURAS
// ===================

// Estrutura final para exibição dos dados
type Borracharia struct {
	Filial   string `json:"Filial"`
	NF       string `json:"NF"` // Alterado para NF
	Cliente  string `json:"Cliente"`
	Vendedor string `json:"Vendedor"`
	Produto  string `json:"Produto"`
	Saldo    int    `json:"Saldo"`
	Emissao  string `json:"Emissao"`
}

// Estrutura original dos dados recebidos do endpoint
type RawBorracharia struct {
	Filial       string `json:"D2_FILIAL"`
	Doc          string `json:"D2_DOC"`
	Serie        string `json:"D2_SERIE"`
	Cliente      string `json:"D2_CLIENTE"`
	Loja         string `json:"D2_LOJA"`
	ClienteNome  string `json:"A1_NOME"`
	Vendedor     string `json:"F2_VEND1"`
	VendedorDesc string `json:"A3_NOME"`
	Codigo       string `json:"D2_COD"`
	Descricao    string `json:"B1_DESC"`
	Saldo        int    `json:"SALDO"`
	Emissao      string `json:"D2_EMISSAO"`
}

// ===================
// REQUISIÇÃO HTTP
// ===================

// Função para iniciar a busca periódica dos dados de borracharia
func StartFetchingBorracharia() {
	go func() {
		for {
			log.Println("Buscando dados de borracharia...")
			borracharia, err := utils.FetchFromEndpoint[RawBorracharia]("http://protheus-vm:9010/rest/MovPortaria/NFSaidasDisponiveis", nil)
			if err != nil {
				log.Printf("Erro ao buscar dados de borracharia: %v", err)
			} else {
				processedBorracharia := processBorracharia(borracharia)
				mutex.Lock()
				borrachariaCache = processedBorracharia // Atualiza o cache
				mutex.Unlock()
				log.Println("Importação de dados de borracharia concluída com sucesso!")
			}
			time.Sleep(1 * time.Minute) // Intervalo de 1 minuto
		}
	}()
}

// ===================
// HANDLER HTTP
// ===================

// Handler HTTP para buscar dados de borracharia com paginação, filtro e ordenação
func GetBorracharia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pega os parâmetros de paginação e ordenação
	page, pageSize := utils.GetPaginationParams(r)
	filters := utils.GetFilters(r)
	sortBy, sortOrder := utils.GetSortingParams(r, map[string]bool{
		"Filial":   true,
		"NF":       true, // Alterado para NF
		"Cliente":  true,
		"Vendedor": true,
		"Produto":  true,
		"Saldo":    true,
		"Emissao":  true,
	})

	mutex.Lock()
	defer mutex.Unlock()

	if len(borrachariaCache) == 0 {
		http.Error(w, "Nenhuma informação de borracharia disponível.", http.StatusNotFound)
		return
	}

	// Chama as funções auxiliares para filtrar, ordenar e paginar
	filteredBorracharia := filterBorracharia(borrachariaCache, filters)

	// Ordenação padrão por Emissao, do mais recente para o mais antigo
	if sortBy == "" {
		sortBy = "Emissao"
		sortOrder = "desc"
	}

	sortedBorracharia := sortBorracharia(filteredBorracharia, sortBy, sortOrder)
	paginatedBorracharia := paginateBorracharia(sortedBorracharia, page, pageSize)

	// Retorna os dados paginados
	json.NewEncoder(w).Encode(paginatedBorracharia)
}

// ===================
// PROCESSAMENTO DOS DADOS
// ===================

// Função para processar e ajustar os dados de borracharia
func processBorracharia(rawBorracharia []RawBorracharia) []Borracharia {
	var processed []Borracharia
	for _, raw := range rawBorracharia {
		borracharia := Borracharia{
			Filial:   utils.TrimString(raw.Filial),
			NF:       utils.TrimString(raw.Doc) + " - " + utils.TrimString(raw.Serie), // Alterado para NF
			Cliente:  utils.TrimString(raw.Cliente) + " " + utils.TrimString(raw.ClienteNome),
			Vendedor: utils.TrimString(raw.Vendedor) + " " + utils.TrimString(raw.VendedorDesc),
			Produto:  utils.TrimString(raw.Codigo) + " " + utils.TrimString(raw.Descricao),
			Saldo:    raw.Saldo,
			Emissao:  utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
		}
		processed = append(processed, borracharia)
	}
	return processed
}

// ===================
// PARÂMETROS DA REQUISIÇÃO
// ===================

// Função para filtrar dados de borracharia
func filterBorracharia(borracharia []Borracharia, filters map[string]string) []Borracharia {
	return utils.FilterData(borracharia, filters, func(item Borracharia, key, value string) bool {
		switch key {
		case "Filial":
			return item.Filial == value
		case "NF": // Alterado para NF
			return item.NF == value
		case "Cliente":
			return item.Cliente == value
		case "Vendedor":
			return item.Vendedor == value
		case "Produto":
			return item.Produto == value
		case "Saldo":
			return strconv.Itoa(item.Saldo) == value
		case "Emissao":
			return item.Emissao == value
		default:
			return true
		}
	})
}

// Função para ordenar dados de borracharia
func sortBorracharia(borracharia []Borracharia, sortBy, sortOrder string) []Borracharia {
	return utils.SortData(borracharia, sortBy, sortOrder, func(i, j Borracharia) bool {
		switch sortBy {
		case "Filial":
			return i.Filial < j.Filial
		case "NF": // Alterado para NF
			return i.NF < j.NF
		case "Cliente":
			return i.Cliente < j.Cliente
		case "Vendedor":
			return i.Vendedor < j.Vendedor
		case "Produto":
			return i.Produto < j.Produto
		case "Saldo":
			return i.Saldo < j.Saldo
		case "Emissao":
			return i.Emissao < j.Emissao
		default:
			return true
		}
	})
}

// Função para paginar dados de borracharia
func paginateBorracharia(borracharia []Borracharia, page, pageSize int) []Borracharia {
	return utils.Paginate(borracharia, page, pageSize)
}
