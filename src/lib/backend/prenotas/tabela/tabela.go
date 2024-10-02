package tabela

import (
	"backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// ===================
// VARIAVEIS GLOBAIS
// ===================
var (
	preNotasCache    []PreNota  // Cache dos dados das pre-notas
	ultimoRec        string     // Armazena o maior REC no cache
	mutex            sync.Mutex // Controle de concorrencia para o cache
	cacheInitialized bool       // Flag para garantir que o cache foi inicializado ao menos uma vez
)

// ===================
// ESTRUTURAS DE DADOS
// ===================

type PreNota struct {
	Filial, NF, Status, Fornecedor, Emissao, Inclusao, Vencimento, Valor,
	Tipo, Prioridade, Usuario, Obs, ObsRev, Revisao, Rec string
}

type RawPreNota struct {
	Filial     string `json:"F1_FILIAL"`
	Doc        string `json:"F1_DOC"`
	Serie      string `json:"F1_SERIE"`
	Status     string `json:"F1_STATUS"`
	Fornecedor string `json:"FORNECE"`
	Emissao    string `json:"F1_EMISSAO"`
	Inclusao   string `json:"F1_DTDIGIT"`
	Vencimento string `json:"VENCIMENTO"`
	ValBrut    string `json:"F1_VALBRUT"`
	Tipo       string `json:"F1_XTIPO"`
	Prior      string `json:"F1_XPRIOR"`
	UsrRa      string `json:"F1_XUSRRA"`
	Obs        string `json:"F1_XOBS"`
	ObsRev     string `json:"F1_ZOBSREV"`
	Rev        string `json:"F1_XREV"`
	Rec        string `json:"REC"`
}

// ===================
// REQUISICAO E CACHE
// ===================

// InitializeCache faz o carregamento inicial dos primeiros 200 registros e depois o restante
func InitializeCache() {
	log.Println("Carregando dados iniciais de pre-notas (primeiros 200)...")
	fetchPreNotasDataInicial(200) // Carregando inicialmente os primeiros 200 registros

	// Inicia o carregamento do restante dos registros de forma assincrona
	go fetchPreNotasRestante()

	// Inicia o polling para buscar novos registros incrementais
	go startPolling()
}

// Carrega os primeiros registros do sistema para um carregamento rapido
func fetchPreNotasDataInicial(limit int) {
	mutex.Lock()
	defer mutex.Unlock()

	// Faz a requisicao dos primeiros "limit" registros
	preNotas, err := utils.FetchFromEndpoint[RawPreNota](fmt.Sprintf("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=%d", limit), nil)
	if err != nil {
		log.Printf("Erro ao buscar pre-notas iniciais: %v", err)
		return
	}

	// Processa os dados e atualiza o cache
	preNotasCache = processPreNotas(preNotas)

	// Atualiza o maior REC do cache
	if len(preNotasCache) > 0 {
		ultimoRec = preNotasCache[len(preNotasCache)-1].Rec
	}
	cacheInitialized = true
	log.Println("Cache inicial de pre-notas (parcial) carregado com sucesso!")
}

// Carrega o restante dos registros enquanto o sistema esta rodando
func fetchPreNotasRestante() {
	log.Println("Carregando restante das pre-notas...")

	mutex.Lock()
	defer mutex.Unlock()

	// Busca o restante dos registros
	preNotasRestantes, err := utils.FetchFromEndpoint[RawPreNota]("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=1&numItem=999999", nil)
	if err != nil {
		log.Printf("Erro ao buscar o restante das pre-notas: %v", err)
		return
	}

	// Processa os registros restantes e adiciona ao cache
	for _, preNota := range preNotasRestantes {
		if preNota.Rec > ultimoRec {
			preNotasCache = append(preNotasCache, processPreNota(preNota))
			ultimoRec = preNota.Rec
		}
	}
	log.Println("Restante das pre-notas carregado com sucesso!")
}

// Polling para buscar novas pre-notas a cada minuto
func startPolling() {
	for {
		time.Sleep(1 * time.Minute)
		fetchPreNotasIncremental()
	}
}

// Funcao para buscar apenas as pre-notas mais novas com base no ultimo REC
func fetchPreNotasIncremental() {
	mutex.Lock()
	defer mutex.Unlock()

	if !cacheInitialized {
		log.Println("Cache nao inicializado ainda. Polling abortado.")
		return
	}

	// Faz uma requisicao para buscar os ultimos 50 registros mais recentes
	preNotasNovas, err := utils.FetchFromEndpoint[RawPreNota]("http://172.16.99.174:8400/rest/PreNota/ListaPreNota?pag=0&numItem=50", nil)
	if err != nil {
		log.Printf("Erro ao buscar pre-notas incrementais: %v", err)
		return
	}

	// Adiciona apenas pre-notas com REC maior que o ultimo registrado
	for _, preNota := range preNotasNovas {
		if preNota.Rec > ultimoRec {
			preNotasCache = append(preNotasCache, processPreNota(preNota))
			ultimoRec = preNota.Rec
		}
	}

	log.Println("Cache de pre-notas atualizado com novos registros.")
}

// ===================
// HANDLER PRINCIPAL
// ===================

func GetPreNotas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mutex.Lock()
	defer mutex.Unlock()

	// Verifica se o cache foi inicializado
	if !cacheInitialized || len(preNotasCache) == 0 {
		http.Error(w, "Nenhuma pre-nota disponivel.", http.StatusNotFound)
		return
	}

	// Aplica filtros e paginacao
	filterableColumns := []string{"Filial", "NF", "Status", "Fornecedor", "Emissao", "Inclusao", "Vencimento", "Valor", "Tipo", "Prioridade", "Usuario", "Obs", "ObsRev", "Rec"}
	filteredPreNotas := applyFilters(preNotasCache, r, filterableColumns)
	sortedPreNotas := applySorting(filteredPreNotas, r)
	paginatedPreNotas := utils.Paginate(sortedPreNotas, r)

	// Retorna os dados filtrados e paginados
	err := json.NewEncoder(w).Encode(paginatedPreNotas)
	if err != nil {
		http.Error(w, "Erro ao gerar a resposta.", http.StatusInternalServerError)
	}
}

// ===================
// FUNCOES AUXILIARES
// ===================

// Funcao para aplicar os filtros nos dados com base nos headers.
func applyFilters(preNotas []PreNota, r *http.Request, filterableColumns []string) []PreNota {
	return utils.FilterData(preNotas, r, filterableColumns)
}

// Funcao para aplicar a classificacao com base nos headers.
func applySorting(preNotas []PreNota, r *http.Request) []PreNota {
	sortBy := r.Header.Get("X-Sort-By")
	sortOrder := r.Header.Get("X-Sort-Order")

	if sortBy == "" {
		sortBy = "Inclusao"
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	return utils.SortByColumn(preNotas, sortBy, sortOrder)
}

// Funcao que processa uma lista de pre-notas e converte para a estrutura final
func processPreNotas(rawPreNotas []RawPreNota) []PreNota {
	var processed []PreNota
	for _, raw := range rawPreNotas {
		processed = append(processed, processPreNota(raw))
	}
	return processed
}

// Funcao que processa uma unica pre-nota bruta e a converte para o formato final
func processPreNota(raw RawPreNota) PreNota {
	nf := utils.TrimString(raw.Doc)
	serie := utils.TrimString(raw.Serie)
	nfCompleta := nf
	if serie != "" {
		nfCompleta += " - " + serie
	}

	emissao := utils.FormatDate(utils.TrimString(raw.Emissao), "20060102", "02/01/2006")
	inclusao := utils.FormatDate(utils.TrimString(raw.Inclusao), "20060102", "02/01/2006")
	vencimento := utils.FormatDate(utils.TrimString(raw.Vencimento), "20060102", "02/01/2006")

	return PreNota{
		Filial:     utils.TrimString(raw.Filial),
		NF:         nfCompleta,
		Status:     utils.DetermineStatus(raw.Status, raw.Rev),
		Fornecedor: utils.TrimString(raw.Fornecedor),
		Emissao:    emissao,
		Inclusao:   inclusao,
		Vencimento: vencimento,
		Valor:      utils.FormatCurrency(utils.TrimString(raw.ValBrut)),
		Tipo:       utils.TrimString(raw.Tipo),
		Prioridade: utils.TrimString(raw.Prior),
		Usuario:    utils.TrimString(raw.UsrRa),
		Obs:        utils.TrimString(raw.Obs),
		ObsRev:     utils.TrimString(raw.ObsRev),
		Revisao:    utils.TrimString(raw.Rev),
		Rec:        utils.TrimString(raw.Rec),
	}
}
