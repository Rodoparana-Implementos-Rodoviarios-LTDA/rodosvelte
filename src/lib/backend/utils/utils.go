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
// O campo é passado como uma string representando o nome da coluna, e o sortOrder define se é ascendente ("asc") ou descendente ("desc").
func SortByColumn[T any](data []T, column string, sortOrder string) []T {
	sort.SliceStable(data, func(i, j int) bool {
		vi := reflect.ValueOf(data[i]).FieldByName(column)
		vj := reflect.ValueOf(data[j]).FieldByName(column)

		// Se o campo não existir, não faz nada
		if !vi.IsValid() || !vj.IsValid() {
			return false
		}

		// Comparação genérica baseada no tipo do campo
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

