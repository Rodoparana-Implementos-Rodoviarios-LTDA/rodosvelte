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

// Cache para armazenar os dados dos movimentos da portaria.
var movimentosPortariaCache []MovimentoPortaria
var mutex sync.Mutex // Controla o acesso concorrente ao cache.

// ===================
// ESTRUTURAS DE DADOS
// ===================

// MovimentoPortaria representa os dados processados do movimento.
type MovimentoPortaria struct {
	Filial, NF, Vendedor, Cliente, Produto, DataHora, Responsavel, Placa, Observacao string
	Saldo                                                                            int
}

// RawMovimentoPortaria mapeia os dados brutos recebidos do endpoint.
type RawMovimentoPortaria struct {
	Filial, Documento, Serie, Cliente, Loja, ClienteNome, Vendedor, VendedorNome,
	Codigo, Item, Descricao, Data, Hora, Responsavel, Placa, Observacao string
	Saldo int
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

// GetMovimentosPortaria responde com os dados de movimentos da portaria atuais em cache.
func GetMovimentosPortaria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	if len(movimentosPortariaCache) == 0 {
		http.Error(w, "Nenhum movimento disponível.", http.StatusNotFound)
		return
	}

	// Retorna os dados diretamente do cache, sem filtros ou ordenação.
	json.NewEncoder(w).Encode(movimentosPortariaCache)
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
