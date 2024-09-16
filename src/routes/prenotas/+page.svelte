<script lang="ts">
	import { columns } from '../../lib/components/prenota/tabela/columns'; // Importa as colunas
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { fetchAndSaveFiliais } from '$lib/services/filiaisFetch';
	import { onMount } from 'svelte';

	// Defina o endpoint no nível da página
	export let endpoint: string = 'api/prenotas';

	// Carrega as filiais no IndexedDB ao carregar a página
	onMount(async () => {
		await fetchAndSaveFiliais(); // Chama a função para buscar e salvar as filiais
	});
</script>

<!-- Interface da Tabela -->
<div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
		<h1 class="text-3xl font-bold">Listagem de Pré Documentos de Entrada</h1>
	</div>

	<!-- Componente da Tabela -->
	<Table {columns} {endpoint} />
</div>
