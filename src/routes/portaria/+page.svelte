<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import { columns } from '$lib/components/portaria/borracharia/columns'; // Importa as colunas
	import Filters from '$lib/components/prenota/tabela/Filters.svelte'; // Componente de Filtros
	import { goto } from '$app/navigation';
	import { IconFilter, IconHistory } from '@tabler/icons-svelte';

	let isLoading = true;
	let currentPage = 1; // Controla a página atual
	let preNotas = [];
	let hasMore = true; // Verifica se há mais páginas para navegar
	let sortBy: string | undefined = undefined;
	let sortOrder: 'asc' | 'desc' | undefined = undefined;

	// Estado dos filtros
	let selectedFilial = '';
	let selectedNF = '';
	let selectedCliente = '';
	let selectedVendedor = '';
	let selectedProduto = '';
	let selectedSaldo = '';
	let selectedEmissao = '';

	// Função que carrega a página com base no número da página
	async function loadPage(page = 1, filters = {}) {
		isLoading = true;
		try {
			const result = await dataFetching('borracharia', page, 10, sortBy, sortOrder, filters);
			preNotas = result.data;
			hasMore = result.hasMore; // Define se há mais páginas
			currentPage = page; // Atualiza a página atual
		} catch (error) {
			console.error('Erro ao carregar dados:', error);
		} finally {
			isLoading = false;
		}
	}

	// Função para capturar o evento de mudança de página e carregar a nova página
	function handlePageChange(event) {
		const nextPage = event.detail;
		loadPage(nextPage); // Carrega a próxima página usando o número passado no evento
	}

	// Função para aplicar os filtros
	function handleFilters(event) {
		selectedFilial = event.detail.selectedFilial;
		selectedNF = event.detail.selectedNF;
		selectedCliente = event.detail.selectedCliente;
		selectedVendedor = event.detail.selectedVendedor;
		selectedProduto = event.detail.selectedProduto;
		selectedSaldo = event.detail.selectedSaldo;
		selectedEmissao = event.detail.selectedEmissao;

		loadPage(1); // Recarrega a página com os novos filtros aplicados
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
		<h1>Listagem de Pré Documentos de Entrada</h1>
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
						pageData={preNotas}
						{currentPage}
						{hasMore}
						{sortBy}
						{sortOrder}
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
