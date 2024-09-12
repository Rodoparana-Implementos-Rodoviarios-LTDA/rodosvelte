<script lang="ts">
  export let table: any; // Recebe a instância da tabela via prop
</script>

<div class="join flex justify-center items-center py-4">
  <!-- Botão para ir à página anterior -->
  <button
    class="join-item btn"
    on:click={() => table.previousPage()}
    disabled={!table.getCanPreviousPage()}
  >
    «
  </button>

  <!-- Mostra a página atual e o total de páginas -->
  <button class="join-item btn">
    Página {table.getState().pagination.pageIndex + 1} de {table.getPageCount()}
  </button>

  <!-- Botão para ir à próxima página -->
  <button
    class="join-item btn"
    on:click={() => table.nextPage()}
    disabled={!table.getCanNextPage()}
  >
    »
  </button>
</div>
