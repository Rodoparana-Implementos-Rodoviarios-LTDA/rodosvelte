<script lang="ts">
    import { onMount } from 'svelte';
    import { columns } from './columns'; // Importa as colunas da Portaria
    import { fetchDados } from '$lib/utils/fetchDados'; // Função para buscar dados
    import type { Portaria } from '$lib/types/tableTypes'; // O tipo para os dados da Portaria
    import Table from '$lib/components/ui/table/Table.svelte'; // Componente de tabela
  
    // Dados da tabela (inicialmente vazios)
    let data: Portaria[] = [];
    const token = sessionStorage.getItem('token') ?? undefined; // Se null, transforma em undefined
  
    // Função para carregar dados do endpoint
    async function carregarDados() {
      try {
        // Faz a requisição para o endpoint '/api/portaria'
        const result = await fetchDados<Portaria>('/api/portaria', token);
        data = result.data; // Atualiza os dados da tabela
      } catch (error) {
        console.error('Erro ao carregar dados:', error);
      }
    }
  
    // Carrega os dados assim que a página é montada
    onMount(() => {
      carregarDados();
    });
  </script>
  
  <div class="container mx-auto">
    <h1 class="text-2xl font-bold mb-4">Tabela de Portaria</h1>
  
    <!-- Renderizando a Tabela -->
    <Table {columns} {data} />
  </div>
  