import { pedidosDetalhadosStore } from '$lib/stores/xmlStore';
import type { Pedido, PedidoProduto } from '$lib/types/Pedidos';

// Função para realizar o fetch e salvar os dados na store
export async function fetchPedidos(cnpj: string) {
	const myHeaders = new Headers();
	myHeaders.append('Content-Type', 'application/json');
	myHeaders.append('CNPJ', cnpj); // Usando o CNPJ no header

	const requestOptions: RequestInit = {
		method: 'GET',
		headers: myHeaders,
		redirect: 'follow' as RequestRedirect
	};

	try {
		const response = await fetch('http://rodoapp:8080/api/pedido', requestOptions);

		if (!response.ok) {
			throw new Error(`Erro ao buscar pedidos: ${response.statusText}`);
		}

		const data = await response.json();

		if (data.length > 0 && data[0].pedidos) {
			// Ordenar os pedidos por código alfabético (campo pedido.pedido)
			const pedidosOrdenados = data[0].pedidos.sort((a: Pedido, b: Pedido) => {
				return a.pedido.localeCompare(b.pedido); // Ordena alfabeticamente pelo código do pedido
			});

			// Transformar os pedidos na estrutura de objeto com o número do pedido como chave
			const pedidosTransformados = pedidosOrdenados.reduce((acc, pedido: Pedido) => {
				const produtosTransformados = pedido.produtos.map((produto: PedidoProduto) => ({
					index: produto.index,
					codigo: produto.codigo,
					nomeProduto: produto.produto,
					quantidade: produto.quantidade,
					valorUnitario: produto.valor_unitario,
					total: produto.total,
					ncm: produto.ncm,
					unidadeMedida: produto.unidade_medida,
					grupo: produto.grupo,
					origem: produto.origem
				}));

				// Obter a condição de pagamento do primeiro produto (assumindo que todos compartilham a mesma)
				const condicaoPagamento = pedido.produtos[0]?.condicao_pagamento || 'N/A';

				// Adiciona cada pedido no objeto com o número do pedido como chave
				acc[pedido.pedido] = {
					pedido: pedido.pedido,
					status: pedido.status,
					pagamento: condicaoPagamento, // Ajustado para ser diretamente no objeto do pedido
					produtos: produtosTransformados
				};

				return acc;
			}, {});

			console.log('Pedidos transformados:', pedidosTransformados);

			// Atualizando a store com os pedidos detalhados transformados (usando o número do pedido como chave)
			pedidosDetalhadosStore.set(pedidosTransformados);
		} else {
			console.log('Nenhum pedido encontrado ou formato de resposta inesperado.');
		}
	} catch (error) {
		console.error('Erro ao buscar pedidos:', error);
	}
}
