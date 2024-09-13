<script lang="ts">
	import { IconFilter } from '@tabler/icons-svelte';
	import { createEventDispatcher } from 'svelte';

	export let columns: { header: string; accessorKey: string; isFilterable: boolean }[] = [];

	const dispatch = createEventDispatcher();
	let filters: Record<string, string> = {};

	// Função chamada quando o valor de um input de filtro muda
	function handleInputChange(columnKey: string, event: Event) {
		const target = event.target as HTMLInputElement;
		filters[columnKey] = target.value;
	}

	// Função que aplica os filtros e envia os dados para o componente pai
	function applyFilters() {
		console.log('Filtros aplicados:', filters); // Verifica quais filtros foram aplicados
		dispatch('applyFilters', filters); // Emite os filtros para o componente pai
		closeDrawer();
	}

	// Função para resetar os filtros e enviar o evento para o componente pai
	function resetFilters() {
		filters = {}; // Limpa os filtros
		dispatch('resetFilters'); // Emite o evento para o componente pai
		closeDrawer();
	}

	// Função para fechar o drawer
	function closeDrawer() {
		const drawerInput = document.getElementById('filter-drawer') as HTMLInputElement;
		if (drawerInput) {
			drawerInput.checked = false;
		}
	}
</script>

<div class="flex justify-center items-center">
	<!-- Botão que abre o drawer -->
	<label for="filter-drawer" class="cursor-pointer flex items-center btn btn-outline w-full">
		<IconFilter class="h-6 w-6" />
	</label>

	<!-- Drawer da DaisyUI -->
	<div class="drawer drawer-end">
		<input id="filter-drawer" type="checkbox" class="drawer-toggle" />
		<div class="drawer-side">
			<label for="filter-drawer" class="drawer-overlay"></label>
			<ul class="menu bg-base-200 text-base-content min-h-full w-96 p-4">
				<!-- Exibe apenas as colunas que são filtráveis -->
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
					<button class="btn btn-outline w-fit m-5" on:click={resetFilters}>
						Reiniciar Filtros
					</button>
					<button class="btn btn-primary w-fit m-5" on:click={applyFilters}>
						Aplicar Filtros
					</button>
				</div>
			</ul>
		</div>
	</div>
</div>

<style>
	.drawer-side {
		z-index: 9999;
	}
</style>
