<script lang="ts">
    import { onMount } from 'svelte';
    import type { ItemNF } from '$lib/types/tableTypes';
    import { IconTagsFilled } from '@tabler/icons-svelte';
  
    export let documentoCompleto: string;
    export let clienteCompleto: string;
    export let filial: string;
  
    export let itensSelecionados: ItemNF[] = [];
  
    let errorMessageFetch = '';
    let usuario: string = 'matheus'; // Pode ser dinâmico conforme necessário
  
    // Função para separar documento e série
    function separarDocumentoESerie(docCompleto: string): { documento: string; serie: string } {
      if (!docCompleto) return { documento: '', serie: '' };
      const partes = docCompleto.split(' - ').map((part) => part.trim());
      const documento = partes[0] || '';
      const serie = partes[1] || '';
      return { documento, serie };
    }
  
    // Função para separar cliente e loja
    function separarClienteLoja(clienteCompleto: string): { cliente: string; loja: string } {
      if (!clienteCompleto) return { cliente: '', loja: '' };
      const partes = clienteCompleto.split(' - ').map((part) => part.trim());
      const cliente = partes[0] || '';
      const loja = partes[1] || '';
      return { cliente, loja };
    }
  
    // Função para buscar itens da NF utilizando o endpoint listanfs
    async function fetchItensNF() {
      const { documento, serie } = separarDocumentoESerie(documentoCompleto);
      const { cliente, loja } = separarClienteLoja(clienteCompleto);
  
      const filialTrimmed = filial?.trim();
      if (!filialTrimmed || !/^\d{4}$/.test(filialTrimmed)) {
        errorMessageFetch = `Filial inválida: Deve conter exatamente 4 dígitos numéricos. Valor atual: '${filialTrimmed}'`;
        console.error(errorMessageFetch);
        return;
      }
  
      // Definir os headers conforme os filtros necessários
      const headers = {
        'Content-Type': 'application/json',
        'X-Filial': filialTrimmed,
        'X-Documento': documento.trim(),
        'X-Serie': serie.trim(),
        'X-Cliente': cliente.trim(),
        'X-Loja': loja.trim()
      };
  
      console.log('Headers da requisição:', headers);
  
      const url = `http://rodoapp:8080/api/pneus/borracharia/listanfs`;
  
      try {
        const response = await fetch(url, {
          method: 'GET',
          headers: headers
        });
  
        console.log('Status da resposta:', response.status);
  
        if (!response.ok) {
          const errorText = await response.text();
          console.error('Erro na resposta da API:', errorText);
          throw new Error(
            `Erro ao buscar itens: ${response.status} ${response.statusText} - ${errorText}`
          );
        }
  
        const data = await response.json();
        console.log('Dados recebidos da API:', data);
  
        if (!data || data.length === 0) {
          console.warn('Nenhum item foi retornado pela API.');
        }
  
        // Processar os dados recebidos e atualizar itensSelecionados
        itensSelecionados =
          data.map((item: any) => ({
            D2_ITEM: item.D2_ITEM?.trim(),
            D2_COD: item.D2_COD?.trim(),
            B1_DESC: item.B1_DESC?.trim(),
            SALDO: Number(item.SALDO), // Garantir que SALDO é um número
            quantity: 0 // Inicializa quantity com zero
          })) || [];
      } catch (error: any) {
        console.error('Erro ao buscar itens da NF:', error);
        errorMessageFetch = `Erro ao carregar itens da nota fiscal: ${error.message}`;
      }
    }
  
    // Carregar os itens ao montar o componente
    onMount(() => {
      fetchItensNF();
    });
  
    // Função para definir a quantidade diretamente
    function definirQuantidade(item: ItemNF, valor: string) {
      let quantidade = parseInt(valor);
      if (isNaN(quantidade) || quantidade < 0) {
        quantidade = 0;
      } else if (quantidade > item.SALDO) {
        quantidade = item.SALDO;
      }
      item.quantity = quantidade;
    }
  
    // Função para lidar com o evento input
    function handleInput(event: Event, item: ItemNF) {
      const target = event.target as HTMLInputElement;
      definirQuantidade(item, target.value);
    }
  </script>
  
  <!-- Exibindo a lista de itens -->
  <div class="p-4">
    <h3 class="text-xl font-bold text-primary mb-6 flex items-center">
      <IconTagsFilled class="mr-2 text-primary" />
      Produtos da Nota Fiscal
    </h3>
  
    {#if itensSelecionados.length > 0}
      <div class="space-y-4">
        <!-- Contêiner que se torna rolável quando há mais de 2 itens -->
        <div
          class="list-container"
          class:overflow-y-auto={itensSelecionados.length > 2}
          class:max-h-64={itensSelecionados.length > 2}
        >
          <ul class="space-y-4">
            {#each itensSelecionados as item (item.D2_ITEM)}
              <li class="card bg-base-200 shadow-md">
                <div class="card-body p-4">
                  <div class="flex justify-between items-center">
                    <div>
                      <h2 class="card-title text-base-content">{item.B1_DESC}</h2>
                      <p class="text-sm text-base-content">Código: {item.D2_COD}</p>
                      <p class="text-sm text-base-content">Saldo: {item.SALDO}</p>
                    </div>
                    <div class="form-control w-24">
                      <label class="label">
                        <span class="label-text">Quantidade</span>
                      </label>
                      <input
                        type="number"
                        min="0"
                        max={item.SALDO}
                        value={item.quantity}
                        on:input={(e) => handleInput(e, item)}
                        class="input input-bordered text-center"
                        placeholder="0"
                        aria-label={`Quantidade do item ${item.D2_ITEM}`}
                      />
                    </div>
                  </div>
                </div>
              </li>
            {/each}
          </ul>
        </div>
      </div>
    {:else}
      <p class="text-neutral-content">Nenhum item encontrado.</p>
    {/if}
  
    {#if errorMessageFetch}
      <p class="text-error mt-4">{errorMessageFetch}</p>
    {/if}
  </div>
  