<script lang="ts">
	import { columns } from '../../lib/components/prenota/tabela/columns'; // Importa as colunas
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import Filtrar from '$lib/components/ui/tabela/Filtrar.svelte';
	import { IconFileTypeXml, IconPencilUp } from '@tabler/icons-svelte';
	
	// Importando os dados retornados do load function
	export let unidadeMedida: any[];
	export let condicoes: any[];
	export let centoCusto: any[];

	// Defina o endpoint no nível da página
	export let endpoint: string = 'api/prenotas';
	let filters = {}; // Inicialmente, nenhum filtro aplicado
	let sortBy = 'Inclusao'; // Coluna padrão para ordenação
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem padrão

	// Função para aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
		filters = event.detail; // Captura os filtros do evento
		console.log('Filtros aplicados:', filters); // Verifica os filtros aplicados
	}

	// Função para resetar filtros e recarregar a tabela
	function resetFilters() {
		filters = {}; // Reseta os filtros
		console.log('Filtros resetados');
	}
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
		<h1 class="text-3xl font-bold">Lançamento de Notas</h1>
	</div>

	<!-- Área dos Filtros e Exportação -->
	<div class="flex flex-col items-end">
		<div class="w-[95dvw] h-[70dvh] 2xl:h-[80dvh]">
			<div class="w-full flex justify-end pb-5 pr-1">
				<div class="join border border-primary">
					<!-- Componente de Filtros -->
					<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />

					<!-- Botões de ação -->
					<button
						class="tooltip tooltip-info tooltip-bottom btn hover:text-primary"
						data-tip="Incluir Nota Manual"
					>
						<a href="/lancamento-notas/incluir"><IconPencilUp /></a>
					</button>

					<button
						class="tooltip tooltip-info tooltip-bottom btn hover:text-primary"
						data-tip="Importar via XML"
					>
						<a href="/lancamento-notas/xml"><IconFileTypeXml /></a>
					</button>
				</div>
			</div>

			<!-- Componente da Tabela -->
			<Table {columns} {endpoint} {filters} {sortBy} {sortOrder} />
		</div>
	</div>
</div>
