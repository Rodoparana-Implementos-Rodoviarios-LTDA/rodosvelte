<script lang="ts">
	import Cabecalho from '$lib/components/prenota/xml/Cabecalho.svelte';
	import ProdutosXMLTable from '$lib/components/prenota/xml/ProdutosXMLTable.svelte';
	import { onMount } from 'svelte';
	import { xmlItemsStore } from '$lib/stores/xmlStore'; // Importa a store de itens

	let loadingProdutos: boolean = false; // Estado de carregamento para os produtos
	let produtosError = null; // Variável para armazenar erro na tabela de produtos

	// Espera os dados da store serem atualizados
	$: {
		if ($xmlItemsStore.length > 0) {
			loadingProdutos = false;
			produtosError = null;
		} else {
			produtosError = 'Nenhum produto encontrado para esta XML.';
			loadingProdutos = false;
		}
	}

	// Iniciar o estado de carregamento ao montar
	onMount(() => {
		loadingProdutos = true;
	});
</script>

<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
	<h1 class="text-3xl font-bold">Importação de XML</h1>
</div>

<main class="h-[85vh] w-[98vw] flex flex-col gap-10">
	<header class="h-fit w-full">
		<Cabecalho />
	</header>
	<section class="w-full h-full">
		<ProdutosXMLTable />
	</section>
</main>
