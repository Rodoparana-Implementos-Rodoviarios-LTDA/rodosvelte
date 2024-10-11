<script lang="ts">
	import { produtosStore, selectedTipoNF, produtosDoPedidoStore } from '$lib/stores/xmlStore';
	import { get } from 'svelte/store';
	import Combobox from '$lib/components/ui/Combobox/Combobox.svelte';

	// Props recebidas da coluna
	export let index: number;
	console.log('index:', index); // Verificar se o index está sendo passado corretamente
	export let origem: string;
	export let ncmsh: string;

	// Verificar se o produto está no pedido
	const Tipo = get(selectedTipoNF);
	const Optiones = get(produtosStore);
	let filtered = Optiones;

	// Função para formatar o índice numérico no formato '0001', '0002', etc.
	function formatIndex(index: number) {
		return String(index).padStart(4, '0'); // Converte o index para o formato '0001'
	}

	// Índice formatado
	const indexFormatado = formatIndex(index);

	$: produtosPedido = $produtosDoPedidoStore;

	$: produtoPedido = produtosPedido.find((produto) => produto.index === indexFormatado);
	$: {

		console.log('Produtos do Pedido disponíveis:', produtosPedido);

	}

	// Filtrar os produtos do Protheus se necessário
	if (Tipo === 'Matéria Prima' || Tipo === 'Revenda') {
		filtered = filtered.filter((option) => {
			const ncmshValue = String(ncmsh || '').trim();
			const origemValue = String(origem || '').trim();
			const optionNcm = String(option.campo1 || '').trim();
			const optionOrigem = String(option.campo2 || '').trim();
			const origemMatches = optionOrigem === origemValue;
			const ncmMatches = optionNcm === ncmshValue;

			return ncmMatches && origemMatches;
		});
	}

	// Função de handle para seleção de produto no Combobox
	async function handleProduct(event: CustomEvent) {
		// Lógica para lidar com a seleção do produto
	}
</script>

<!-- Renderização condicional: mostrar produto do pedido ou Combobox do Protheus -->
{#if produtoPedido}
	<div class="text-primary text0s">
		<p>{produtoPedido.codigo} - {produtoPedido.produto}</p>
	</div>
{:else}
	<!-- Mostra o Combobox de produtos do Protheus -->
	<Combobox
		placeholder="Produto do Protheus"
		on:select={handleProduct}
		options={filtered}
		buttonClass="bg-base-300"
		tit1="NCM:"
		tit2="Origem:"
	/>
{/if}
