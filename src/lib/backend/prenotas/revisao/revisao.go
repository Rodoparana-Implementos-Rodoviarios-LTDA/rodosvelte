package revisao

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
	"backend/utils"
)

// ===================
// VARIÁVEIS GLOBAIS
// ===================

var revisaoCache map[string][]Revisao // Cache para armazenar revisões por chave
var mutex sync.Mutex                  // Exclusão mútua para proteger o cache

// ===================
// ESTRUTURAS
// ===================

type Revisao struct {
	Usuario   string `json:"usuario"`
	Data      string `json:"data"`   // Mantém o valor da data como está
	Hora      string `json:"hora"`   // Mantém o valor da hora como está
	Campo     string `json:"campo"`
	Anterior  string `json:"anterior"`
	Chave     string `json:"chave"`  // Também o campo "rec"
	Atual     string `json:"atual"`
	Status    string `json:"status"` // Campo status adicionado
}

// Estrutura para o agrupamento por `rec` e histórico
type RevisaoAgrupada struct {
	Rec       string    `json:"rec"`
	Historico []Revisao `json:"historico"`
}

// ===================
// REQUISIÇÃO HTTP
// ===================

// Função para iniciar a busca periódica das revisões
func StartFetchingRevisao() {
	go func() {
		for {
			log.Println("Buscando revisões...")

			// Faz a requisição utilizando a função genérica de utils com headers (Content-Type)
			revisoes, err := utils.FetchFromEndpoint[Revisao](
				"http://protheus-app:8400/rest/reidoapsdu/getHistorico",
				map[string]string{"Content-Type": "application/json"}, // Headers com Content-Type
			)
			if err != nil {
				log.Printf("Erro ao buscar revisões: %v", err)
			} else {
				processedRevisoes := processRevisoes(revisoes)

				mutex.Lock()
				revisaoCache = processedRevisoes // Atualiza o cache
				mutex.Unlock()

				log.Println("Importação de revisões concluída com sucesso!")
			}
			time.Sleep(1 * time.Minute) // Atualiza o cache a cada 1 minuto
		}
	}()
}

// ===================
// PROCESSAMENTO DE DADOS
// ===================

// Função que processa as revisões, adiciona o campo de status, e organiza por `rec`
func processRevisoes(revisoes []Revisao) map[string][]Revisao {
	processed := make(map[string][]Revisao)

	for _, revisao := range revisoes {
		// Define o status com base no campo
		switch revisao.Campo {
		case "F1_ZOBSREV":
			revisao.Status = "Revisar"
		case "XX":
			revisao.Status = "Criado"
		default:
			revisao.Status = "Desconhecido"
		}

		// Organiza as revisões por chave (chave == rec)
		processed[revisao.Chave] = append(processed[revisao.Chave], revisao)
	}

	return processed
}

// ===================
// ENDPOINT HTTP
// ===================

// Função para obter as revisões com base na chave (rec), ou todas se o parâmetro não for fornecido
func GetRevisaoByRec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Pega o parâmetro 'rec' da URL
	rec := r.URL.Query().Get("rec")

	mutex.Lock()
	defer mutex.Unlock()

	// Se o parâmetro 'rec' não for fornecido, retorna todas as revisões agrupadas
	if rec == "" {
		if len(revisaoCache) == 0 {
			log.Println("Nenhuma revisão encontrada no cache.")
			http.Error(w, "Nenhuma revisão encontrada", http.StatusNotFound)
			return
		}

		// Agrupa todas as revisões por `rec`
		result := groupRevisoesByRec(revisaoCache)
		log.Printf("Retornando todas as revisões agrupadas: %v", result)
		json.NewEncoder(w).Encode(result)
		return
	}

	// Se o parâmetro 'rec' for fornecido, retorna as revisões filtradas por rec
	revisoes, found := revisaoCache[rec]
	if !found || len(revisoes) == 0 {
		log.Printf("Nenhuma revisão encontrada para o rec: %s", rec)
		http.Error(w, "Nenhuma revisão encontrada para este rec", http.StatusNotFound)
		return
	}

	// Retorna as revisões específicas para o rec fornecido, agrupadas
	result := RevisaoAgrupada{
		Rec:       rec,
		Historico: revisoes,
	}
	log.Printf("Retornando revisões para rec: %s", rec)
	json.NewEncoder(w).Encode(result)
}

// Função para agrupar revisões por `rec`
func groupRevisoesByRec(revisoes map[string][]Revisao) []RevisaoAgrupada {
	var grouped []RevisaoAgrupada
	for rec, historico := range revisoes {
		grouped = append(grouped, RevisaoAgrupada{
			Rec:       rec,
			Historico: historico,
		})
	}
	return grouped
}
