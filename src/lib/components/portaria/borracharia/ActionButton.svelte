<script lang="ts">
    import { fly } from 'svelte/transition';
    import { circOut } from 'svelte/easing';
    import { IconCopyPlus, IconTruckReturn, IconUserDollar, IconUser, IconCar } from '@tabler/icons-svelte';
    import ProductItens from './productItens.svelte'; // Componente de Itens
    import type { ItemNF } from '$lib/types/tableTypes'; // Tipagem dos itens

    import toast, { Toaster } from 'svelte-french-toast'; // Importando toast e Toaster

    export let documentoCompleto: string;
    export let clienteCompleto: string;
    export let filial: string;
    export let observacao: string | null = '';
    export let usuario: string = 'matheus';

    let responsible = '';
    let plate = '';
    let retiradaPor = 'C'; // Valor padrão para o radio (Cliente)
    let isLoading = false;
    let drawerVisible = false;

    let itensSelecionados: ItemNF[] = []; // Armazena os itens selecionados

    // Função para abrir o drawer
    function openDrawer() {
        drawerVisible = true;
    }

    // Função para enviar as solicitações via POST
    async function handleConfirm() {
        console.log('handleConfirm - Início da função');

        if (!responsible || !plate) {
            toast.error('🚨 Preencha todos os campos obrigatórios!', { className: 'bg-error text-white' });  // Erro estilizado
            console.error('Campos obrigatórios faltando - responsável ou placa.');
            return;
        }

        const { documento, serie } = separarDocumentoESerie(documentoCompleto);
        const { cliente, loja } = separarClienteLoja(clienteCompleto);

        const itensParaIncluir = itensSelecionados.filter(item => item.quantity > 0);
        if (itensParaIncluir.length === 0) {
            toast.error('❗ Nenhum item selecionado para inclusão!', { className: 'bg-warning text-white' });  // Alerta estilizado
            console.error('Nenhum item selecionado para inclusão.');
            return;
        }

        try {
            isLoading = true;

            const promises = itensParaIncluir.map(async (item) => {
                const body = {
                    filial: filial.trim(),
                    documento: documento.trim(),
                    serie: serie.trim(),
                    cliente: cliente.trim(),
                    loja: loja.trim(),
                    produto: item.D2_COD.trim(),
                    item: item.D2_ITEM.trim(),
                    quantidade: item.quantity,
                    retiradopor: retiradaPor,
                    responsavel: responsible.trim(),
                    placa: plate.trim(),
                    observacoes: observacao?.trim() || '',
                    usuario: usuario.trim(),
                    origem: 'S'
                };

                console.log('Enviando POST com o seguinte corpo:', JSON.stringify(body, null, 2));

                const response = await fetch('http://protheus-vm:9010/rest/MovPortaria/IncluirItem', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body),
                });

                console.log('Status da resposta do POST:', response.status);

                if (!response.ok) {
                    console.error(`Erro na inclusão do item ${item.D2_ITEM}`);
                    throw new Error(`Erro ao incluir item ${item.D2_ITEM}`);
                }

                console.log(`Item ${item.D2_ITEM} incluído com sucesso.`);
                return null;
            });

            await Promise.all(promises);

            // Se não houver erros, todos os itens foram incluídos com sucesso
            toast.success('🎉 Todos os itens foram incluídos com sucesso!', { className: 'bg-success text-white' });  // Sucesso estilizado
            drawerVisible = false; // Fecha o drawer

        } catch (error) {
            console.error('Erro na inclusão:', error);
            toast.error(`❌ Erro ao incluir itens: ${error.message}`, { className: 'bg-error text-white' });  // Mensagem de erro estilizada
        } finally {
            isLoading = false;
        }
    }

    // Função para fechar o drawer
    function closeDrawer(event: Event) {
        const target = event.target as HTMLElement;
        if (!target.closest('.drawer-content')) {
            drawerVisible = false;
        }
    }

    // Separar documento/serie e cliente/loja
    function separarDocumentoESerie(docCompleto: string): { documento: string; serie: string } {
        const [documento, serie] = docCompleto.split(' - ').map(part => part.trim());
        return { documento: documento || '', serie: serie || '' };
    }

    function separarClienteLoja(clienteCompleto: string): { cliente: string; loja: string } {
        const [cliente, loja] = clienteCompleto.split(' - ').map(part => part.trim());
        return { cliente: cliente || '', loja: loja || '' };
    }
</script>

<!-- Toaster para exibir notificações -->
<Toaster />

<!-- Botão que abre o Drawer -->
<button class="btn btn-primary" on:click={openDrawer}>
    <IconCopyPlus />
</button>

<!-- Drawer -->
{#if drawerVisible}
    <!-- Overlay que fecha o Drawer ao clicar fora -->
    <div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>

    <!-- Conteúdo do Drawer com animação fly -->
    <div class="drawer drawer-content" transition:fly={{ x: 400, easing: circOut }}>
        <div class="w-full h-screen bg-base-100 text-info p-4">
            <h2 class="text-xl font-bold mb-4 text-primary">Seleção de Produtos</h2>

            <!-- Componente ProductItens que exibe a lista de produtos -->
            <ProductItens {documentoCompleto} {clienteCompleto} {filial} bind:itensSelecionados />

            <!-- Radio Buttons para Cliente e Rodoparaná -->
            <div class="mb-4 flex flex-col space-y-2">
                <label class="flex items-center cursor-pointer">
                    <IconUserDollar class="mr-2 text-primary" />
                    <input type="radio" name="retiradaPor" value="C" bind:group={retiradaPor} class="radio radio-base-content ml-2" />
                    <span class="label-text text-primary text-base ml-2">Cliente</span>
                </label>

                <label class="flex items-center cursor-pointer">
                    <IconTruckReturn class="mr-2 text-primary" />
                    <input type="radio" name="retiradaPor" value="R" bind:group={retiradaPor} class="radio radio-base-content ml-2" />
                    <span class="label-text text-primary text-base ml-2">Rodoparaná</span>
                </label>
            </div>

            <!-- Inputs para o nome do responsável e a placa -->
            <div class="mb-4 space-y-2">
                <div class="relative">
                    <IconUser class="absolute left-3 top-3 text-primary" />
                    <input type="text" placeholder="Nome do responsável" bind:value={responsible} class="input input-bordered w-full pl-10" />
                </div>
                <div class="relative">
                    <IconCar class="absolute left-3 top-3 text-primary" />
                    <input type="text" placeholder="Placa do carro" bind:value={plate} class="input input-bordered w-full pl-10" />
                </div>
            </div>

            <!-- Campo de Observações -->
            <div class="mb-4">
                <textarea class="textarea textarea-bordered w-full" placeholder="Observações" bind:value={observacao}></textarea>
            </div>

            <!-- Botões de Ação -->
            <div class="flex justify-between mt-6">
                <button class="btn btn-primary" on:click={handleConfirm} disabled={isLoading}>
                    {#if isLoading} Enviando... {:else} Confirmar {/if}
                </button>
                <button class="btn btn-outline" on:click={() => (drawerVisible = false)}>Cancelar</button>
            </div>
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
