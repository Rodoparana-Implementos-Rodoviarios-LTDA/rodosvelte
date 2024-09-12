<script lang="ts">
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
  </script>
  
  <!-- Botão que abre o Drawer -->
  <button class="btn btn-primary" on:click={handleClick}>
    Abrir Drawer
  </button>
  
  <!-- Drawer -->
  {#if isDrawerOpen}
    <div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>
    <div class="fixed right-0 top-0 h-full w-80 bg-white z-50 p-4">
      <h2 class="text-lg font-bold">Seleção de Pneus</h2>
      <p>Saldo disponível: {saldoMaximo}</p>
  
      <!-- Conteúdo do Drawer -->
      <div>
        <input type="number" placeholder="Quantidade de pneus" min="1" class="input input-bordered mb-4" />
        <input type="text" placeholder="Nome do responsável" class="input input-bordered mb-4" />
        <input type="text" placeholder="Placa do carro" class="input input-bordered mb-4" />
      </div>
  
      <!-- Botões de Ação dentro do Drawer -->
      <div class="flex justify-between mt-4">
        <button class="btn btn-sm btn-primary" on:click={closeDrawer}>Confirmar</button>
        <button class="btn btn-sm btn-outline" on:click={closeDrawer}>Cancelar</button>
      </div>
    </div>
  {/if}
  