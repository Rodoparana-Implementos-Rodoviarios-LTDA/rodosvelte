<script lang="ts">
    import { IconCheck, IconDots } from '@tabler/icons-svelte';
    import type { Revisao } from '$lib/types/revisao';
  
    // Recebe os eventos e o saldo como prop
    export let eventos: Revisao[] = [];
    export let saldo: number = 0; // Valor do saldo, que deve ser passado como prop
  
    let quantidadePneus = 1; // Número de pneus selecionados
    let responsavel = ''; // Nome do responsável
    let placa = ''; // Placa do carro
  
    // Função para enviar os dados (pode ser ajustada conforme sua lógica de backend)
    function enviarSolicitacao() {
      console.log('Quantidade de pneus:', quantidadePneus);
      console.log('Nome do responsável:', responsavel);
      console.log('Placa do carro:', placa);
    }
  </script>
  
  <div class="h-full flex flex-col justify-between items-center">
    <div class="w-full flex justify-between">
      <h2 class="text-3xl text-center font-bold pb-10">Solicitar Revisão e Adição de Pneus</h2>
      <button class="btn btn-outline">Anexos</button>
    </div>
  
    <!-- Formulário para adicionar pneus e solicitar revisão -->
    <div class="flex flex-col w-full flex-1 gap-3">
      
      <!-- Slider para selecionar a quantidade de pneus -->
      <div class="mb-4">
        <label for="quantidadePneus" class="block text-sm font-medium">Quantidade de pneus: {quantidadePneus}</label>
        <input
          type="range"
          id="quantidadePneus"
          min={1}
          max={saldo}
          bind:value={quantidadePneus}
          class="range range-primary"
        />
        <p class="text-sm">Saldo disponível: {saldo}</p>
      </div>
  
      <!-- Nome do responsável e placa do carro -->
      <input
        type="text"
        placeholder="Nome do responsável"
        bind:value={responsavel}
        class="input input-bordered w-full mb-4"
      />
  
      <input
        type="text"
        placeholder="Placa do carro"
        bind:value={placa}
        class="input input-bordered w-full mb-4"
      />
  
      <!-- Botão de Solicitar revisão -->
      <button class="btn btn-primary w-fit" on:click={enviarSolicitacao}>
        Solicitar revisão
      </button>
    </div>
  
    <div class="divider text-xl">Histórico de Eventos</div>
  
    <!-- Exibição do histórico de eventos -->
    <ul class="timeline timeline-snap-icon max-md:timeline-compact timeline-vertical max-h-[80vh] overflow-y-auto">
      {#each eventos as event, index}
        <li>
          {#if index > 0}
            <hr class="bg-primary" />
          {/if}
          <div class="timeline-start text-end">
            <div class="text-xl capitalize">{event.status}</div>
            <time class="text-xs font-mono">{event.data} - {event.hora}</time>
          </div>
  
          <div class="timeline-middle">
            {#if event.status === 'classificar'}
              <div class="rounded-full p-1 m-0.5 bg-primary/20">
                <IconDots stroke={3} class="h-3 w-3 text-black" />
              </div>
            {:else}
              <div class="rounded-full p-1 m-0.5 bg-primary">
                <IconCheck stroke={3} class="h-3 w-3 text-black" />
              </div>
            {/if}
          </div>
  
          <div class="timeline-end timeline-box w-full">
            <div class="flex justify-between">
              <div class="text-start">
                <div class="text-sm font-black">Usuário</div>
                <div class="text-sm font-light">{event.usuario}</div>
              </div>
              <div class="text-end">
                <div class="text-sm font-black">Campo</div>
                <div class="text-sm font-light">{event.campo}</div>
              </div>
            </div>
  
            {#if event.anterior}
              <div class="flex justify-between">
                <div class="text-start">
                  <div class="text-sm font-black">Anterior</div>
                  <div class="text-sm font-light">{event.anterior}</div>
                </div>
                <div class="text-end">
                  <div class="text-sm font-black">Valor Atual</div>
                  <div class="text-sm font-light max-w-36">{event.atual}</div>
                </div>
              </div>
            {/if}
          </div>
  
          <hr class="bg-primary" />
        </li>
      {/each}
    </ul>
  </div>
  