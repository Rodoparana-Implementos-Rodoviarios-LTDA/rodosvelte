<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import SortableHeader from './SortableHeader.svelte';
	import type { Column } from '$lib/types/tableTypes';

	// Declarando o tipo genérico <T> para aceitar diferentes tipos de dados
	export let columns: Column<T>[] = [];
	export let pageData: T[] = [];
	export let currentPage: number = 1;
	export let hasMore: boolean = true;
	export let sortBy: string | undefined = undefined;
	export let sortOrder: 'asc' | 'desc' | undefined = undefined;
	export let pageSize: number = 15; // Número de itens por página (padrão)

	const dispatch = createEventDispatcher();

	// Função para alternar a ordenação das colunas
	function toggleSort(columnKey: keyof T) {
		if (sortBy === columnKey) {
			sortOrder = sortOrder === 'asc' ? 'desc' : 'asc';
		} else {
			sortBy = String(columnKey); // Garantir que columnKey seja tratado como string
			sortOrder = 'asc';
		}
		dispatch('sortChange', { sortBy, sortOrder });
	}

	// Função para navegar para a página anterior
	function prevPage() {
		if (currentPage > 1) {
			dispatch('pageChange', currentPage - 1);
		}
	}

	// Função para navegar para a próxima página
	function nextPage() {
		if (hasMore) {
			dispatch('pageChange', currentPage + 1);
		}
	}

	// Função para alterar o número de itens por página
	function handlePageSizeChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		const newPageSize = parseInt(target.value);
		dispatch('pageSizeChange', newPageSize); // Emite o novo tamanho de página para o componente pai
	}
</script>

<!-- Tabela de dados -->
<div class="overflow-hidden z-1">
	<div class="overflow-auto scroll-insvisible w-[95dvw] h-[70dvh] 2xl:h-[75dvh] border border-primary rounded-btn shadow-md">
		<table class="table table-pin-rows table-pin-cols z-0">
			<thead class="h-16">
				<tr>
					{#each columns as column}
						<th class="">
							<!-- Se o header for um componente, renderiza ele, caso contrário, renderiza o texto -->
							{#if typeof column.header === 'function'}
								<svelte:component this={column.header} {columns} />
								<!-- Passa as colunas para o componente Filtrar -->
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
	<div class="flex justify-between pt-2">
		<!-- Controle para selecionar o número de itens por página -->
		<div>
			 <select class="select select-ghost w-full max-w-xs" on:change={handlePageSizeChange}>
				<option disabled selected>Itens por Página</option>
				<option value="5">5</option>
				<option value="10">10</option>
				<option value="15">15</option>
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
</div>
