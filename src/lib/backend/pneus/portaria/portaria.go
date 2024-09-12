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
// VARIÁVEIS GLOBAIS
// ===================

var movimentosPortariaCache []MovimentoPortaria
var mutex sync.Mutex

// ===================
// ESTRUTURAS DE DADOS
// ===================

// MovimentoPortaria representa os dados processados do movimento.
type MovimentoPortaria struct {
	Filial, NF, Cliente, Produto, TipoMov, DataHora, Responsavel, Placa, Observacao string
	Saldo                                                                           int
}

// RawMovimentoPortaria mapeia os dados brutos recebidos do endpoint.
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
}

// ===================
// REQUISIÇÃO E CACHE
// ===================

// StartFetchingMovimentosPortaria inicia a atualização periódica do cache de movimentos da portaria.
func StartFetchingMovimentosPortaria() {
	go func() {
		for {
			log.Println("Buscando movimentos da portaria...")
			movimentos, err := utils.FetchFromEndpoint[RawMovimentoPortaria]("http://protheus-vm:9010/rest/MovPortaria/MovimentosPorNF", nil)
			if err != nil {
				log.Printf("Erro ao buscar movimentos da portaria: %v", err)
				continue
			}

			mutex.Lock()
			movimentosPortariaCache = processMovimentosPortaria(movimentos)
			mutex.Unlock()

			log.Println("Cache de movimentos da portaria atualizado com sucesso!")
			time.Sleep(1 * time.Minute) // Atualiza o cache a cada minuto
		}
	}()
}

// ===================
// HANDLER HTTP
// ===================

// GetMovimentosPortaria responde com os dados de movimentos da portaria filtrados, classificados e paginados.
func GetMovimentosPortaria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	if len(movimentosPortariaCache) == 0 {
		http.Error(w, "Nenhum movimento disponível.", http.StatusNotFound)
		return
	}

	// Aplicar Filtros
	filterableColumns := []string{"Filial", "NF", "Cliente", "Produto", "DataHora", "TipoMov"}
	filteredMovimentos := applyFilters(movimentosPortariaCache, r, filterableColumns)

	// Aplicar Classificação
	sortedMovimentos := applySorting(filteredMovimentos, r)

	// Aplicar Paginação
	paginatedMovimentos := utils.Paginate(sortedMovimentos, r)

	// Retorna os dados filtrados, classificados e paginados
	json.NewEncoder(w).Encode(paginatedMovimentos)
}

// ===================
// FUNÇÕES AUXILIARES
// ===================

// Função para aplicar filtros com base nos headers.
func applyFilters(movimentos []MovimentoPortaria, r *http.Request, filterableColumns []string) []MovimentoPortaria {
	return utils.FilterData(movimentos, r, filterableColumns)
}

// Função para aplicar a classificação com base nos headers.
func applySorting(movimentos []MovimentoPortaria, r *http.Request) []MovimentoPortaria {
	// Headers que determinam a coluna e ordem de classificação
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	// Define valores padrão se os headers não forem fornecidos
	if sortBy == "" {
		sortBy = "DataHora" // Valor padrão
	}
	if sortOrder == "" {
		sortOrder = "desc" // Valor padrão
	}

	// Aplica a função de classificação genérica do utils
	return utils.SortByColumn(movimentos, sortBy, sortOrder)
}

// ===================
// PROCESSAMENTO DOS DADOS
// ===================

// processMovimentosPortaria converte os dados brutos em dados processados para exibição.
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
			Saldo:       raw.Saldo,
		}
		processed = append(processed, movimento)
	}
	return processed
}
