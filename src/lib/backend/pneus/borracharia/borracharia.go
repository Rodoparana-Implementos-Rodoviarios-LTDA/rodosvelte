package borracharia

import (
	"backend/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

// ===================
// VARIAVEIS GLOBAIS
// ===================
var (
	borrachariaCache []Borracharia
	cacheTimestamp   time.Time
	cacheDuration    = 1 * time.Hour
	mutex            sync.Mutex
	cacheInitialized bool
	cacheUpdating    bool
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type Borracharia struct {
	Filial, NF, Vendedor, Cliente, Emissao string
}

type RawBorracharia struct {
	Filial       string `json:"F2_FILIAL"`
	Documento    string `json:"F2_DOC"`
	Serie        string `json:"F2_SERIE"`
	Emissao      string `json:"F2_EMISSAO"`
	Vendedor     string `json:"F2_VEND1"`
	VendedorNome string `json:"A3_NOME"`
	Cliente      string `json:"F2_CLIENTE"`
	Loja         string `json:"F2_LOJA"`
	ClienteNome  string `json:"A1_NOME"`
}

// ===================
// REQUISICAO E CACHE
// ===================

func InitializeCache() {
	log.Println("Carregando dados iniciais de borracharia...")
	fetchBorrachariaData()
}

func fetchBorrachariaData() {
	mutex.Lock()

	if cacheUpdating {
		log.Println("Atualizacao do cache de borracharia ja esta em andamento, aguardando...")
		mutex.Unlock()
		return
	}

	cacheUpdating = true
	mutex.Unlock()

	resp, err := http.Get("http://protheus-vm:9010/rest/MovPortaria/NFSaidasDisponiveis")
	if err != nil {
		log.Printf("Erro ao buscar dados de borracharia: %v", err)
		mutex.Lock()
		cacheUpdating = false
		mutex.Unlock()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta da API: %v", err)
		mutex.Lock()
		cacheUpdating = false
		mutex.Unlock()
		return
	}

	var rawBorracharia []RawBorracharia
	if err := json.Unmarshal(body, &rawBorracharia); err != nil {
		log.Printf("Erro ao decodificar JSON da borracharia: %v", err)
		mutex.Lock()
		cacheUpdating = false
		mutex.Unlock()
		return
	}

	mutex.Lock()
	borrachariaCache = processBorracharia(rawBorracharia)
	cacheTimestamp = time.Now()
	cacheInitialized = true
	cacheUpdating = false
	mutex.Unlock()

	log.Println("Cache de borracharia atualizado com sucesso!")
}

// ===================
// HANDLER HTTP
// ===================

func GetBorracharia(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	if !cacheInitialized || len(borrachariaCache) == 0 || time.Since(cacheTimestamp) > cacheDuration {
		log.Println("Cache de borracharia expirado ou vazio. Carregando novos dados em background...")
		go fetchBorrachariaData()
	}

	filterableColumns := []string{"Filial", "NF", "Vendedor", "Cliente", "Emissao"}
	filteredBorracharia := applyFilters(borrachariaCache, r, filterableColumns)

	if len(filteredBorracharia) == 0 {
		mutex.Unlock()
		http.Error(w, "Nenhuma informacao de borracharia disponivel.", http.StatusNotFound)
		return
	}

	sortedBorracharia := applySorting(filteredBorracharia, r)
	paginatedBorracharia := utils.Paginate(sortedBorracharia, r)

	mutex.Unlock()

	if err := json.NewEncoder(w).Encode(paginatedBorracharia); err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

func applyFilters(borracharia []Borracharia, r *http.Request, filterableColumns []string) []Borracharia {
	return utils.FilterData(borracharia, r, filterableColumns)
}

func applySorting(borracharia []Borracharia, r *http.Request) []Borracharia {
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	if sortBy == "" {
		sortBy = "Emissao"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	return utils.SortByColumn(borracharia, sortBy, sortOrder)
}

func processBorracharia(rawBorracharia []RawBorracharia) []Borracharia {
	var processed []Borracharia
	for _, raw := range rawBorracharia {
		borracharia := Borracharia{
			Filial:   utils.TrimString(raw.Filial),
			NF:       utils.TrimString(raw.Documento) + " - " + utils.TrimString(raw.Serie),
			Vendedor: utils.TrimString(raw.Vendedor) + " - " + utils.TrimString(raw.VendedorNome),
			Cliente:  utils.TrimString(raw.Cliente) + " - " + utils.TrimString(raw.Loja) + " - " + utils.TrimString(raw.ClienteNome),
			Emissao:  utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006"),
		}
		processed = append(processed, borracharia)
	}
	return processed
}
