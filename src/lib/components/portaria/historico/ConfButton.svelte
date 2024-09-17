<script lang="ts">
  import { saldoStore } from '$lib/components/portaria/historico/saldoStore'; // Importa a store de saldo
  import { fly } from 'svelte/transition'; // Transição fly
  import { circOut } from 'svelte/easing'; // Easing para a transição
  import { createEventDispatcher } from 'svelte'; // Dispatcher de eventos
  import { onMount } from 'svelte';
	import { IconCopyPlus } from '@tabler/icons-svelte';

  export let documento: string;
  export let produto: string;
  let saldoMaximo = 0;

  let quantity = 0; // Quantidade inicial para o slider
  let isChecked = false; // Estado do checkbox para confirmar a conferência
  let isLoading = false;
  let errorMessage = "";
  let drawerVisible = false;

  const dispatch = createEventDispatcher();

  // Inscreve-se na store para receber o saldo da borracharia
  onMount(() => {
    saldoStore.subscribe(value => {
      saldoMaximo = value;
    });
  });

  // Função para lidar com a ação de confirmar e enviar os dados
  async function handleConfirm() {
    if (!isChecked || quantity <= 0) {
      errorMessage = "Por favor, confirme a conferência e insira uma quantidade válida.";
      return;
    }

    const body = {
      documento,
      produto,
      quantidade: quantity,
      usuario: "usuarioAtual",
    };

    try {
      isLoading = true;
      const response = await fetch('http://172.16.99.174:8400/rest/MovPortaria/IncluirItem', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
      });

      if (!response.ok) {
        throw new Error('Erro ao enviar a solicitação.');
      }

      console.log('Solicitação enviada com sucesso');
      errorMessage = ''; // Limpa a mensagem de erro
      drawerVisible = false; // Fecha o drawer após a confirmação
    } catch (error) {
      console.error(error);
      errorMessage = 'Erro ao enviar a solicitação.';
    } finally {
      isLoading = false;
    }
  }

  // Função para fechar o drawer ao clicar fora
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
</button>

<!-- Drawer -->
{#if drawerVisible}
  <!-- Overlay que fecha o Drawer ao clicar fora -->
  <div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>

  <!-- Conteúdo do Drawer com animação fly -->
  <div class="drawer drawer-content" transition:fly={{ x: 400, easing: circOut }}>
    <div class="w-full h-screen bg-base-100 text-info p-8">
      <h2 class="text-lg font-bold mb-4">Conferência de Produto</h2>

      <!-- Informações do Documento -->
      <p><strong>Documento:</strong> {documento}</p>
      <p><strong>Produto:</strong> {produto}</p>
      <p><strong>Saldo disponível:</strong> {saldoMaximo}</p>

      <!-- Range Slider para quantidade -->
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

      <!-- Checkbox para confirmar a conferência -->
      <div class="form-control mb-6">
        <label class="cursor-pointer label">
          <span class="label-text">Confirmar conferência</span>
          <input type="checkbox" class="checkbox checkbox-success" bind:checked={isChecked} />
        </label>
      </div>

      <!-- Botões de Ação -->
      <div class="flex justify-between mt-6">
        <button class="btn btn-primary" on:click={handleConfirm} disabled={isLoading}>
          {#if isLoading}
            Enviando...
          {:else}
            Confirmar
          {/if}
        </button>
        <button class="btn btn-outline" on:click={() => drawerVisible = false}>
          Cancelar
        </button>
      </div>

      <!-- Mensagem de erro -->
      {#if errorMessage}
        <p class="text-error mt-4">{errorMessage}</p>
      {/if}
    </div>
  </div>
{/if}

<style>
  .drawer {
    position: fixed;
    top: 0;
    right: 0;
    height: 100vh;
    width: 400px;
    display: flex;
    justify-content: space-between;
    z-index: 50;
  }
</style>
