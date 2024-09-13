<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import SortableHeader from './SortableHeader.svelte';
	import Filtrar from './Filtrar.svelte'; // Importa o componente de filtros
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetch para os dados
	import type { Column } from '$lib/types/tableTypes';

	// Props recebidas
	export let columns: Column<any>[] = [];
	export let endpoint: string;
	export let pageSize: number = 15; // Tamanho da página padrão

	let pageData: any[] = [];
	let currentPage: number = 1;
	let hasMore: boolean = true;
	let sortBy: string | undefined = 'Inclusao';
	let sortOrder: 'asc' | 'desc' = 'desc';
	let filters: Record<string, string> = {};
	let isLoading = true;

	const dispatch = createEventDispatcher();

	// Função para carregar os dados da página
	async function loadPage(page: number = 1) {
		isLoading = true;
		try {
			const result = await dataFetching(endpoint, sortBy!, sortOrder!, page, pageSize, filters);
			pageData = result.data;
			hasMore = result.hasMore;
		} catch (error) {
			console.error(`Erro ao carregar dados de ${endpoint}:`, error);
		} finally {
			isLoading = false;
		}
	}

	// Alternar a ordenação das colunas
	function toggleSort(columnKey: keyof any) {
		if (sortBy === columnKey) {
			sortOrder = sortOrder === 'asc' ? 'desc' : 'asc';
		} else {
			sortBy = String(columnKey);
			sortOrder = 'asc';
		}
		loadPage(); // Recarrega a tabela com a nova ordenação
	}

	// Navegar para a página anterior
	function prevPage() {
		if (currentPage > 1) {
			currentPage -= 1;
			loadPage(currentPage);
		}
	}

	// Navegar para a próxima página
	function nextPage() {
		if (hasMore) {
			currentPage += 1;
			loadPage(currentPage);
		}
	}

	// Alterar o número de itens por página
	function handlePageSizeChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		pageSize = parseInt(target.value);
		currentPage = 1; // Voltar para a primeira página ao mudar o tamanho da página
		loadPage(currentPage);
	}

	// Aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
		filters = event.detail;
		currentPage = 1;
		loadPage(currentPage);
	}

	// Resetar filtros
	function resetFilters() {
		filters = {};
		currentPage = 1;
		loadPage(currentPage);
	}

	// Carrega a primeira página ao montar o componente
	onMount(() => {
		loadPage();
	});
</script>

<!-- Tabela de dados -->
<div class="overflow-hidden z-1">
	{#if isLoading}
		<p>Carregando dados...</p>
	{:else}
		<!-- Adiciona o componente de filtro -->
		<div class="flex justify-end pb-5">
			<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />
		</div>
		<div
			class="overflow-auto scroll-insvisible w-[95dvw] h-[70dvh] 2xl:h-[75dvh] border border-primary rounded-btn shadow-md"
		>
			<table class="table table-pin-rows table-pin-cols z-0">
				<thead class="h-16">
					<tr>
						{#each columns as column}
							<th class="">
								<!-- Se o header for um componente, renderiza ele, caso contrário, renderiza o texto -->
								{#if typeof column.header === 'function'}
									<svelte:component this={column.header} {columns} />
								{:else}
									<SortableHeader
										title={column.header}
										isSorted={sortBy === String(column.accessorKey) ? sortOrder : false}
										onClick={() => toggleSort(column.accessorKey)}
									/>
								{/if}
							</th>
						{/each}
					</tr>
				</thead>

				<tbody>
					{#if pageData.length === 0}
						<tr>
							<td colspan={columns.length}>Nenhum dado encontrado</td>
						</tr>
					{:else}
						{#each pageData as row}
							<tr>
								{#each columns as column}
									<td class="relative max-w-56 text-start font-medium break-words">
										{#if column.component}
											<svelte:component
												this={column.component}
												{...column.props ? column.props(row) : {}}
											/>
										{:else}
											{row[String(column.accessorKey)]}
										{/if}
									</td>
								{/each}
							</tr>
						{/each}
					{/if}
				</tbody>
			</table>
		</div>

		<!-- Controles de navegação -->
		<div class="flex justify-between pt-5">
			<!-- Controle para selecionar o número de itens por página -->
			<div>
				<select class="select select-ghost w-full max-w-xs" on:change={handlePageSizeChange}>
					<option value="5">5</option>
					<option value="10">10</option>
					<option value="15" selected>15</option>
					<option value="30">30</option>
					<option value="50">50</option>
					<option value="100">100</option>
				</select>
			</div>

			<!-- Botões de navegação -->
			<div class="join">
				<button class="join-item btn" on:click={prevPage} disabled={currentPage === 1}>«</button>
				<button class="join-item btn">{currentPage}</button>
				<button class="join-item btn" on:click={nextPage} disabled={!hasMore}>»</button>
			</div>
		</div>
	{/if}
</div>
