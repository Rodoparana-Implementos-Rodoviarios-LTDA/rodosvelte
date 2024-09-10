<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/prenota/tabela/columns'; // Importa as colunas
	import Filters from '$lib/components/prenota/tabela/Filters.svelte'; // Componente de Filtros
	import type { PreNota } from '$lib/types/tableTypes';
  
	let isLoading = true;
	let currentPage = 1;
	let preNotas: PreNota[] = [];
	let hasMore = true;
	let sortBy: string | undefined = undefined;
	let sortOrder: 'asc' | 'desc' | undefined = undefined;
  
	// Estados para filtros
	let selectedStatus = '';
	let selectedFilial = '';
	let selectedFornecedor = '';
	let selectedTipo = '';
	let selectedPrioridade = '';
  
	// Função para carregar os dados da página atual
	async function loadPage(page = 1, filters = {}) {
	  isLoading = true;
	  try {
		const result = await dataFetching(page, 10, sortBy, sortOrder, filters);
		preNotas = result.data;
		hasMore = result.hasMore;
	  } catch (error) {
		console.error('Erro ao carregar pre_notas:', error);
	  } finally {
		isLoading = false;
	  }
	}
  
	// Função para mudar a página
	function changePage(newPage: number) {
	  currentPage = newPage;
	  loadPage(currentPage);
	}
  
	// Função para mudar a ordenação
	function changeSort({
	  sortBy: newSortBy,
	  sortOrder: newSortOrder
	}: {
	  sortBy: string;
	  sortOrder: 'asc' | 'desc';
	}) {
	  sortBy = newSortBy;
	  sortOrder = newSortOrder;
	  loadPage(currentPage);
	}
  
	// Função para aplicar os filtros
	function handleFilters(event) {
	  selectedStatus = event.detail.selectedStatus;
	  selectedFilial = event.detail.selectedFilial;
	  selectedFornecedor = event.detail.selectedFornecedor;
	  selectedTipo = event.detail.selectedTipo;
	  selectedPrioridade = event.detail.selectedPrioridade;
  
	  loadPage(1); // Recarrega a primeira página com os filtros
	}
  
	// Carrega a primeira página quando o componente é montado
	onMount(() => {
	  loadPage(currentPage);
	});
  </script>
  
  <div class="h-full">
	<!-- Cabeçalho com o título e botão de filtros -->
	<div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
	  <h1>Listagem de Pré Documentos de Entrada</h1>
	  <label for="filters-drawer" class="btn btn-outline btn-secondary">Filtros</label>
	</div>
  
	<!-- Drawer de Filtros -->
	<div class="drawer drawer-end">
	  <input id="filters-drawer" type="checkbox" class="drawer-toggle" />
	  <div class="drawer-content">
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
			  on:pageChange={(e) => changePage(e.detail)}
			  on:sortChange={(e) => changeSort(e.detail)}
			/>
		  </div>
		{/if}
	  </div>
	
	  <!-- Drawer Lateral de Filtros -->
	  <div class="drawer-side">
		<label for="filters-drawer" class="drawer-overlay"></label>
		<Filters
		  bind:selectedStatus
		  bind:selectedFilial
		  bind:selectedFornecedor
		  bind:selectedTipo
		  bind:selectedPrioridade
		  on:applyFilters={handleFilters}
		/>
	  </div>
	</div>
  </div>
  