<script lang="ts">
	import { onMount } from 'svelte';
	import Table from '$lib/components/ui/tabela/Table.svelte';
	import { dataFetching } from '$lib/services/dataFetching.svelte';
	import { columns } from '$lib/components/portaria/borracharia/columns';
	import { goto } from '$app/navigation'; // Para redirecionamento
	import type { BorrachariaData } from '$lib/types/tableTypes';

	let isLoading = true;
	let currentPage = 1;
	let borrachariaData: BorrachariaData[] = [];
	let hasMore = true;
	let sortBy: string | undefined = undefined;
	let sortOrder: 'asc' | 'desc' = 'asc';

	// Função para carregar os dados da página atual
	async function loadPage(page = 1) {
		isLoading = true;
		try {
			const result = await dataFetching('portaria', page, 10, sortBy, sortOrder);
			console.log('API Response:', result);
			borrachariaData = result.data;
			hasMore = result.hasMore;
		} catch (error) {
			console.error('Erro ao carregar os dados:', error);
		} finally {
			isLoading = false;
		}
	}

	function changePage(newPage: number) {
		currentPage = newPage;
		loadPage(currentPage);
	}

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

	// Função para voltar para a tabela principal de Borracharia
	function goToBorracharia() {
		goto('/portaria'); // Alterar o caminho para a tabela de Borracharia
	}

	onMount(() => {
		loadPage(currentPage);
	});
</script>

<div class="h-full">
	<div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
		<h1>Histórico Borracharia</h1>
		<!-- Botão para voltar para a página principal da Borracharia -->
		<button class="btn btn-outline btn-primary" on:click={goToBorracharia}>
			Voltar para Borracharia
		</button>
	</div>
	<div class="flex-1">
		{#if isLoading}
			<p>Carregando dados...</p>
		{:else}
			<Table
				{columns}
				pageData={borrachariaData}
				{currentPage}
				{hasMore}
				{sortBy}
				{sortOrder}
				on:pageChange={(e) => changePage(e.detail)}
				on:sortChange={(e) => changeSort(e.detail)}
			/>
		{/if}
	</div>
</div>
