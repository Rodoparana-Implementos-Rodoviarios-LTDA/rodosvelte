<script lang="ts">
	import { columns } from '$lib/components/portaria/conferencia/columns'; // Importa as colunas
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import Filtrar from '$lib/components/ui/tabela/Filtrar.svelte';
	import { IconChevronDown } from '@tabler/icons-svelte';

	// Defina o endpoint no nível da página
	export let endpoint: string = 'api/portaria';
	let filters = {}; // Inicialmente, nenhum filtro aplicado
	let sortBy = 'DataHora'; // Coluna padrão para ordenação
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem padrão

	// Função para aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
		filters = event.detail; // Captura os filtros do evento
		// Dispara um evento para atualizar a tabela
		const applyEvent = new CustomEvent('applyFilters', { detail: filters });
		dispatchEvent(applyEvent);
	}

	// Função para resetar filtros e recarregar a tabela
	function resetFilters() {
		filters = {}; // Reseta os filtros
		// Dispara um evento personalizado para resetar a tabela
		const resetEvent = new CustomEvent('resetFilters');
		dispatchEvent(resetEvent);
	}
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
		<h1 class="text-3xl font-bold">Conferência de Saída de Itens</h1>
	</div>

	<!-- Área dos Filtros e Exportação -->
	<div class="flex flex-col items-end">
		<div class="flex justify-end space-x-4 pb-5">
			<!-- Dropdown de Navegação -->
			<div class="dropdown dropdown-end">
				<div tabindex="0" role="button" class="btn btn-outline m-1"><IconChevronDown /></div>
				<ul
					tabindex="0"
					class="dropdown-content menu bg-base-100 rounded-box z-[1] border border-secondary w-52 p-2 shadow"
				>
					<li><a class="text-lg" href="/controle-itens">Borracharia</a></li>
					<li>
						<a class="text-lg" href="/controle-itens/conferencia">Itens Conferidos</a>
					</li>
				</ul>
			</div>

			<!-- Componente de Filtros -->
			<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />
		</div>

		<!-- Componente da Tabela -->
		<Table {columns} {endpoint} {filters} {sortBy} {sortOrder} />
	</div>
</div>
