<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/prenota/tabela/columns'; // Importa as colunas
	import type { PreNota } from '$lib/types/tableTypes';
	import Filtrar from '$lib/components/prenota/tabela/Filtrar.svelte'; // Importa o componente Filtrar

	// Estados de carregamento e paginação
	let isLoading = true;
	let currentPage = 1; // Controla a página atual
	let pageSize = 15; // Tamanho da página (pode ser dinâmico)
	let preNotas: PreNota[] = [];
	let hasMore = true;

	// Estado para definir o endpoint, ordenação e filtros
	export let endpoint: string = 'api/prenotas';
	export let sortBy: string = 'Inclusao'; // Campo para ordenar (dinâmico)
	export let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem de classificação (dinâmico)
	let filters: Record<string, string> = {}; // Filtros aplicados

	// Função para carregar os dados da página atual
	async function loadPage(page: number = 1) {
		isLoading = true;
		try {
			const result = await dataFetching(endpoint, sortBy, sortOrder, page, pageSize, filters);
			preNotas = result.data;
			hasMore = result.hasMore;
		} catch (error) {
			console.error(`Erro ao carregar dados de ${endpoint}:`, error);
		} finally {
			isLoading = false;
		}
	}

	// Função que escuta o evento de filtros aplicados e carrega a nova página com os filtros
	function handleFiltersApplied(event: CustomEvent<Record<string, string>>) {
		filters = event.detail; // Atualiza os filtros recebidos do Filtrar.svelte
		currentPage = 1; // Reseta para a primeira página quando os filtros são aplicados
		loadPage(currentPage); // Recarrega os dados com os filtros aplicados
	}

	// Função para resetar os filtros
	function handleFiltersReset(event: Event) {
		filters = {}; // Reseta os filtros no componente pai
		currentPage = 1; // Reseta para a primeira página
		loadPage(currentPage); // Recarrega os dados sem filtros
	}

	// Função para mudança de ordenação
	function handleSortChange(event: CustomEvent<{ sortBy: string; sortOrder: 'asc' | 'desc' }>) {
		sortBy = event.detail.sortBy;
		sortOrder = event.detail.sortOrder;
		loadPage(); // Recarrega a página com a nova ordenação
	}

	// Função para mudança de página
	function handlePageChange(event: CustomEvent<number>) {
		currentPage = event.detail; // Atualiza a página atual
		loadPage(currentPage); // Carrega os dados da nova página
	}

	// Carrega a primeira página quando o componente é montado
	onMount(() => {
		loadPage(currentPage);
	});
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título -->
	<div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
		<h1>Listagem de Pré Documentos de Entrada</h1>

		<!-- Componente de filtros -->
		<Filtrar
			{columns}
			on:applyFilters={handleFiltersApplied}
			on:resetFilters={handleFiltersReset}
		/>
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
				{sortBy}
				{sortOrder}
				on:sortChange={handleSortChange}
				on:pageChange={handlePageChange}
			/>
		</div>
	{/if}
</div>
