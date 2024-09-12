<script lang="ts">
    export let column;   // A coluna que está sendo renderizada
    export let rowData;   // Os dados da linha correspondente
  
    // Verifica se a célula usa uma função de renderização personalizada
    let cellContent;
    if (typeof column.cell === 'function') {
      cellContent = column.cell(rowData);   // Renderiza o conteúdo personalizado
    } else {
      cellContent = rowData[column.key];    // Renderiza o valor correspondente à chave
    }
  </script>
  
  <td class={column.class ?? ''}>
    {#if typeof cellContent === 'string'}
      {cellContent}  <!-- Renderiza texto -->
    {:else}
      {@html cellContent}  <!-- Renderiza HTML dinâmico se for HTML -->
    {/if}
  </td>
  