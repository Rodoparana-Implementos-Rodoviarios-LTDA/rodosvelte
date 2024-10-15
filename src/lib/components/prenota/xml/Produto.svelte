<script lang="ts">
	import Combobox from '$lib/components/ui/Combobox/Combobox.svelte';
	import { produtosDoPedidoStore, produtosStore, selectedTipoNF } from '$lib/stores/xmlStore';
	import { formatarProdutosDoPedidoFiltrados } from './produtosDoPedidoStore'; // Função de filtro para pedidos
	import { obterProdutosProtheusFiltrados } from './produtosStore'; // Função de filtro para produtos do Protheus
	import type { ComboboxOption } from '$lib/types/Produtos';
  
	// Props recebidas da coluna (para NCM e origem)
	export let origem: string;
	export let ncmsh: string;
	export let quantidade: number;
	export let valorUnitario: number;
	export let valorTotal: number;
  
	// Variável para armazenar as opções de produtos
	let produtosOptions: ComboboxOption[] = [];
  
	// Obtenção reativa das stores
	$: tipoNF = $selectedTipoNF;
	$: produtosProtheus = $produtosStore;
	$: produtosDoPedido = $produtosDoPedidoStore;
  
	// Lógica para decidir os produtos a serem mostrados, reativamente
	$: {
	  if (produtosDoPedido && produtosDoPedido.length > 0) {
		// Caso tenha produtos do pedido, mostrar produtos do pedido
		if (tipoNF) {
		  // Se houver pedido e tipo de NF, aplicar os filtros
		  produtosOptions = formatarProdutosDoPedidoFiltrados(ncmsh, origem, quantidade, valorUnitario, valorTotal);
		} else {
		  // Caso tenha só produtos do pedido, sem tipo de NF, mostrar todos
		  produtosOptions = formatarProdutosDoPedidoFiltrados('', '', 0, 0, 0);
		}
	  } else if (!produtosDoPedido || produtosDoPedido.length === 0) {
		// Caso não tenha produtos do pedido, mas tenha tipo de NF, usar produtos do Protheus
		if (tipoNF) {
		  produtosOptions = obterProdutosProtheusFiltrados(ncmsh, origem);
		} else {
		  // Caso nenhum pedido e nenhum tipo de NF, sem filtro
		  produtosOptions = [];
		}
	  }
	}
  
	// Função para lidar com a seleção de produto no Combobox
	function handleProductSelect(event: CustomEvent) {
	  const produtoSelecionado = event.detail.selected;
	  console.log('Produto selecionado:', produtoSelecionado);
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
  