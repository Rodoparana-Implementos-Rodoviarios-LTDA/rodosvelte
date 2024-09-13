<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/portaria/borracharia/columns'; // Importa as colunas
	import Filters from '$lib/components/prenota/tabela/Filters.svelte'; // Componente de Filtros
	import { goto } from '$app/navigation';
	import { IconFilter, IconHistory } from '@tabler/icons-svelte'; // Ícones de filtro e histórico

	// Estados de carregamento e paginação
	let isLoading = true;
	let currentPage = 1; // Controla a página atual
	let pageSize = 10; // Tamanho da página (pode ser dinâmico)
	let borracharia = [];
	let hasMore = true; // Verifica se há mais páginas para navegar

	// Estado da ordenação
	let sortBy: string = 'Inclusao'; // Campo para ordenar (dinâmico)
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem de classificação (dinâmico)

	// Estado dos filtros
	let selectedFilial = '';
	let selectedNF = '';
	let selectedCliente = '';
	let selectedVendedor = '';
	let selectedProduto = '';
	let selectedSaldo = '';
	let selectedEmissao = '';

	// Função que carrega a página com base no número da página e filtros
	async function loadPage(page: number = 1, filters = {}) {
		isLoading = true;
		try {
			const result = await dataFetching('borracharia', page, pageSize, sortBy, sortOrder, filters);
			borracharia = result.data;
			hasMore = result.hasMore; // Define se há mais páginas
		} catch (error) {
			console.error('Erro ao carregar dados:', error);
		} finally {
			isLoading = false;
		}
	}

	// Função para capturar o evento de mudança de página e carregar a nova página
	function handlePageChange(event) {
		currentPage = event.detail; // Atualiza a página atual
		loadPage(currentPage); // Carrega a próxima página usando o número passado no evento
	}

	// Função para aplicar os filtros
	function handleFilters(event) {
		// Captura os valores dos filtros
		selectedFilial = event.detail.selectedFilial;
		selectedNF = event.detail.selectedNF;
		selectedCliente = event.detail.selectedCliente;
		selectedVendedor = event.detail.selectedVendedor;
		selectedProduto = event.detail.selectedProduto;
		selectedSaldo = event.detail.selectedSaldo;
		selectedEmissao = event.detail.selectedEmissao;

		// Recarrega a primeira página com os filtros aplicados
		loadPage(1, {
			selectedFilial,
			selectedNF,
			selectedCliente,
			selectedVendedor,
			selectedProduto,
			selectedSaldo,
			selectedEmissao
		});
	}

	// Função para capturar o evento de ordenação
	function handleSortChange(event) {
		sortBy = event.detail.sortBy;
		sortOrder = event.detail.sortOrder;
		loadPage(currentPage); // Recarrega a página com a nova ordenação
	}

	// Carrega a primeira página assim que o componente é montado
	onMount(() => {
		loadPage(currentPage);
	});
</script>

<!-- Main Content -->
<div class="h-full">
	<!-- Header com Filtros e Histórico -->
	<div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
		<h1>Lista de pneus borracharia</h1>
		<div>
			<!-- Botão para abrir o filtro -->
			<label for="filters-drawer" class="btn btn-outline btn-secondary">
				<IconFilter />
			</label>

			<!-- Botão para abrir o histórico -->
			<button class="btn btn-outline btn-secondary" on:click={() => goto('/portaria/historico')}>
				<IconHistory />
			</button>
		</div>
	</div>

	<!-- Drawer for Filters -->
	<div class="drawer drawer-end">
		<input id="filters-drawer" type="checkbox" class="drawer-toggle" />
		<div class="drawer-content">
			<!-- Tabela de dados -->
			{#if isLoading}
				<p>Carregando dados...</p>
			{:else}
				<div class="flex-1">
					<Table
						{columns}
						pageData={borracharia}
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

		<!-- Drawer Sidebar for Filters -->
		<div class="drawer-side">
			<label for="filters-drawer" class="drawer-overlay"></label>
			<Filters
				bind:selectedFilial
				bind:selectedNF
				bind:selectedCliente
				bind:selectedVendedor
				bind:selectedProduto
				bind:selectedSaldo
				bind:selectedEmissao
				on:applyFilters={handleFilters}
			/>
		</div>
	</div>
</div>
