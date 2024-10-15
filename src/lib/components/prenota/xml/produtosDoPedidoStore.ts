import { get } from 'svelte/store';
import { produtosDoPedidoStore, selectedTipoNF } from '$lib/stores/xmlStore';
import type { ComboboxOption } from '$lib/types/Produtos';

/**
 * Aplica os filtros conforme o tipo de NF, NCM, origem e outros parâmetros
 */
export function aplicarFiltrosProdutosDoPedido(
	ncmsh: string,
	origem: string,
	quantidade: number,
	valorUnitario: number,
	valorTotal: number
): ComboboxOption[] {
	const tipoNF = get(selectedTipoNF); // Obter o tipo de NF selecionado
	const produtosDoPedido = get(produtosDoPedidoStore); // Obter os produtos do pedido

	// Filtrar os produtos de acordo com as condições do tipo de NF
	return produtosDoPedido.filter((produto: any) => {
		let ncmMatches = true;
		let origemMatches = true;
		let quantidadeMatches = true;
		let valorUnitarioMatches = true;
		let valorTotalMatches = true;

		// Revenda e Matéria Prima: devem bater NCM e Origem
		if (tipoNF === 'Revenda' || tipoNF === 'Matéria Prima') {
			ncmMatches = produto.campo1.trim() === ncmsh.trim();
			origemMatches = produto.campo2.trim() === origem.trim();
		}

		// Comparar a quantidade
		if (quantidade) {
			quantidadeMatches = produto.quantidade === quantidade;
		}

		// Comparar o valor unitário
		if (valorUnitario) {
			valorUnitarioMatches = produto.valorUnitario === valorUnitario;
		}

		// Comparar o valor total
		if (valorTotal) {
			valorTotalMatches = produto.valorTotal === valorTotal;
		}

		// Retornar verdadeiro se todos os critérios forem atendidos
		return (
			ncmMatches && origemMatches && quantidadeMatches && valorUnitarioMatches && valorTotalMatches
		);
	});
}

/**
 * Formata os produtos filtrados para exibição no Combobox.
 */
export function formatarProdutosDoPedidoFiltrados(
	ncmsh: string,
	origem: string,
	quantidade: number,
	valorUnitario: number,
	valorTotal: number
): ComboboxOption[] {
	const produtosFiltrados = aplicarFiltrosProdutosDoPedido(
		ncmsh,
		origem,
		quantidade,
		valorUnitario,
		valorTotal
	);

	// Formatar os produtos para o Combobox
	return produtosFiltrados.map((produto) => ({
		label: `${produto.codigo} - ${produto.nomeProduto}`,
		value: produto.codigo,
		campo1: produto.quantidade.toString(), // Quantidade
		campo2: produto.total.toFixed(2) // Valor total
	}));
}
