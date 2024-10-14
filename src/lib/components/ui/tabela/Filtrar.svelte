<script lang="ts">
	import { filterStore, setFilter, resetFilters } from '$lib/stores/filterStore';
	import { IconFilterFilled } from '@tabler/icons-svelte';
	export let columns: { header: string; accessorKey: string; isFilterable: boolean }[] = [];

	// Reatividade automática à store de filtros
	$: $filterStore;

	function handleInputChange(columnKey: string, event: Event) {
		const target = event.target as HTMLInputElement;
		setFilter(columnKey, target.value); // Atualiza diretamente a store de filtros
	}

	function applyFilters() {
		closeDrawer();
	}

	function resetAllFilters() {
		resetFilters(); // Reseta a store de filtros
		closeDrawer();
	}

	function closeDrawer() {
		const drawerInput = document.getElementById('filter-drawer') as HTMLInputElement;
		if (drawerInput) {
			drawerInput.checked = false;
		}
	}
</script>

<!-- Drawer do DaisyUI -->
<div class="drawer drawer-end ">
	<input id="filter-drawer" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content">
		<!-- Botão que abre o drawer -->
		<label
			for="filter-drawer"
			class="tooltip tooltip-info tooltip-bottom btn drawer-button hover:text-primary flex justify-center items-center"
			data-tip="Filtrar dados"
		>
			<IconFilterFilled />
		</label>
	</div>

	<div class="drawer-side z-[9999]">
		<label for="filter-drawer" class="drawer-overlay"></label>
		<ul class="menu bg-base-200 text-base-content min-h-full w-96 p-4">
			<!-- Inputs de Filtro para cada coluna filtrável -->
			{#each columns as column}
				{#if column.isFilterable}
					<li class="px-5">
						<input
							type="text"
							placeholder={`Filtrar ${column.header}`}
							class="input input-bordered w-full my-2"
							on:input={(e) => handleInputChange(column.accessorKey, e)}
						/>
					</li>
				{/if}
			{/each}

			<div class="flex gap-2">
				<button class="btn btn-outline w-fit m-5" on:click={resetFilters}>Reiniciar Filtros</button>
				<button class="btn btn-primary w-fit m-5" on:click={applyFilters}>Aplicar Filtros</button>
			</div>
		</ul>
	</div>
</div>
