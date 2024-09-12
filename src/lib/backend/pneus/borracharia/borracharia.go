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

// Cache de borracharia para armazenar os dados processados na memória.
var borrachariaCache []Borracharia
var mutex sync.Mutex // Controla o acesso ao cache.

// ===================
// ESTRUTURAS DE DADOS
// ===================

// Borracharia representa a estrutura final dos dados exibidos.
type Borracharia struct {
	Filial, NF, Cliente, Vendedor, Produto, Emissao string
	Saldo                                           int
}

// RawBorracharia mapeia os dados originais recebidos do endpoint.
type RawBorracharia struct {
	Filial, Doc, Serie, Cliente, ClienteNome, Vendedor, VendedorDesc,
	Codigo, Descricao, Emissao string
	Saldo int
}

// ===================
// REQUISIÇÃO E CACHE
// ===================

// StartFetchingBorracharia inicia a atualização periódica do cache de borracharia.
func StartFetchingBorracharia() {
	go func() {
		for {
			log.Println("Buscando dados de borracharia...")
			borracharia, err := utils.FetchFromEndpoint[RawBorracharia]("http://protheus-vm:9010/rest/MovPortaria/NFSaidasDisponiveis", nil)
			if err != nil {
				log.Printf("Erro ao buscar dados de borracharia: %v", err)
				continue
			}

			mutex.Lock()
			borrachariaCache = processBorracharia(borracharia)
			mutex.Unlock()

			log.Println("Cache de borracharia atualizado com sucesso!")
			time.Sleep(1 * time.Minute) // Atualiza a cada minuto.
		}
	}()
}

// ===================
// HANDLER HTTP
// ===================

// GetBorracharia responde com os dados de borracharia atuais em cache, sem filtros ou classificação.
func GetBorracharia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	if len(borrachariaCache) == 0 {
		http.Error(w, "Nenhuma informação de borracharia disponível.", http.StatusNotFound)
		return
	}

	// Retorna os dados sem filtrar, ordenar ou paginar.
	json.NewEncoder(w).Encode(borrachariaCache)
}

// ===================
// PROCESSAMENTO DOS DADOS
// ===================

// processBorracharia converte os dados brutos em dados processados para exibição.
func processBorracharia(rawBorracharia []RawBorracharia) []Borracharia {
	var processed []Borracharia
	for _, raw := range rawBorracharia {
		processed = append(processed, Borracharia{
			Filial:   utils.TrimString(raw.Filial),
			NF:       utils.TrimString(raw.Doc) + " - " + utils.TrimString(raw.Serie),
			Cliente:  utils.TrimString(raw.Cliente) + " " + utils.TrimString(raw.ClienteNome),
			Vendedor: utils.TrimString(raw.Vendedor) + " " + utils.TrimString(raw.VendedorDesc),
			Produto:  utils.TrimString(raw.Codigo) + " " + utils.TrimString(raw.Descricao),
			Saldo:    raw.Saldo,
			Emissao:  utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
		})
	}
	return processed
}
