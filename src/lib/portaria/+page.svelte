<script lang="ts">
	import { onMount } from 'svelte';
	import { columns } from './columns'; // Importa as colunas específicas de Borracharia
	import { fetchDados } from '$lib/utils/fetchDados'; // Função para buscar dados da API
	import type { Borracharia } from '$lib/types/tableTypes'; // Tipo específico para Borracharia
	import Table from '$lib/components/ui/table/Table.svelte'; // Componente de tabela reutilizável

	// Dados da tabela (inicialmente vazios)
	let data: Borracharia[] = [];

	// Função para buscar dados do endpoint borracharia
	async function carregarDados() {
		try {
			// Faz a requisição para o endpoint específico de borracharia
			const result = await fetchDados<Borracharia>('/api/borracharia');
			data = result.data; // Atualiza os dados da tabela
		} catch (error) {
			console.error('Erro ao carregar dados da borracharia:', error);
		}
	}

	// Carrega os dados assim que a página é montada
	onMount(() => {
		carregarDados();
	});
</script>

<div class="container mx-auto">
	<h1 class="text-2xl font-bold mb-4">Tabela de Borracharia</h1>

	<!-- Renderizando a Tabela -->
	<Table {columns} {data} />
</div>
