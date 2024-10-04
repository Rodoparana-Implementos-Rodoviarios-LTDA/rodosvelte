<script lang="ts">
	import { onMount } from 'svelte';
	import { getOptionsData } from '$lib/services/idb'; // Função para buscar dados do IndexedDB
	import { columns } from './columns'; // Importa as colunas definidas
	import type { XML } from '$lib/types';

	export let xmlKey: string; // Chave recebida via prop
	let items: XML[] = []; // Armazena os itens da XML
	let loading: boolean = false;
	let error: string | null = null;
	let pageSize: number = 10; // Número de itens por página
	let currentPage: number = 1; // Página atual
	let paginatedItems: XML[] = []; // Itens paginados

	// Função para buscar os itens do IndexedDB com base na chave xmlKey
	async function fetchItensFromIndexedDB() {
		try {
			loading = true;
			error = null;

			// Busca os dados no IndexedDB
			const result = await getOptionsData('xml', xmlKey);
			if (result && result.data && result.data.itens) {
				items = result.data.itens; // Armazena os itens
				updatePaginatedItems(); // Atualiza a paginação
				console.log('Itens carregados do IndexedDB:', items);
			} else {
				throw new Error('Nenhum item encontrado para a chave fornecida.');
			}
		} catch (err) {
			error = (err as Error).message;
			console.error('Erro ao buscar itens no IndexedDB:', error);
		} finally {
			loading = false;
		}
	}

	// Função para atualizar os itens paginados
	function updatePaginatedItems() {
		const start = (currentPage - 1) * pageSize;
		const end = start + pageSize;
		paginatedItems = items.slice(start, end);
	}

	// Função para ir para a próxima página
	function nextPage() {
		if (currentPage < Math.ceil(items.length / pageSize)) {
			currentPage += 1;
			updatePaginatedItems();
		}
	}

	// Função para ir para a página anterior
	function prevPage() {
		if (currentPage > 1) {
			currentPage -= 1;
			updatePaginatedItems();
		}
	}

	// Monta o componente e busca os itens
	onMount(() => {
		fetchItensFromIndexedDB();
	});
</script>

<!-- Renderiza a tabela -->
<div class="overflow-auto border border-primary rounded-md shadow-md w-full">
	{#if loading}
		<p>Carregando itens...</p>
	{:else if error}
		<p class="text-red-500">{error}</p>
	{:else if paginatedItems.length === 0}
		<p>Nenhum item encontrado.</p>
	{:else}
		<table class="table table-zebra w-full">
			<thead>
				<tr>
					{#each columns as column}
						<th class={column.class}>{column.header}</th>
					{/each}
				</tr>
			</thead>
			<tbody>
				{#each paginatedItems as row}
					<tr>
						{#each columns as column}
							<td class={column.class}>
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
			</tbody>
		</table>

		<!-- Controles de paginação -->
		<div class="flex justify-between items-center pt-4">
			<button class="btn" on:click={prevPage} disabled={currentPage === 1}>Página anterior</button>
			<span>Página {currentPage} de {Math.ceil(items.length / pageSize)}</span>
			<button
				class="btn"
				on:click={nextPage}
				disabled={currentPage === Math.ceil(items.length / pageSize)}
			>
				Próxima página
			</button>
		</div>
	{/if}
</div>
