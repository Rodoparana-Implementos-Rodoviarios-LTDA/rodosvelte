<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/prenota/tabela/columns'; // Importa as colunas
	import type { PreNota } from '$lib/types/tableTypes';
  
	// Estados de carregamento e paginação
	let isLoading = true;
	let currentPage = 1;
	let preNotas: PreNota[] = [];
	let hasMore = true;
  
	// Estado para definir o endpoint dinamicamente
	export let endpoint: string = "api/prenotas"; // Valor default (pode ser passado de fora)
	
	// Função para carregar os dados da página atual
	async function loadPage() {
	  isLoading = true;
	  try {
		// Chama a função dataFetching com os headers travados
		const result = await dataFetching(endpoint);
		preNotas = result.data;
		hasMore = result.hasMore;
	  } catch (error) {
		console.error(`Erro ao carregar dados de ${endpoint}:`, error);
	  } finally {
		isLoading = false;
	  }
	}
	
	// Carrega a primeira página quando o componente é montado
	onMount(() => {
	  loadPage();
	});
  </script>
  
  <!-- Interface da Tabela -->
  <div class="h-full">
	<!-- Cabeçalho com o título -->
	<div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
	  <h1>Listagem de Pré Documentos de Entrada</h1>
	</div>
  
	<!-- Conteúdo da Tabela -->
	{#if isLoading}
	  <p>Carregando dados...</p>
	{:else}
	  <div class="flex-1">
		<Table
		  {columns}
		  pageData={preNotas}
		  {currentPage}
		  {hasMore}
		/>
	  </div>
	{/if}
  </div>
  