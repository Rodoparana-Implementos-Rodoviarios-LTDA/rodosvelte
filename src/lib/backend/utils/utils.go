// utils/utils.go
package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ===================
// REQUISIÇÕES EXTERNAS
// ===================

// FetchFromEndpoint realiza uma requisição genérica em um endpoint e retorna os dados decodificados
func FetchFromEndpoint[T any](url string, headers map[string]string) ([]T, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %v", err)
	}

	// Adiciona os headers (se houver)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada: %d %s", resp.StatusCode, resp.Status)
	}

	var result []T
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %v", err)
	}

	return result, nil
}

//==============================================\\==================================================

// ===================
// PARÂMETROS E ORGANIZAÇÃO
// ===================

// === PAGINAÇÃO ===
// GetPaginationParams pega os parâmetros de paginação da URL
func GetPaginationParams(r *http.Request) (int, int) {
	page := 1
	pageSize := 10

	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("pageSize")

	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}
	if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 {
		pageSize = ps
	}

	return page, pageSize
}

// Paginate aplica a paginação a uma lista de itens genérica
func Paginate[T any](items []T, page, pageSize int) []T {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(items) {
		return []T{}
	}
	if end > len(items) {
		end = len(items)
	}

	return items[start:end]
}

// === ORDENAÇÃO ===
// GetSortingParams pega os parâmetros de ordenação da URL e valida com base em campos permitidos
func GetSortingParams(r *http.Request, validSortFields map[string]bool) (string, string) {
	sortBy := r.URL.Query().Get("sortBy")
	if _, valid := validSortFields[sortBy]; !valid {
		sortBy = "" // Campo padrão ou vazio se inválido
	}

	sortOrder := r.URL.Query().Get("sortOrder")
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc" // Ordenação padrão
	}

	return sortBy, sortOrder
}

// SortData ordena os dados com base em um campo específico e a direção da ordenação
func SortData[T any](items []T, sortBy, sortOrder string, compareFunc func(i, j T) bool) []T {
	sort.SliceStable(items, func(i, j int) bool {
		if sortOrder == "asc" {
			return compareFunc(items[i], items[j])
		}
		return compareFunc(items[j], items[i])
	})
	return items
}

// === FILTRAGEM ===
// GetFilters pega os filtros da URL e os converte em um mapa chave-valor
func GetFilters(r *http.Request) map[string]string {
	filters := make(map[string]string)

	// Captura todos os parâmetros de query
	for key, values := range r.URL.Query() {
		filters[key] = values[0] // Pega apenas o primeiro valor
	}

	return filters
}

// FilterData aplica filtros dinâmicos a uma lista de dados
func FilterData[T any](items []T, filters map[string]string, compareFunc func(item T, key, value string) bool) []T {
	var filtered []T

	for _, item := range items {
		matchesAll := true
		for key, value := range filters {
			if !compareFunc(item, key, value) {
				matchesAll = false
				break
			}
		}

		if matchesAll {
			filtered = append(filtered, item)
		}
	}

	return filtered
}

//==============================================\\==================================================

// ===================
// FORMATAÇÃO DE DADOS
// ===================

// FormatCurrency formata valores monetários no formato brasileiro (Real)
func FormatCurrency(valueStr string) string {
	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		return valueStr
	}
	return fmt.Sprintf("%.2f", value)
}

// FormatDate formata uma string de data de entrada no formato `fromFormat` para o formato `toFormat`
// Exemplo de uso: FormatDate("20220609", "20060102", "02/01/2006") -> "09/06/2022"
func FormatDate(dateStr, fromFormat, toFormat string) string {
	// Remove qualquer espaço adicional na data
	dateStr = strings.TrimSpace(dateStr)

	// Parse a data no formato de entrada (fromFormat)
	parsedDate, err := time.Parse(fromFormat, dateStr)
	if err != nil {
		log.Printf("Erro ao fazer o parsing da data: %v", err)
		return dateStr // Retorna a data original em caso de erro
	}

	// Retorna a data formatada no formato desejado (toFormat)
	return parsedDate.Format(toFormat)
}

// TrimString remove espaços em branco de uma string.
func TrimString(input string) string {
	return strings.TrimSpace(input)
}

// DetermineStatus determina o status com base em Status e Revisão.
func DetermineStatus(status, revisao string) string {
	if TrimString(revisao) != "" {
		return "Revisar"
	}
	if TrimString(status) != "" {
		return "Classificado"
	}
	return "Pendente"
}

// Mapear valores de Z08_ORIGEM para descrições legíveis
func MapOrigem(origem string) string {
	switch origem {
	case "S":
		return "Saída"
	case "E":
		return "Entrada"
	default:
		return origem
	}
}

// Mapear valores de Z08_TIPMOV para descrições legíveis
func MapTipoMovimento(tipoMov string) string {
	switch tipoMov {
	case "1":
		return "Venda"
	case "2":
		return "Carregamento"
	case "3":
		return "Saída"
	default:
		return tipoMov
	}
}
