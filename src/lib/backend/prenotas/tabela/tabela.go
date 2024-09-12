package tabela

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

// preNotasCache mantém em memória os dados das pre_notas, evitando buscas frequentes no endpoint.
var preNotasCache []PreNota
var mutex sync.Mutex // Controla o acesso concorrente ao cache.

// ===================
// ESTRUTURAS DE DADOS
// ===================

// PreNota representa a estrutura final da nota.
type PreNota struct {
	Filial, NF, Status, Fornecedor, Emissao, Inclusao, Vencimento, Valor,
	Tipo, Prioridade, Usuario, Obs, ObsRev, Revisao, Rec string
}

// RawPreNota mapeia os dados originais do endpoint.
type RawPreNota struct {
	Filial, Doc, Serie, Status, Fornece, Emissao, Inclusao, Vencimento,
	ValBrut, Tipo, Prior, UsrRa, Obs, ObsRev, Rev, Rec string
}

// ===================
// REQUISIÇÃO E CACHE
// ===================

// StartFetchingPreNotas inicia a atualização periódica do cache de pre_notas.
func StartFetchingPreNotas() {
	go func() {
		for {
			log.Println("Buscando pre_notas...")
			preNotas, err := utils.FetchFromEndpoint[RawPreNota]("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=99999", nil)
			if err != nil {
				log.Printf("Erro ao buscar pre_notas: %v", err)
				continue
			}

			mutex.Lock()
			preNotasCache = processPreNotas(preNotas)
			mutex.Unlock()
			log.Println("Cache de pre_notas atualizado com sucesso!")
			time.Sleep(1 * time.Minute) // Atualiza o cache a cada minuto
		}
	}()
}
// ===================
// GET
// ===================

// GetPreNotas responde com os dados de pre_notas classificados por uma coluna específica.
func GetPreNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pega os parâmetros de classificação dos headers
	sortBy := r.Header.Get("X-Sort-By")       // Header para a coluna de classificação
	sortOrder := r.Header.Get("X-Sort-Order") // Header para a ordem de classificação
	if sortBy == "" {
		sortBy = "Inclusao" // Valor padrão
	}
	if sortOrder == "" {
		sortOrder = "desc" // Valor padrão
	}

	mutex.Lock()
	defer mutex.Unlock()

	if len(preNotasCache) == 0 {
		http.Error(w, "Nenhuma pre_nota disponível.", http.StatusNotFound)
		return
	}

	// Chama a função genérica de classificação
	sortedPreNotas := utils.SortByColumn(preNotasCache, sortBy, sortOrder)

	// Retorna os dados classificados
	json.NewEncoder(w).Encode(sortedPreNotas)
}

// ===================
// PROCESSAMENTO
// ===================

// processPreNotas transforma os dados brutos em pre_notas processadas.
func processPreNotas(rawPreNotas []RawPreNota) []PreNota {
	var processed []PreNota
	for _, raw := range rawPreNotas {
		processed = append(processed, PreNota{
			Filial:     utils.TrimString(raw.Filial),
			NF:         utils.TrimString(raw.Doc) + " - " + utils.TrimString(raw.Serie),
			Status:     utils.DetermineStatus(raw.Status, raw.Rev),
			Fornecedor: utils.TrimString(raw.Fornece),
			Emissao:    utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
			Inclusao:   utils.FormatDate(utils.TrimString(raw.Inclusao), "20060102", "02/01/2006"),
			Vencimento: utils.FormatDate(utils.TrimString(raw.Vencimento), "20060102", "02/01/2006"),
			Valor:      utils.FormatCurrency(utils.TrimString(raw.ValBrut)),
			Tipo:       utils.TrimString(raw.Tipo),
			Prioridade: utils.TrimString(raw.Prior),
			Usuario:    utils.TrimString(raw.UsrRa),
			Obs:        utils.TrimString(raw.Obs),
			ObsRev:     utils.TrimString(raw.ObsRev),
			Revisao:    utils.TrimString(raw.Rev),
			Rec:        utils.TrimString(raw.Rec),
		})
	}
	return processed
}

