<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/table/Table.svelte'; // Componente da Tabela
	import { fetchDados } from '$lib/utils/fetchDados'; // Função de fetch para os dados
	import { columns } from '$lib/components/prenota/tabela/columns'; // Importa as colunas

	import type { PreNota } from '$lib/types/tableTypes';
	import {IconFilter} from "@tabler/icons-svelte"
	let isLoading = true;
	let currentPage = 1;
	let preNotas: PreNota[] = [];
	let hasMore = true;
	let sortBy: string | undefined = undefined;
	let sortOrder: 'asc' | 'desc' | undefined = undefined;

	// Estado para definir o endpoint dinamicamente
	export let endpoint: string = "prenotas"; // Valor default (pode ser passado de fora)
  
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
		// Passa o endpoint dinâmico para a função de fetch
		const result = await fetchDados(endpoint, page, 15, sortBy, sortOrder, filters);
		preNotas = result.data;
		hasMore = result.hasMore;
	  } catch (error) {
		console.error(`Erro ao carregar dados de ${endpoint}:`, error);
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
	  <label for="filters-drawer" class="btn btn-outline z-00"><IconFilter/></label>
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
	  <div class="drawer-side z-50">
		<label for="filters-drawer" class="drawer-overlay"></label>

	  </div>
	</div>
</div>