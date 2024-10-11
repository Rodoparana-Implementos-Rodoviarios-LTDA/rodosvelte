package pedido

import (
	"backend/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Estrutura do pedido (fornecedor)
type Pedido struct {
	Fornecedor string  `json:"fornecedor"`
	CNPJ       string  `json:"cnpj"`
	Pedidos    []Order `json:"pedidos"` // Lista de pedidos agrupados
}

// Estrutura para o agrupamento de itens por pedido
type Order struct {
	Pedido   string `json:"pedido"`
	Status   string `json:"status"`
	Produtos []Item `json:"produtos"`
}

// Estrutura dos itens do pedido
type Item struct {
	Index             string  `json:"index"`
	Codigo            string  `json:"codigo"`
	Produto           string  `json:"produto"`
	Quantidade        int     `json:"quantidade"`
	ValorUnitario     float64 `json:"valor_unitario"`
	Total             float64 `json:"total"`
	CondicaoPagamento string  `json:"condicao_pagamento"`
	NCM               string  `json:"ncm"`
	UnidadeMedida     string  `json:"unidade_medida"`
	Grupo             string  `json:"grupo"`
	Origem            string  `json:"origem"`
}

// GetPedidos é o handler que responde à rota /api/pedido
func GetPedidos(w http.ResponseWriter, r *http.Request) {
	// Verifica se o CNPJ foi enviado via header
	cnpj := r.Header.Get("CNPJ")
	if cnpj == "" {
		http.Error(w, "CNPJ não fornecido", http.StatusBadRequest)
		return
	}

	// Chama a função FetchPedido para obter os dados
	pedidos, err := FetchPedido(cnpj)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao buscar pedidos: %v", err), http.StatusInternalServerError)
		return
	}

	// Retorna os pedidos em formato JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pedidos); err != nil {
		http.Error(w, "Erro ao gerar a resposta", http.StatusInternalServerError)
	}
}

// FetchPedido realiza a requisição ao endpoint para buscar os pedidos com base no CNPJ
func FetchPedido(cnpj string) ([]Pedido, error) {
	url := "http://protheus-app:8400/rest/reidoapsdu/consultar/likefor"

	// Headers para a requisição, incluindo o CNPJ
	headers := map[string]string{
		"Content-Type": "application/json",
		"busca":        cnpj,
	}

	// Usa a função FetchRawFromEndpoint da utils para fazer a requisição
	rawResponse, err := utils.FetchRawFromEndpoint(url, headers)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pedidos: %w", err)
	}

	// Limpa o JSON para remover caracteres indesejados (se necessário)
	cleanedResponse := cleanJSON(rawResponse)

	// Parseia o JSON para a estrutura Pedido
	var pedidosRaw []map[string]interface{}
	if err := json.Unmarshal([]byte(cleanedResponse), &pedidosRaw); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	// Processa e transforma os dados
	pedidos := transformPedidos(pedidosRaw)

	return pedidos, nil
}

// Transforma os dados do JSON original para o formato esperado, agrupando por número de pedido
func transformPedidos(pedidosRaw []map[string]interface{}) []Pedido {
	var pedidos []Pedido

	for _, pedidoRaw := range pedidosRaw {
		// Extrai e renomeia os campos do fornecedor
		fornecedor := Pedido{
			Fornecedor: getString(pedidoRaw, "APELIDO"),
			CNPJ:       getString(pedidoRaw, "A2_CGC"),
		}

		// Agrupamento dos pedidos pelo campo "C7_NUM"
		pedidoMap := make(map[string]*Order)

		// Processa a lista de itens dentro de "PEDIDOS"
		if itensRaw, ok := pedidoRaw["PEDIDOS"].([]interface{}); ok {
			for _, itemRaw := range itensRaw {
				if itemMap, ok := itemRaw.(map[string]interface{}); ok {
					pedidoNum := getString(itemMap, "C7_NUM")
					status := strings.ReplaceAll(getString(itemMap, "STATUS"), pedidoNum+" ", "")

					// Verifica se já existe um pedido com o mesmo número
					if _, exists := pedidoMap[pedidoNum]; !exists {
						// Cria uma nova entrada para o pedido, se ainda não existe
						pedidoMap[pedidoNum] = &Order{
							Pedido:   pedidoNum,
							Status:   status,
							Produtos: []Item{},
						}
					}

					// Cria o item do pedido
					item := Item{
						Index:             getString(itemMap, "C7_ITEM"),
						Codigo:            getString(itemMap, "C7_PRODUTO"),
						Produto:           getString(itemMap, "B1_DESC"),
						Quantidade:        getInt(itemMap, "C7_QUANT"),
						ValorUnitario:     getFloat(itemMap, "C7_PRECO"),
						Total:             getFloat(itemMap, "C7_TOTAL"),
						CondicaoPagamento: getString(itemMap, "C7_COND"),
						NCM:               getString(itemMap, "B1_POSIPI"),
						UnidadeMedida:     getString(itemMap, "B1_UM"),
						Grupo:             getString(itemMap, "B1_GRUPO"),
						Origem:            getString(itemMap, "B1_ORIGEM"),
					}

					// Adiciona o item à lista de produtos do pedido
					pedidoMap[pedidoNum].Produtos = append(pedidoMap[pedidoNum].Produtos, item)
				}
			}
		}

		// Converte o mapa de pedidos agrupados em uma lista
		for _, order := range pedidoMap {
			fornecedor.Pedidos = append(fornecedor.Pedidos, *order)
		}

		// Adiciona o fornecedor com seus pedidos à lista final
		pedidos = append(pedidos, fornecedor)
	}

	return pedidos
}

// Função auxiliar para limpar o JSON (caso tenha conteúdo não desejado)
func cleanJSON(raw string) string {
	return strings.ReplaceAll(strings.ReplaceAll(raw, "\n", ""), "\t", "")
}

// Funções auxiliares para extrair os dados corretamente
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key].(string); ok {
		return val
	}
	return ""
}

func getInt(data map[string]interface{}, key string) int {
	if val, ok := data[key].(float64); ok {
		return int(val)
	}
	return 0
}

func getFloat(data map[string]interface{}, key string) float64 {
	if val, ok := data[key].(float64); ok {
		return val
	}
	return 0.0
}
