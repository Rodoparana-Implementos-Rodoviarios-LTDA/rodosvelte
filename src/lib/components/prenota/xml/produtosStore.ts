import { get } from 'svelte/store';
import { produtosStore, selectedTipoNF } from '$lib/stores/xmlStore';
import type { ComboboxOption } from '$lib/types/Produtos';

/**
 * Formata os produtos do Protheus para exibição no Combobox.
 */
export function formatarProdutosProtheus(produtos: any[]): ComboboxOption[] {
	return produtos.map((produto: any) => ({
		label: `${produto.B1_COD} - ${produto.B1_DESC}`,
		value: produto.B1_COD,
		campo1: produto.B1_POSIPI, // NCM
		campo2: produto.B1_ORIGEM // Origem
	}));
}

/**
 * Aplica o filtro de NCM e origem nos produtos do Protheus, dependendo do tipo de NF.
 */
export function aplicarFiltroProtheus(
	produtos: ComboboxOption[],
	ncmsh: string,
	origem: string,
	tipoNF: string
): ComboboxOption[] {
	// Apenas para tipos de NF "Matéria Prima" ou "Revenda", aplicamos os filtros de NCM e origem
	if (tipoNF === 'Matéria Prima' || tipoNF === 'Revenda') {
		return produtos.filter((produto) => {
			const ncmshValue = ncmsh.trim();
			const origemValue = origem.trim();
			const produtoNcm = produto.campo1.trim();
			const produtoOrigem = produto.campo2.trim();

			// Filtro baseado em NCM e origem
			const ncmMatches = produtoNcm === ncmshValue;
			const origemMatches = produtoOrigem === origemValue;

			return ncmMatches && origemMatches;
		});
	}

	// Se não for "Matéria Prima" ou "Revenda", retorna os produtos sem aplicar filtros
	return produtos;
}

/**
 * Retorna os produtos do Protheus, aplicando a formatação e os filtros de acordo com o tipo de NF, NCM e origem.
 */
export function obterProdutosProtheusFiltrados(ncmsh: string, origem: string): ComboboxOption[] {
	const tipoNF = get(selectedTipoNF); // Obter o tipo de NF selecionado
	const produtosProtheus = get(produtosStore); // Obter os produtos do Protheus da store

	// Formatar os produtos do Protheus
	let produtosFormatados = formatarProdutosProtheus(produtosProtheus);

	// Aplicar filtro de NCM e origem baseado no tipo de NF
	produtosFormatados = aplicarFiltroProtheus(produtosFormatados, ncmsh, origem, tipoNF);

	return produtosFormatados;
}
