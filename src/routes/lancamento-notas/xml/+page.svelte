<script lang="ts">
	import Cabecalho from '$lib/components/prenota/xml/Cabecalho.svelte';
	import ProdutosXMLTable from '$lib/components/prenota/xml/ProdutosXMLTable.svelte';
	import { onMount } from 'svelte';

	let xmlKey = ''; // Armazena a chave XML recebida do cabeçalho

	// Função para capturar o evento emitido pelo Cabecalho
	function handleChaveXML(event) {
		xmlKey = event.detail.xmlKey; // Atualiza a chave XML com o valor do evento
		localStorage.setItem('xmlKey', xmlKey); // Salva a chave no LocalStorage
	}

	// Restaurar a chave XML do LocalStorage ao carregar a página
	onMount(() => {
		const storedKey = localStorage.getItem('xmlKey');
		if (storedKey) {
			xmlKey = storedKey;
		}
	});
</script>

<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
	<h1 class="text-3xl font-bold">Importação de XML</h1>
</div>

<main class="h-[85vh] w-[98vw] flex flex-col justify-end items-baseline gap-10">
	<header class="h-fit w-full">
		<!-- Ouvindo o evento chaveXML e passando a função de captura -->
		<Cabecalho on:chaveXML={handleChaveXML} />
	</header>
	<section class="w-full h-full">
		<!-- Renderiza a tabela de produtos apenas se o xmlKey estiver disponível -->
		{#if xmlKey}
			<ProdutosXMLTable {xmlKey} />
		{:else}
			<p class="text-center text-lg">Por favor, insira uma chave XML para exibir os detalhes.</p>
		{/if}
	</section>
</main>
