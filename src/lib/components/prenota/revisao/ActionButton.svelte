<script lang="ts">
  import { fly } from "svelte/transition";
  import { circOut } from "svelte/easing";
  import { IconCopyPlus } from '@tabler/icons-svelte';
  import Timeline from './Timeline.svelte'; // Importa o componente Timeline
  import Produtos from './Produtos.svelte';
  import SolicitaRevisao from './SolicitaRevisao.svelte';

  export let status: string;
  export let rec: number;

  let tabs = ['Historico', 'Solicitar Revisão', 'Produtos'];
  let activeTab = 0;
  let drawerVisible = false; // Controla a visibilidade do drawer

  // Função que altera a aba ativa
  function setActiveTab(index: number) {
    activeTab = index;
  }

  // Função para fechar o Drawer ao clicar fora
  function closeDrawer(event: Event) {
    const target = event.target as HTMLElement;
    if (!target.closest('.drawer-content')) {
      drawerVisible = false;
    }
  }
</script>

<!-- Botão que abre o Drawer -->
<button class="btn btn-primary" on:click={() => drawerVisible = true}>
  <IconCopyPlus />
  <span class="sr-only">Abrir Histórico</span>
</button>

<!-- Drawer -->
{#if drawerVisible}
  <!-- Overlay que fecha o Drawer ao clicar fora -->
  <div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>

  <!-- Conteúdo do Drawer -->
  <div class="drawer drawer-content" transition:fly={{ x: 400, easing: circOut }}>
    <div class="w-full h-screen bg-base-100 text-info p-8">
      <div class="tabs tabs-boxed">
        {#each tabs as tab, index}
          <button
            class="tab tab-lifted"
            class:tab-active={activeTab == index}
            on:click={() => setActiveTab(index)}>{tab}</button
          >
        {/each}
      </div>
      <div class="w-full h-full pt-5">
        <div class="w-full h-full" class:hidden={activeTab != 0}>
          <Timeline {rec} {status} />
        </div>
        <div class="w-full h-full" class:hidden={activeTab != 1}>
          <SolicitaRevisao {rec}/>
        </div>
        <div class="w-full h-full" class:hidden={activeTab != 2}>
          <Produtos />
        </div>
      </div>
    </div>
  </div>
{/if}

<style>
  .drawer {
    background: #89dfff;
    position: fixed;
    top: 0;
    right: 0;
    height: 100vh;
    width: 400px; /* Largura do drawer */
    display: flex;
    justify-content: space-between;
    z-index: 50;
  }
</style>
