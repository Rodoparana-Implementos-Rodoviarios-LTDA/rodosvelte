<script lang="ts">
	import { produtosDoPedidoStore, produtosStore, selectedTipoNF } from '$lib/stores/xmlStore';
	import Combobox from '$lib/components/ui/Combobox/Combobox.svelte';
	import { get } from 'svelte/store';

	// Props recebidas da coluna (para NCM e origem)
	export let origem: string;
	export let ncmsh: string;

	// Armazenar os produtos disponíveis no pedido e no Protheus
	let produtosOptions: ComboboxOption[] = [];
	let produtosProtheus = get(produtosStore); // Produtos disponíveis do Protheus

	// Obtenha o tipo de NF (Matéria Prima, Revenda, etc.)
	const tipoNF = get(selectedTipoNF);

	// Função para formatar os produtos
	function FormataProdutos(produtos: Produto[]): ComboboxOption[] {
		return produtos.map((produto) => ({
			label: `${produto.codigo} - ${produto.nomeProduto}`,
			value: produto.codigo,
			campo1: produto.quantidade.toString(),
			campo2: produto.total.toFixed(2)
		}));
	}

	// Função para aplicar o filtro baseado em NCM e origem
	function aplicarFiltro(produtos: ComboboxOption[]): ComboboxOption[] {
		return produtos.filter((produto) => {
			const ncmshValue = String(ncmsh || '').trim();
			const origemValue = String(origem || '').trim();
			const produtoNcm = String(produto.campo1 || '').trim(); // NCM do produto
			const produtoOrigem = String(produto.campo2 || '').trim(); // Origem do produto

			// Aplicar o filtro baseado no NCM e origem
			const ncmMatches = produtoNcm === ncmshValue;
			const origemMatches = produtoOrigem === origemValue;

			return ncmMatches && origemMatches;
		});
	}

	// Lógica para formatar e filtrar os produtos
	$: {
		if ($produtosDoPedidoStore.length > 0) {
			// Formatar os produtos dos pedidos selecionados
			produtosOptions = FormataProdutos($produtosDoPedidoStore);

			// Aplicar o filtro de NCM e origem se for "Revenda" ou "Matéria Prima"
			if (tipoNF === 'Matéria Prima' || tipoNF === 'Revenda') {
				produtosOptions = aplicarFiltro(produtosOptions);
			}
		} else {
			// Caso não haja produtos dos pedidos selecionados, aplicar a lógica do Protheus
			if (tipoNF === 'Matéria Prima' || tipoNF === 'Revenda') {
				produtosProtheus = produtosProtheus.filter((option) => {
					const ncmshValue = String(ncmsh || '').trim();
					const origemValue = String(origem || '').trim();
					const optionNcm = String(option.campo1 || '').trim();
					const optionOrigem = String(option.campo2 || '').trim();

					// Filtro baseado no NCM e origem
					const ncmMatches = optionNcm === ncmshValue;
					const origemMatches = optionOrigem === origemValue;

					return ncmMatches && origemMatches;
				});
			}

			// Formatar os produtos do Protheus para o Combobox
			produtosOptions = produtosProtheus.map((option) => ({
				label: `${option.label}`,
				value: option.value,
				campo1: option.campo1, // NCM
				campo2: option.campo2 // Origem
			}));
		}
	}

	// Função para lidar com a seleção de produto
	function handleProductSelect(event: CustomEvent) {
		const produtoSelecionado = event.detail.selected;
		console.log('Produto selecionado:', produtoSelecionado);
		// Lógica para salvar o produto selecionado ou realizar alguma ação
	}
</script>

<!-- Renderização do Combobox -->
<Combobox
	placeholder="Selecione um produto"
	options={produtosOptions}
	on:select={handleProductSelect}
	tit1="Saldo:"
	tit2="Valor total:"
	buttonClass="bg-base-300 w-full"
/>
