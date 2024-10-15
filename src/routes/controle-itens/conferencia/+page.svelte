<script lang="ts">
	import { onMount } from 'svelte';
	import { columns } from '$lib/components/portaria/conferencia/columns'; // Importa as colunas
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import Filtrar from '$lib/components/ui/tabela/Filtrar.svelte';

	// Defina o endpoint no nível da página
	export let endpoint: string = 'api/portaria';
	let filters: Record<string, string> = {}; // Inicialmente, nenhum filtro aplicado
	let sortBy = 'DataHora'; // Coluna padrão para ordenação
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem padrão
	let isTableReady = false; // Estado para garantir que a tabela seja recarregada corretamente

	// Função para aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
		filters = event.detail;
		console.log('Filtros aplicados:', filters);
		isTableReady = false; // Sinaliza que a tabela deve ser recarregada
	}

	// Função para resetar filtros
	function resetFilters() {
		filters = {};
		console.log('Filtros resetados');
		isTableReady = false; // Sinaliza que a tabela deve ser recarregada
	}

	// Garante que os filtros são aplicados após a página carregar
	onMount(() => {
		try {
			console.log('Página montada, aplicando filtros:', filters);

			// Se existir filtros salvos, aplica-os na inicialização
			if (Object.keys(filters).length > 0) {
				applyFilters(new CustomEvent('applyFilters', { detail: filters }));
			}
			isTableReady = true; // Permite que a tabela seja renderizada
		} catch (error) {
			console.error('Erro ao carregar a página ou aplicar filtros:', error);
		}
	});
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
		<h1 class="text-3xl font-bold">Conferência de Saída de Itens</h1>
	</div>

	<!-- Área dos Filtros e Exportação -->
	<div class="flex flex-col items-end">
		<div class="w-[95dvw] h-[70dvh] 2xl:h-[80dvh]">
			<div class="w-full flex justify-end pb-5 pr-1">
				<div class="join border border-primary">
					<!-- Componente de Filtros -->
					<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />
				</div>
			</div>

			<!-- Componente da Tabela -->
			{#if isTableReady}
				<Table {columns} {endpoint} {sortBy} {sortOrder} />
			{:else}
				<p>Carregando tabela...</p>
			{/if}
		</div>
	</div>
</div>
