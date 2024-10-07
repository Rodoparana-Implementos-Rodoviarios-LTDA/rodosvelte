<script lang="ts">
	import { xmlItemsStore } from '$lib/stores/xmlStore'; // Importa a store dos itens da XML
	import { fetchState } from '$lib/services/conexaoNFE'; // Importa a store de estado da requisição
	import { columns } from './columns'; // Importa as colunas definidas
	import type { ConexaoNFE } from '$lib/types/ConexaoNFE'; // Tipo da XML

	let items: ConexaoNFE['itens'] | null = null; // Armazena os itens da XML

	// Usa o estado da store fetchState
	$: estadoAtual = $fetchState;

	// Usa o sinal $ para obter a reatividade automática da store xmlItemsStore
	$: items = $xmlItemsStore;

	// Função para parar a propagação do evento dentro dos componentes da célula
	function handleClick(event: Event) {
		event.stopPropagation(); // Para a propagação manualmente
	}
</script>

{#if estadoAtual === 'inicial'}
	<div class="h-full flex items-center justify-center bg-base-300">
		<p class="text-center text-info text-lg">
			Por favor, insira uma chave XML para buscar os produtos.
		</p>
	</div>
{:else if estadoAtual === 'carregando'}
	<div class="flex w-full flex-col gap-4">
		<div class="flex items-center gap-4">
			<div class="skeleton h-16 w-16 shrink-0 rounded-full"></div>
			<div class="flex flex-col gap-4 w-full">
				<div class="skeleton h-4 w-full"></div>
				<div class="skeleton h-4 w-full"></div>
			</div>
		</div>
		<div class="skeleton h-32 w-full"></div>
	</div>
{:else if estadoAtual === 'erro'}
	<p class="text-red-500">Erro ao buscar os produtos da XML. Tente novamente.</p>
{:else if estadoAtual === 'sucesso' && items && items.length > 0}
	<div class="bg-base-300 h-full rounded-md shadow-lg w-full">
		<div class="overflow-auto h-full">
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
		</div>
	</div>
{:else}
	<div class="h-full flex items-center justify-center">
		<p class="text-center text-lg">Nenhum produto encontrado para esta XML.</p>
	</div>
{/if}
