<script lang="ts">
  import { createEventDispatcher } from 'svelte'; // Importa o dispatcher

  export let saldoMaximo: number; // Recebe o saldo máximo
  let quantity = 0; // Quantidade inicial para o range
  let responsible = "";
  let plate = "";
  let selectedOption = "Cliente"; // Variável para controlar a seleção entre Cliente e Motorista

  const dispatch = createEventDispatcher(); // Cria o dispatcher

  // Função de exemplo para lidar com a ação de confirmar
  function handleConfirm() {
    console.log("Confirmação realizada");
    console.log("Quantidade:", quantity);
    console.log("Responsável:", responsible);
    console.log("Placa:", plate);
    console.log("Tipo de seleção:", selectedOption);
  }
</script>

<div class="h-full">
  <h2 class="text-lg font-bold mb-4">Seleção de Pneus</h2>
  <p class="mb-4">Saldo disponível: {saldoMaximo}</p>



  <!-- Range Slider de quantidade de pneus -->
  <div class="mb-4">
    <input
      type="range"
      min="0"
      max={saldoMaximo} 
      bind:value={quantity}
      class="range range-primary"
    />
    <p>Quantidade selecionada: {quantity}</p>
  </div>
  <!-- Radio Button para selecionar entre Cliente e Motorista -->
  <div class="mb-4 flex space-x-4">
    <label class="flex items-center">
      <input 
        type="radio" 
        name="radio-role" 
        class="radio radio-primary" 
        value="Cliente" 
        bind:group={selectedOption} 
        checked
      />
      <span class="ml-2">Cliente</span>
    </label>

    <label class="flex items-center">
      <input 
        type="radio" 
        name="radio-role" 
        class="radio radio-primary" 
        value="Motorista" 
        bind:group={selectedOption}
      />
      <span class="ml-2">Motorista</span>
    </label>
  </div>
  

  <!-- Inputs para o nome do responsável e a placa -->
  <div class="mb-4 space-y-2">
    <input
      type="text"
      placeholder="Nome do responsável"
      bind:value={responsible}
      class="input input-bordered w-full"
    />
    <input
      type="text"
      placeholder="Placa do carro"
      bind:value={plate}
      class="input input-bordered w-full"
    />
  </div>
  

  <!-- Botões de Ação -->
  <div class="flex justify-between mt-6">
    <button class="btn btn-primary" on:click={handleConfirm}>
      Confirmar
    </button>
    <button class="btn btn-outline" on:click={() => dispatch('closeDrawer')}>
      Cancelar
    </button>
  </div>
</div>
