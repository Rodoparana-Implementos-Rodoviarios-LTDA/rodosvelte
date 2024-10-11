import { pedidosDetalhadosStore, pedidoComboStore } from '$lib/stores/xmlStore';
import type { Pedido, PedidoProduto, PedidoOption } from '$lib/types/Pedidos';

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

			// Salvando os pedidos detalhados completos na store 'pedidosDetalhadosStore'
			pedidosDetalhadosStore.set(pedidosOrdenados);

			// Preparando os dados para o combobox (somente pedido, primeiro produto e status)
			const pedidosCombo: PedidoOption[] = pedidosOrdenados.map((pedido: Pedido) => {
				const primeiroProduto: PedidoProduto = pedido.produtos[0]; // Pega o primeiro produto da lista

				return {
					label: pedido.pedido, // Exibe o número do pedido no Combobox
					value: pedido.pedido, // Valor que será usado internamente
					produto: primeiroProduto.produto, // Nome do primeiro produto
					status: pedido.status // Status do pedido
				};
			});

			// Atualizando a store 'pedidoComboStore' com os dados formatados para o combobox
			pedidoComboStore.set(pedidosCombo);
		} else {
			console.log('Nenhum pedido encontrado ou formato de resposta inesperado.');
			pedidosDetalhadosStore.set([]); // Limpar a store se não houver pedidos
			pedidoComboStore.set([]); // Limpar a store do combobox se não houver pedidos
		}
	} catch (error) {
		console.error('Erro ao buscar pedidos:', error);
		pedidosDetalhadosStore.set([]); // Limpar a store em caso de erro
		pedidoComboStore.set([]); // Limpar a store do combobox em caso de erro
	}
}
