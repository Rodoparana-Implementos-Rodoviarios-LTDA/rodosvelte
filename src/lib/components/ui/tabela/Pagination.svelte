<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let currentPage: number = 1; // Página atual
	export let pageSize: number = 15; // Número de linhas por página
	export let hasMore: boolean = false; // Indica se há mais páginas

	const dispatch = createEventDispatcher(); // Para disparar eventos

	// Função chamada ao mudar o tamanho da página
	function handlePageSizeChange(event: Event) {
		const target = event.target as HTMLSelectElement;
		const newPageSize = parseInt(target.value);
		pageSize = newPageSize; // Atualiza o tamanho da página
		dispatch('pageSizeChange', pageSize); // Emite o evento de mudança de tamanho da página
	}

	// Função para ir à página anterior
	function goToPreviousPage() {
		if (currentPage > 1) {
			currentPage--;
			dispatch('pageChange', currentPage); // Emite o evento de mudança de página
		}
	}

	// Função para ir à próxima página
	function goToNextPage() {
		if (hasMore) {
			currentPage++;
			dispatch('pageChange', currentPage); // Emite o evento de mudança de página
		}
	}
</script>

<div class="w-full flex justify-between pt-5">
	<!-- Controle para selecionar o número de itens por página -->
	<div>
		<select class="select w-full max-w-xs" bind:value={pageSize} on:change={handlePageSizeChange}>
			<option disabled selected>{pageSize}</option>
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
		<button class="join-item btn" on:click={goToPreviousPage} disabled={currentPage === 1}>«</button
		>
		<button class="join-item btn">{currentPage}</button>
		<button class="join-item btn" on:click={goToNextPage} disabled={!hasMore}>»</button>
	</div>
</div>
