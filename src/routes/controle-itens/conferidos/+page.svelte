<script lang="ts">
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { columns } from '$lib/components/portaria/itensConferidos/columns'; // Colunas atualizadas
	import Filtrar from '$lib/components/ui/tabela/Filtrar.svelte';
	import { onMount } from 'svelte';
	import { dataFetching } from '$lib/services/dataFetching'; // Função de fetching de dados

	// Endpoint específico para conferidos
	const endpoint: string = 'api/pneus/historico'; // Endpoint correto

	let filters: { filial?: string } = {}; // Inicialmente, nenhum filtro aplicado
	let sortBy = 'Z08_DATA'; // Coluna padrão para ordenação
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem padrão
	let data: any[] = []; // Variável para armazenar os dados da API
	let isLoading = true; // Estado de carregamento
	let errorMessage = ''; // Mensagem de erro
	let hasMore = false; // Controle de paginação

	// Função para buscar os dados diretamente do endpoint
	async function fetchData() {
		try {
			isLoading = true;

			// Chama a função de dataFetching com os parâmetros apropriados
			const { data: apiData, hasMore: apiHasMore } = await dataFetching<any>(
				endpoint, sortBy, sortOrder, 1, 50, filters
			);

			data = apiData;
			hasMore = apiHasMore;

			console.log('Dados recebidos da API:', data);
		} catch (error: any) {
			console.error('Erro ao buscar dados da API:', error);
			errorMessage = `Erro ao carregar dados: ${error.message}`;
		} finally {
			isLoading = false;
		}
	}

	// Função para aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
		filters = event.detail; // Captura os filtros do evento
		fetchData(); // Recarrega os dados com os novos filtros
	}

	// Função para resetar os filtros
	function resetFilters() {
		filters = {}; // Reseta os filtros
		fetchData(); // Recarrega os dados sem filtros
	}

	// Carregar os dados quando o componente é montado
	onMount(() => {
		fetchData(); // Chama a função para obter os dados ao montar o componente
	});
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
		<h1 class="text-3xl font-bold">Itens Conferidos - Lista de Produtos</h1>
	</div>

	<!-- Área dos Filtros e Exportação -->
	<div class="flex flex-col items-end">
		<div class="flex justify-end space-x-4 pb-5">
			<!-- Componente de Filtros -->
			<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />
		</div>

		<!-- Mensagem de erro -->
		{#if errorMessage}
			<p class="text-error mt-4">{errorMessage}</p>
		{/if}

		<!-- Mensagem de carregamento -->
		{#if isLoading}
			<p>Carregando...</p>
		{:else}
			<!-- Componente da Tabela com dados carregados diretamente -->
			<Table {columns} {endpoint} />
		{/if}
	</div>
</div>
