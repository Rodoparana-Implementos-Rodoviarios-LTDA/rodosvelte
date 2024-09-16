<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/portaria/borracharia/columns'; // Importa as colunas
	import type { borracharia } from '$lib/types/tableTypes';

	// Estados de carregamento e paginação
	let isLoading = true;
	let currentPage = 1; // Controla a página atual
	let pageSize = 10; // Tamanho da página (pode ser dinâmico)
	let borracharia: borracharia[] = [];
	let filteredBorracharia: borracharia[] = []; // Dados filtrados
	let hasMore = true;

	// Estado para definir o endpoint e ordenação
	export let endpoint: string = 'api/borracharia';
	export let sortBy: string = 'Inclusao'; // Campo para ordenar (dinâmico)
	export let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem de classificação (dinâmico)

	// Função para carregar os dados da página atual
	async function loadPage(page: number = 1) {
		isLoading = true;
		try {
			const result = await dataFetching(endpoint, sortBy, sortOrder, page, pageSize);
			borracharia = result.data;
			applyFilter(); // Aplica o filtro para remover linhas com "REC"
			hasMore = result.hasMore;
		} catch (error) {
			console.error(`Erro ao carregar dados de ${endpoint}:`, error);
		} finally {
			isLoading = false;
		}
	}

	// Função para aplicar o filtro nas Notas Fiscais
	function applyFilter() {
		// Remove linhas onde a Nota Fiscal contém "REC"
		filteredBorracharia = borracharia.filter(item => !item.NF.includes("REC"));
	}

	// Escuta quando o componente Table emite o evento de ordenação
	function handleSortChange(event) {
		sortBy = event.detail.sortBy;
		sortOrder = event.detail.sortOrder;
		loadPage(); // Recarrega a página com a nova ordenação
	}

	// Escuta quando o componente Table emite o evento de mudança de página
	function handlePageChange(event) {
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
		<h1>Lista de Notas Borracharia</h1>
	</div>

	<!-- Conteúdo da Tabela -->
	{#if isLoading}
		<p>Carregando dados...</p>
	{:else}
		<div class="flex-1">
			<Table 
				{columns} 
				pageData={filteredBorracharia} 
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
