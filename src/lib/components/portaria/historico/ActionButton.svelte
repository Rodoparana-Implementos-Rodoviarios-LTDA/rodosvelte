<script lang="ts">
	import { IconCopyPlus } from "@tabler/icons-svelte";
  import ContentHist from "./ContentDrawer.svelte"; // Importa o novo componente

  export let saldoMaximo: number; // O saldo que será passado
  let isDrawerOpen = false; // Controla a abertura do Drawer

  // Função para abrir o Drawer
  function handleClick() {
    isDrawerOpen = true;
  }

  // Função para fechar o Drawer
  function closeDrawer() {
    isDrawerOpen = false;
  }

  // Função para detectar cliques fora do drawer
  function handleOutsideClick(event) {
    // Verifica se o clique é fora do conteúdo do drawer
    if (event.target.closest('.drawer-content') === null) {
      closeDrawer();
    }
  }
</script>

<!-- Botão que abre o Drawer -->
<button class="btn btn-primary" on:click={handleClick}>
  <IconCopyPlus/>
</button>

<!-- Drawer -->
{#if isDrawerOpen}
  <!-- Overlay para fechar o drawer ao clicar fora -->
  <div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={handleOutsideClick}></div>

  <!-- Conteúdo do Drawer -->
  <div class="fixed right-0 top-0 h-full w-80 bg-base-200 z-50 p-4 shadow-lg drawer-content">


    <!-- Insere o conteúdo do histórico com saldo máximo -->
    <ContentHist {saldoMaximo} on:closeDrawer={closeDrawer} />
  </div>
{/if}
