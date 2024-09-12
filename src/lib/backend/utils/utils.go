package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ===================
// REQUISIÇÕES EXTERNAS
// ===================

// FetchFromEndpoint realiza uma requisição HTTP genérica e decodifica a resposta em um slice do tipo T.
// Uso de generics para lidar com qualquer estrutura de dados.
func FetchFromEndpoint[T any](url string, headers map[string]string) ([]T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	// Adiciona headers se fornecidos
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req) // Usando DefaultClient ao invés de criar um novo cliente
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resposta inesperada: %d %s", resp.StatusCode, resp.Status)
	}

	var result []T
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return result, nil
}

// ===================
// FORMATAÇÃO DE DADOS
// ===================

// FormatCurrency converte uma string de valor monetário para o formato brasileiro.
// Melhorias: Verificação simplificada e retorno do valor formatado com precisão definida.
func FormatCurrency(valueStr string) string {
	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		return valueStr // Retorna o valor original se não puder ser convertido
	}
	return fmt.Sprintf("%.2f", value) // Mantém 2 casas decimais para valores monetários
}

// FormatDate transforma uma string de data em um formato de origem para outro formato alvo.
// Melhorias: Simplificação da lógica de erro e mensagens de log mais claras.
func FormatDate(dateStr, fromFormat, toFormat string) string {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return "" // Retorna vazio se a string da data for vazia
	}

	parsedDate, err := time.Parse(fromFormat, dateStr)
	if err != nil {
		log.Printf("Erro ao fazer o parsing da data (%s): %v", dateStr, err)
		return dateStr // Retorna a data original em caso de erro
	}

	return parsedDate.Format(toFormat)
}

// TrimString remove espaços em branco de uma string.
// Simplesmente invoca strings.TrimSpace para clareza e consistência.
func TrimString(input string) string {
	return strings.TrimSpace(input)
}

// DetermineStatus define o status com base em status e revisão.
// Refatoração: Tornar a lógica mais clara e remover redundâncias.
func DetermineStatus(status, revisao string) string {
	switch {
	case TrimString(revisao) != "":
		return "Revisar"
	case TrimString(status) != "":
		return "Classificado"
	default:
		return "Pendente"
	}
}

// MapOrigem transforma um código de origem em uma descrição legível.
// Refatoração: Código mais compacto com switch simplificado.
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

// MapTipoMovimento converte um código de tipo de movimento para uma descrição legível.
// Refatoração: Compactação do switch.
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

// ===================
// PARÂMETROS DE GET
// ===================

// SortByColumn classifica uma lista de qualquer tipo com base em um campo específico.
func SortByColumn[T any](data []T, column string, sortOrder string) []T {
	sort.SliceStable(data, func(i, j int) bool {
		vi := reflect.ValueOf(data[i]).FieldByName(column)
		vj := reflect.ValueOf(data[j]).FieldByName(column)

		// Se o campo for uma data, converte para time.Time e compara
		if column == "Inclusao" || column == "Emissao" || column == "Vencimento" {
			timeFormat := "02/01/2006"
			dateI, errI := time.Parse(timeFormat, vi.String())
			dateJ, errJ := time.Parse(timeFormat, vj.String())

			if errI != nil || errJ != nil {
				log.Printf("Erro ao converter string para data: %v, %v", errI, errJ)
				return false
			}

			// Compara as datas de acordo com o sortOrder
			if sortOrder == "desc" {
				return dateI.After(dateJ)
			}
			return dateI.Before(dateJ)
		}

		// Comparação padrão para strings, inteiros e floats
		switch vi.Kind() {
		case reflect.String:
			if sortOrder == "desc" {
				return vi.String() > vj.String()
			}
			return vi.String() < vj.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if sortOrder == "desc" {
				return vi.Int() > vj.Int()
			}
			return vi.Int() < vj.Int()
		case reflect.Float32, reflect.Float64:
			if sortOrder == "desc" {
				return vi.Float() > vj.Float()
			}
			return vi.Float() < vj.Float()
		default:
			return false
		}
	})
	return data
}
// FilterData filtra uma lista de qualquer tipo com base nos headers enviados. 
// Para cada coluna especificada, ele compara o valor no header com o valor no dado correspondente.
func FilterData[T any](data []T, r *http.Request, filterableColumns []string) []T {
	var filteredData []T

	for _, item := range data {
		matchesAll := true

		for _, column := range filterableColumns {
			// Pega o valor do header correspondente ao filtro da coluna
			filterValue := r.Header.Get("X-Filter-" + column)

			// Se o filtro não estiver presente ou estiver vazio, ignora essa coluna
			if filterValue != "" {
				// Busca o valor do campo correspondente no item
				fieldValue := reflect.ValueOf(item).FieldByName(column)

				// Verifica se o campo existe e é uma string
				if fieldValue.IsValid() && fieldValue.Kind() == reflect.String {
					// Comparação por substring (case-insensitive)
					if !strings.Contains(strings.ToLower(fieldValue.String()), strings.ToLower(filterValue)) {
						matchesAll = false
						break
					}
				} else {
					matchesAll = false
					break
				}
			}
		}

		// Se o item combinar com todos os filtros, ele é adicionado ao resultado
		if matchesAll {
			filteredData = append(filteredData, item)
		}
	}

	return filteredData
}
// Paginate aplica a paginação a uma lista de itens genéricos, retornando apenas os itens da página solicitada.
func Paginate[T any](items []T, r *http.Request) []T {
	// Define os valores padrão de paginação
	page := 1      // Página padrão
	pageSize := 10 // Número de itens por página padrão

	// Pega os parâmetros de paginação do header
	pageParam := r.Header.Get("X-Page")
	pageSizeParam := r.Header.Get("X-Page-Size")

	// Converte os parâmetros recebidos
	if p, err := strconv.Atoi(pageParam); err == nil && p > 0 {
		page = p
	}
	if ps, err := strconv.Atoi(pageSizeParam); err == nil && ps > 0 {
		pageSize = ps
	}

	// Calcula os índices de início e fim da página
	start := (page - 1) * pageSize
	end := start + pageSize

	// Verifica se os índices estão dentro do limite
	if start >= len(items) {
		return []T{} // Página vazia se o início estiver fora do limite
	}
	if end > len(items) {
		end = len(items) // Ajusta o fim se ultrapassar o número de itens
	}

	// Retorna os itens da página correspondente
	return items[start:end]
}