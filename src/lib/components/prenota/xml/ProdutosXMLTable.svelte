<script lang="ts">
	import { xmlItemsStore } from '$lib/stores/xmlStore'; // Importa a store dos itens
	import { columns } from './columns'; // Importa as colunas definidas
	import { onMount } from 'svelte';
	import { writable } from 'svelte/store';

	let loading = false; // Controla o estado de carregamento
	let error: string | null = null; // Armazena possíveis erros
	let items = []; // Armazena os itens da XML

	// Monitora a store dos itens da XML
	$: items = $xmlItemsStore;

	// Função de inicialização, setando o estado de carregamento
	onMount(() => {
		loading = items.length === 0; // Se os itens estiverem vazios, exibe o estado de carregamento
		error = items.length === 0 ? 'Nenhum item encontrado' : null;
	});
</script>

<!-- Renderiza a tabela com scroll interno -->
<div class="overflow-auto border border-primary rounded-md shadow-md w-full h-[400px]">
	{#if loading}
		<p>Carregando itens...</p>
	{:else if error}
		<p class="text-red-500">{error}</p>
	{:else if items.length === 0}
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
				{#each items as row}
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
	{/if}
</div>
