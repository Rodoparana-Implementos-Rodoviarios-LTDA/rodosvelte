package portaria

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

var movimentosPortariaCache []MovimentoPortaria
var mutex sync.Mutex

// ===================
// ESTRUTURAS
// ===================

type MovimentoPortaria struct {
	Filial      string `json:"Filial"`
	NF          string `json:"NF"`       // Documento + Série
	Vendedor    string `json:"Vendedor"` // F2_VEND1 + A3_NOME
	Cliente     string `json:"Cliente"`  // D2_CLIENTE + D2_LOJA + A1_NOME
	Produto     string `json:"Produto"`  // D2_COD + D2_ITEM + B1_DESC
	Saldo       int    `json:"Saldo"`    // SALDO
	DataHora    string `json:"DataHora"` // Data + Hora combinadas
	Responsavel string `json:"Responsavel"`
	Placa       string `json:"Placa"`
	Observacao  string `json:"Observacao"`
}

type RawMovimentoPortaria struct {
	Filial       string `json:"D2_FILIAL"`
	Documento    string `json:"D2_DOC"`
	Serie        string `json:"D2_SERIE"`
	Cliente      string `json:"D2_CLIENTE"`
	Loja         string `json:"D2_LOJA"`
	ClienteNome  string `json:"A1_NOME"`
	Vendedor     string `json:"F2_VEND1"`
	VendedorNome string `json:"A3_NOME"`
	Codigo       string `json:"D2_COD"`
	Item         string `json:"D2_ITEM"`
	Descricao    string `json:"B1_DESC"`
	Saldo        int    `json:"SALDO"`
	Data         string `json:"Z08_DATA"`
	Hora         string `json:"Z08_HORA"`
	Responsavel  string `json:"Z08_RESPON"`
	Placa        string `json:"Z08_PLACA"`
	Observacao   string `json:"Z08_OBSERV"`
}

// ===================
// REQUISIÇÃO HTTP
// ===================

func StartFetchingMovimentosPortaria() {
	go func() {
		for {
			log.Println("Buscando movimentos da portaria...")
			movimentos, err := utils.FetchFromEndpoint[RawMovimentoPortaria]("http://protheus-vm:9010/rest/MovPortaria/MovimentosPorNF", nil)
			if err != nil {
				log.Printf("Erro ao buscar movimentos da portaria: %v", err)
			} else {
				processedMovimentos := processMovimentosPortaria(movimentos)
				mutex.Lock()
				movimentosPortariaCache = processedMovimentos // Atualiza o cache
				mutex.Unlock()
				log.Println("Importação de movimentos da portaria concluída com sucesso!")
			}
			time.Sleep(1 * time.Minute) // Intervalo de 1 minuto
		}
	}()
}

func GetMovimentosPortaria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pega os parâmetros de paginação e ordenação
	page, pageSize := utils.GetPaginationParams(r)
	filters := utils.GetFilters(r)
	sortBy, sortOrder := utils.GetSortingParams(r, map[string]bool{
		"Filial":   true,
		"NF":       true,
		"Cliente":  true,
		"DataHora": true,
		"TipoMov":  true,
	})

	mutex.Lock()
	defer mutex.Unlock()

	if len(movimentosPortariaCache) == 0 {
		http.Error(w, "Nenhum movimento disponível.", http.StatusNotFound)
		return
	}

	// Chama as funções auxiliares para filtrar, ordenar e paginar
	filteredMovimentos := filterMovimentosPortaria(movimentosPortariaCache, filters)
	sortedMovimentos := sortMovimentosPortaria(filteredMovimentos, sortBy, sortOrder)
	paginatedMovimentos := paginateMovimentosPortaria(sortedMovimentos, page, pageSize)

	// Retorna os dados paginados
	json.NewEncoder(w).Encode(paginatedMovimentos)
}

// Função para processar e ajustar os dados dos movimentos
func processMovimentosPortaria(rawMovimentos []RawMovimentoPortaria) []MovimentoPortaria {
	var processed []MovimentoPortaria
	for _, raw := range rawMovimentos {
		movimento := MovimentoPortaria{
			Filial:      utils.TrimString(raw.Filial),
			NF:          utils.TrimString(raw.Documento) + " - " + utils.TrimString(raw.Serie),
			Vendedor:    utils.TrimString(raw.Vendedor) + " - " + utils.TrimString(raw.VendedorNome),
			Cliente:     utils.TrimString(raw.Cliente) + " - " + utils.TrimString(raw.Loja) + " - " + utils.TrimString(raw.ClienteNome),
			Produto:     utils.TrimString(raw.Codigo) + " - " + utils.TrimString(raw.Item) + " - " + utils.TrimString(raw.Descricao),
			Saldo:       raw.Saldo,
			DataHora:    utils.FormatDate(utils.TrimString(raw.Data), "20060102", "02/01/2006") + " " + utils.TrimString(raw.Hora),
			Responsavel: utils.TrimString(raw.Responsavel),
			Placa:       utils.TrimString(raw.Placa),
			Observacao:  utils.TrimString(raw.Observacao),
		}
		processed = append(processed, movimento)
	}
	return processed
}

// ===================
// PARÂMETROS DA REQUISIÇÃO
// ===================

func filterMovimentosPortaria(movimentos []MovimentoPortaria, filters map[string]string) []MovimentoPortaria {
	return utils.FilterData(movimentos, filters, func(item MovimentoPortaria, key, value string) bool {
		switch key {
		case "Filial":
			return item.Filial == value
		case "NF":
			return item.NF == value
		case "Cliente":
			return item.Cliente == value
		case "DataHora":
			return item.DataHora == value
		default:
			return true
		}
	})
}

func sortMovimentosPortaria(movimentos []MovimentoPortaria, sortBy, sortOrder string) []MovimentoPortaria {
	return utils.SortData(movimentos, sortBy, sortOrder, func(i, j MovimentoPortaria) bool {
		switch sortBy {
		case "Filial":
			return i.Filial < j.Filial
		case "NF":
			return i.NF < j.NF
		case "Cliente":
			return i.Cliente < j.Cliente
		case "DataHora":
			return i.DataHora < j.DataHora
		default:
			return true
		}
	})
}

func paginateMovimentosPortaria(movimentos []MovimentoPortaria, page, pageSize int) []MovimentoPortaria {
	return utils.Paginate(movimentos, page, pageSize)
}
