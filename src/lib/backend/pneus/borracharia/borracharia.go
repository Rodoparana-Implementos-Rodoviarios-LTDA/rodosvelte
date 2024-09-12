package borracharia

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

var borrachariaCache []Borracharia
var mutex sync.Mutex

// ===================
// ESTRUTURAS DE DADOS
// ===================

// Borracharia representa os dados processados da borracharia.
type Borracharia struct {
	Filial, NF, Vendedor, Cliente, Produto, Emissao string
	Saldo                                           int
}

// RawBorracharia mapeia os dados brutos recebidos do endpoint.
type RawBorracharia struct {
	Filial       string `json:"D2_FILIAL"`
	Documento    string `json:"D2_DOC"`
	Serie        string `json:"D2_SERIE"`
	Emissao      string `json:"D2_EMISSAO"`
	Vendedor     string `json:"F2_VEND1"`
	VendedorNome string `json:"A3_NOME"`
	Cliente      string `json:"D2_CLIENTE"`
	Loja         string `json:"D2_LOJA"`
	ClienteNome  string `json:"A1_NOME"`
	Codigo       string `json:"D2_COD"`
	Item         string `json:"D2_ITEM"`
	Descricao    string `json:"B1_DESC"`
	Saldo        int    `json:"SALDO"`
}

// ===================
// REQUISIÇÃO E CACHE
// ===================

// StartFetchingBorracharia inicia a atualização periódica do cache de dados de borracharia.
func StartFetchingBorracharia() {
	go func() {
		for {
			borracharia, err := utils.FetchFromEndpoint[RawBorracharia]("http://protheus-vm:9010/rest/MovPortaria/NFSaidasDisponiveis", nil)
			if err != nil {
				log.Printf("Erro ao buscar dados de borracharia: %v", err)
				continue
			}

			mutex.Lock()
			borrachariaCache = processBorracharia(borracharia)
			mutex.Unlock()

			log.Println("Cache de borracharia atualizado com sucesso!")
			time.Sleep(1 * time.Minute) // Atualiza o cache a cada minuto.
		}
	}()
}

// ===================
// HANDLER HTTP
// ===================

// GetBorracharia responde com os dados de borracharia filtrados, classificados e paginados.
func GetBorracharia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	if len(borrachariaCache) == 0 {
		http.Error(w, "Nenhuma informação de borracharia disponível.", http.StatusNotFound)
		return
	}

	// Aplicar Filtros
	filterableColumns := []string{"Filial", "NF", "Vendedor", "Cliente", "Produto", "Emissao"}
	filteredBorracharia := applyFilters(borrachariaCache, r, filterableColumns)

	// Aplicar Classificação
	sortedBorracharia := applySorting(filteredBorracharia, r)

	// Aplicar Paginação
	paginatedBorracharia := utils.Paginate(sortedBorracharia, r)

	// Retorna os dados filtrados, classificados e paginados
	json.NewEncoder(w).Encode(paginatedBorracharia)
}

// ===================
// FUNÇÕES AUXILIARES
// ===================

// Função para aplicar filtros com base nos headers.
func applyFilters(borracharia []Borracharia, r *http.Request, filterableColumns []string) []Borracharia {
	return utils.FilterData(borracharia, r, filterableColumns)
}

// Função para aplicar a classificação com base nos headers.
func applySorting(borracharia []Borracharia, r *http.Request) []Borracharia {
	// Headers que determinam a coluna e ordem de classificação
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	// Define valores padrão se os headers não forem fornecidos
	if sortBy == "" {
		sortBy = "Emissao" // Valor padrão
	}
	if sortOrder == "" {
		sortOrder = "desc" // Valor padrão
	}

	// Aplica a função de classificação genérica do utils
	return utils.SortByColumn(borracharia, sortBy, sortOrder)
}

// ===================
// PROCESSAMENTO DOS DADOS
// ===================

// processBorracharia converte os dados brutos em dados processados para exibição.
func processBorracharia(rawBorracharia []RawBorracharia) []Borracharia {
	var processed []Borracharia
	for _, raw := range rawBorracharia {
		borracharia := Borracharia{
			Filial:   utils.TrimString(raw.Filial),
			NF:       utils.TrimString(raw.Documento) + " - " + utils.TrimString(raw.Serie),
			Vendedor: utils.TrimString(raw.Vendedor) + " " + utils.TrimString(raw.VendedorNome),
			Cliente:  utils.TrimString(raw.Cliente) + " " + utils.TrimString(raw.Loja) + " " + utils.TrimString(raw.ClienteNome),
			Produto:  utils.TrimString(raw.Codigo) + " " + utils.TrimString(raw.Item) + " " + utils.TrimString(raw.Descricao),
			Emissao:  utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
			Saldo:    raw.Saldo,
		}
		processed = append(processed, borracharia)
	}
	return processed
}
