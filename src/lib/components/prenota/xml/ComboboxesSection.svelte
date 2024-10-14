<script lang="ts">
	// Imports de componentes
	import Combobox from '$lib/components/ui/Combobox/Combobox.svelte';
	// Importando os tipos necessários
	import type { ComboboxOption } from '$lib/types/Produtos';
	import type { CondicaoOption } from '$lib/types/CargaInicio';
	import type { Pedido } from '$lib/types/Pedidos';

	// Imports de stores
	import {
		tiposNFOptions,
		prioridadeOptions,
		condicoesStore,
		selectedTipoNF,
		selectedPriority,
		selectedCondicao,
		pedidosDetalhadosStore,
		produtosDoPedidoStore
	} from '$lib/stores/xmlStore';

	// Imports de funções
	import { fetchProdutos } from '$lib/services/fetchProdutos';
	import { get } from 'svelte/store';

	/********************* Combobox de Pedidos *********************/

	let pedidosOptions: ComboboxOption[] = []; // Opções para o Combobox de pedidos
	let pedidosSelecionados: Pedido[] = [];

	// Reatividade para atualizar pedidosOptions sempre que a store for atualizada

	$: pedidosOptions = FormataPedido($pedidosDetalhadosStore);

	// Função que formata os pedidos (mantendo como objeto)
	function FormataPedido(pedidos: Record<string, Pedido>): ComboboxOption[] {
		// Verifica se é um objeto válido
		if (typeof pedidos !== 'object' || pedidos === null) {
			return [];
		}

		// Mapeia os valores do objeto de pedidos (sem transformar em array explicitamente)
		return Object.keys(pedidos).map((key) => {
			const pedido = pedidos[key];
			return {
				label: pedido.pedido, // Nome do pedido
				value: pedido.pedido, // Número do pedido
				campo1: pedido.produtos[0].pedidoProduto,
				campo2: pedido.status
			};
		});
	}

	// Função de manipulação da seleção de pedidos
	function handlePedidoSelect(event: CustomEvent) {
		const selectedValues = event.detail.selected.map((item) => item.value); // Obtém os valores dos pedidos selecionados

		// Filtrar os produtos dos pedidos selecionados
		pedidosSelecionados = selectedValues.map((value) => $pedidosDetalhadosStore[value]);

		// Extrair todos os produtos dos pedidos selecionados
		const produtosSelecionados = pedidosSelecionados.flatMap((pedido) => pedido.produtos);

		// Atualiza a store com os produtos dos pedidos selecionados (para mostrar no Combobox)
		produtosDoPedidoStore.set(produtosSelecionados);

		if (pedidosSelecionados.length > 0) {
			const primeiroPedido = pedidosSelecionados[0]; // Pega o primeiro pedido selecionado
			const condicaoPagamentoDoPedido = primeiroPedido.pagamento; // Condição de pagamento do pedido

			// Atualiza o valor de condPedSel com a condição de pagamento
			condPedSel = condicaoPagamentoDoPedido;
		}
	}
	/********************* Combobox de Tipo de NF *********************/

	// Função de manipulação da seleção de tipo de NF
	async function handleTipoDeNFSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedTipoNF.set(selectedValue.label);
		await fetchProdutos();
	}

	/********************* Combobox de Prioridade *********************/

	// Função de manipulação da seleção de prioridade
	function handlePrioridadeSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedPriority.set(selectedValue.label);
	}

	/********************* Combobox de Condição de Pagamento ***********/

	let condicaoOptions: CondicaoOption[] = []; // Opções para o Combobox de condição de pagamento
	let condPedSel = ''; // Condição de pagamento selecionada

	// Formatação das opções de condição de pagamento
	function formatCondicaoOptions(condicoes) {
		return condicoes.map((condicao) => ({
			label: condicao.E4_DESCRI.trim(),
			value: condicao.E4_CODIGO.trim()
		}));
	}
	$: condicaoOptions = formatCondicaoOptions($condicoesStore);

	// Função de manipulação da seleção de condição de pagamento
	function handleCondicaoSelect(event: CustomEvent) {
		const selectedValue = event.detail.selected;
		selectedCondicao.set(selectedValue.label);
	}
</script>

<!-- Estrutura do componente com os Comboboxes -->
<div id="Inputs" class="card w-1/2 p-5 gap-5 bg-base-100 justify-evenly">
	<div class="flex justify-center">
		<span class="font-bold text-lg">Adicionar Informações sobre a NF</span>
	</div>

	<div class="flex w-full justify-evenly gap-5">
		<div class="w-full">
			<Combobox
				options={pedidosOptions}
				placeholder="Selecione um ou mais Pedidos"
				on:select={handlePedidoSelect}
				buttonClass="bg-base-300 w-full"
				multiple
				tit2="Status: "
			/>
		</div>

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
		<div class="w-full">
			<Combobox
				options={prioridadeOptions}
				placeholder="Prioridade"
				on:select={handlePrioridadeSelect}
				buttonClass="bg-base-300 w-full"
			/>
		</div>

		<div class="w-full">
			<Combobox
				options={condicaoOptions}
				placeholder="Forma de Pagamento"
				value={condPedSel}
				on:select={handleCondicaoSelect}
				buttonClass="bg-base-300 w-full"
			/>
		</div>
	</div>
</div>
