<script lang="ts">
  import { onMount } from 'svelte';
  import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
  import { columns } from '$lib/components/portaria/historico/columns'; // Colunas da tabela
  import { fetchHistorico } from '$lib/services/histFetching'; // Função de fetching
  import type { HistoricoData } from '$lib/types/tableTypes';

  let isLoading = true;
  let paginaAtual = 1;      // Página inicial
  let tamanhoPagina = 10;   // Tamanho da página
  let historicoDados: HistoricoData[] = [];
  let possuiMais = true;
  let ordenarPor = 'DataHora';
  let ordemOrdenacao: 'asc' | 'desc' = 'desc';

  // Variáveis para armazenar os dados da linha selecionada
  let selectedRow: HistoricoData | null = null; // Objeto que guarda os dados da linha selecionada

  let isDrawerOpen = false; // Estado para controlar a visibilidade do drawer

  // Função para carregar os dados da API
  async function carregarPagina(pagina: number = 1) {
    isLoading = true;
    try {
      const result = await fetchHistorico(pagina, tamanhoPagina, ordenarPor, ordemOrdenacao);
      historicoDados = result.data;
      possuiMais = result.hasMore;
      paginaAtual = pagina;
    } catch (error) {
      console.error('Erro ao carregar dados:', error);
    } finally {
      isLoading = false;
    }
  }

  // Função para abrir o drawer com os dados da linha selecionada a partir do evento
  function abrirDrawer(event: CustomEvent) {
    const row = event.detail;  // Captura o objeto da linha a partir de event.detail
    selectedRow = row;         // Armazena os dados da linha clicada
    isDrawerOpen = true;       // Abre o drawer
  }

  // Função para capturar mudanças de ordenação
  function handleSortChange(event) {
    ordenarPor = event.detail.sortBy;
    ordemOrdenacao = event.detail.sortOrder;
    carregarPagina(paginaAtual);  // Recarregar com a nova ordenação
  }

  // Função para capturar mudança de página
  function handlePageChange(event) {
    const novaPagina = event.detail;
    if (possuiMais || novaPagina < paginaAtual) {
      carregarPagina(novaPagina);
    } else {
      console.log('Não há mais páginas para carregar.');
    }
  }

  // Carrega a primeira página ao montar o componente
  onMount(() => {
    carregarPagina(paginaAtual); // Carrega a primeira página ao montar
  });
</script>

<!-- Interface da Tabela -->
<div class="h-full">
  <!-- Cabeçalho com título -->
  <div class="flex justify-between items-center p-2 font-bold text-xl text-accent">
    <h1>Histórico de Movimentos</h1>
  </div>

  <!-- Exibir a Tabela de Dados -->
  {#if isLoading}
    <p>Carregando dados...</p>
  {:else}
    <div class="flex-1">
      <Table
        {columns} 
        pageData={historicoDados}
        currentPage={paginaAtual}
        hasMore={possuiMais}
        sortBy={ordenarPor} 
        sortOrder={ordemOrdenacao}
        on:sortChange={handleSortChange}
        on:pageChange={handlePageChange}
        on:rowClick={abrirDrawer}
      />
    </div>
  {/if}

  <!-- Drawer de Conferência -->
  {#if isDrawerOpen && selectedRow}
    <div 
      class="drawer fixed inset-0 bg-black bg-opacity-50 z-40 flex items-center justify-center"
      role="dialog" 
      aria-modal="true" 
      aria-labelledby="drawer-title"
    >
      <div class="drawer-side bg-white p-6 w-[400px] rounded-lg shadow-lg">
        <!-- Conteúdo do Drawer -->
        <h2 id="drawer-title" class="text-xl font-bold mb-6">Conferência de Produto</h2>
        
        <!-- Informações do Documento -->
        <div class="mb-4">
          <p><strong>Documento:</strong> {selectedRow.NF}</p>
          <p><strong>Produto:</strong> {selectedRow.Produto}</p>
          <p><strong>Responsável:</strong> {selectedRow.Responsavel}</p>
          <p><strong>Saldo Máximo:</strong> {selectedRow.Saldo.toFixed(2)}</p>
        </div>

        <!-- Botões de Ação -->
        <div class="flex justify-between">
          <button class="btn btn-outline" on:click={() => isDrawerOpen = false}>Cancelar</button>
        </div>
      </div>
    </div>
  {/if}
</div>

<style>
  .drawer-content {
    padding: 2rem;
  }
</style>
