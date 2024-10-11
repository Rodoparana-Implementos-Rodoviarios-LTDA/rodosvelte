<script lang="ts">
	// Imports de componentes e stores
	import Combobox from '$lib/components/ui/Combobox/Combobox.svelte';
	import {
		tiposNFOptions,
		prioridadeOptions,
		condicoesStore,
		selectedTipoNF,
		selectedPriority,
		selectedCondicao,
		pedidoComboStore,
		pedidosDetalhadosStore,
		produtosDoPedidoStore
	} from '$lib/stores/xmlStore';
	import { fetchProdutos } from '$lib/services/fetchProdutos';
	import { get } from 'svelte/store';
	import type { Pedido } from '$lib/types/Pedidos';
	import type { CondicaoOption } from '$lib/types/CargaInicio';

	let selectedPedidos = []; // Para armazenar os pedidos selecionados
	let ProdPedSel = []; // Produtos dos pedidos selecionados
	let condicoesPagamentoDoPedido = []; // Condições de pagamento dos pedidos selecionados

	// Dados para Combobox de pedidos e condições de pagamento
	let pedidosOptions = [];
	let condicaoOptionsFiltradas = []; // Condições filtradas
	let condicaoOptions: CondicaoOption[] = [];
	$: pedidosOptions = formatPedidosOptions($pedidoComboStore);
	$: condicaoOptions = formatCondicaoOptions($condicoesStore);

	// Função para formatar os pedidos no Combobox
	function formatPedidosOptions(pedidos) {
		return pedidos.map((pedido) => ({
			label: `Pedido: ${pedido.label}`,
			value: pedido.value,
			campo1: pedido.produto,
			campo2: pedido.status
		}));
	}

	// Função para formatar as condições de pagamento
	function formatCondicaoOptions(condicoes) {
		return condicoes.map((condicao) => ({
			label: condicao.E4_DESCRI.trim(),
			value: condicao.E4_CODIGO.trim()
		}));
	}

	// Função para filtrar as condições de pagamento pelos métodos usados nos pedidos selecionados
	function filtrarCondicoesPorPedidos(pedidosSelecionados) {
		// Coleta todas as condições de pagamento dos produtos nos pedidos selecionados
		const condicoesPagamento = pedidosSelecionados.flatMap((pedido) =>
			pedido.produtos.map((produto) => produto.condicao_pagamento)
		);

		// Remove duplicatas
		condicoesPagamentoDoPedido = [...new Set(condicoesPagamento)];

		// Filtra as opções de condição de pagamento com base nos métodos usados nos pedidos
		condicaoOptionsFiltradas = condicaoOptions.filter((condicao) =>
			condicoesPagamentoDoPedido.includes(condicao.value)
		);
	}

	// Função para tratar a seleção de Pedidos (múltipla seleção)
	function handlePedidoSelect(event: CustomEvent) {
		const selectedPedidosNovos = event.detail.selected;
		console.log('Pedidos selecionados:', selectedPedidosNovos);

		// Obtenha todos os pedidos detalhados da store
		const pedidosDetalhados: Pedido[] = get(pedidosDetalhadosStore);

		// Filtra os pedidos selecionados
		selectedPedidos = pedidosDetalhados.filter((pedido) =>
			selectedPedidosNovos.find((selected) => selected.value === pedido.pedido)
		);

		// Filtrar as condições de pagamento com base nos pedidos selecionados
		filtrarCondicoesPorPedidos(selectedPedidos);

		// Coletar todos os produtos dos pedidos selecionados
		ProdPedSel = selectedPedidos.flatMap((pedido) => pedido.produtos);

		// Atualiza a store com os produtos selecionados
		produtosDoPedidoStore.set(ProdPedSel);
	}

	// Outras funções auxiliares para tratar Tipo de NF e Prioridade...
	async function handleTipoDeNFSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedTipoNF.set(selectedValue.label);
		await fetchProdutos();
		console.log('Tipo de NF selecionado:', selectedValue);
	}

	function handlePrioridadeSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedPriority.set(selectedValue.label);
		console.log('Prioridade selecionada:', selectedValue);
	}

	function handleCondicaoSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedCondicao.set(selectedValue.label);
		console.log('Condição de pagamento selecionada:', selectedValue);
	}

	$: {
		console.log('Condições de pagamento disponíveis:', condicaoOptionsFiltradas);
	}
</script>

<!-- Estrutura do componente -->
<div id="Inputs" class="card w-1/2 p-5 gap-5 bg-base-100 justify-evenly">
	<div class="flex justify-center">
		<span class="font-bold text-lg">Adicionar Informações sobre a NF</span>
	</div>

	<div class="flex w-full justify-evenly gap-5">
		<!-- Combobox para Pedidos -->
		<div class="w-full">
			<Combobox
				options={pedidosOptions}
				placeholder="Selecione um ou mais Pedidos"
				on:select={handlePedidoSelect}
				buttonClass="bg-base-300 w-full"
				multiple
				tit1="Primeiro Produto:"
			/>
		</div>
		<!-- Combobox para Tipo de NF -->
		<div class="w-full">
			<Combobox
				options={tiposNFOptions}
				placeholder="Tipo de NF"
				on:select={handleTipoDeNFSelect}
				buttonClass="bg-base-300 w-full"
			/>
		</div>
	</div>

	<div class="flex w-full justify-evenly gap-5">
		<!-- Combobox para Prioridade -->
		<div class="w-full">
			<Combobox
				options={prioridadeOptions}
				placeholder="Prioridade"
				on:select={handlePrioridadeSelect}
				buttonClass="bg-base-300 w-full"
			/>
		</div>

		<!-- Combobox para Condição de Pagamento -->
		<div class="w-full">
			<Combobox
				options={condicaoOptionsFiltradas} F
				placeholder="Forma de Pagamento"
				on:select={handleCondicaoSelect}
				buttonClass="bg-base-300 w-full"
			/>
		</div>
	</div>
</div>
