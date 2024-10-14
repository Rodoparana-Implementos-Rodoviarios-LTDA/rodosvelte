export interface PedidoProduto {
	index: string; // Índice do produto no pedido
	codigo: string; // Código do produto
	produto: string; // Descrição do produto
	quantidade: number; // Quantidade do produto
	valor_unitario: number; // Valor unitário
	total: number; // Valor total
	condicao_pagamento: string; // Condição de pagamento
	ncm: string; // NCM
	unidade_medida: string; // Unidade de medida
	grupo: string; // Grupo de produtos
	origem: string; // Origem do produto
}

export interface Pedido {
	pedido: string; // Número do pedido
	status: string; // Status do pedido (aprovado, pendente, etc.)
	produtos: PedidoProduto[]; // Lista de produtos dentro de cada pedido
}

export interface PedidoOption {
	label: string; // Exibição no Combobox (Número do pedido)
	value: string; // Valor (Número do pedido)
	produto: string; // Nome do produto
	status: string; // Status do pedido
}
