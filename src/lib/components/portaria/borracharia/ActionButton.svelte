<script lang="ts">
    import { fly } from 'svelte/transition';
    import { circOut } from 'svelte/easing';
    import { IconCopyPlus, IconTruckReturn, IconUserDollar, IconUser, IconCar } from '@tabler/icons-svelte';

    // Importando o componente ProductItens
    import ProductItens from './productItens.svelte';

    // Importando o tipo ItemNF
    import type { ItemNF } from '$lib/types/tableTypes';

    // Recebendo as informações via props
    export let documentoCompleto: string;
    export let clienteCompleto: string;
    export let filial: string;
    export let observacao: string | null = '';
    export let usuario: string = 'matheus';

    // Variáveis de estado
    let responsible = '';
    let plate = '';
    let retiradaPor = 'C'; // Valor padrão para o radio (Cliente)
    let isLoading = false;
    let errorMessage = '';
    let drawerVisible = false;

    // Variável para armazenar os itens selecionados
    let itensSelecionados: ItemNF[] = [];

    // Abrir o drawer
    function openDrawer() {
        drawerVisible = true;
    }

    // Função para enviar as solicitações via POST
    async function handleConfirm() {
        console.log('handleConfirm - Início da função');

        // Validar campos obrigatórios
        if (!responsible || !plate) {
            errorMessage = 'Preencha todos os campos obrigatórios.';
            console.error(
                'Erro: Campos obrigatórios faltando - responsible:',
                responsible,
                'plate:',
                plate
            );
            return;
        }

        // Extrair documento e série
        const { documento, serie } = separarDocumentoESerie(documentoCompleto);
        const { cliente, loja } = separarClienteLoja(clienteCompleto);

        // Filtrar itens com quantidade > 0
        const itensParaIncluir = itensSelecionados.filter(item => item.quantity > 0);

        if (itensParaIncluir.length === 0) {
            errorMessage = 'Nenhum item selecionado para inclusão.';
            console.error('Erro: Nenhum item selecionado para inclusão.');
            return;
        }

        try {
            isLoading = true;
            errorMessage = ''; // Limpar mensagens de erro anteriores

            // Criar um array de promessas para enviar as requisições em paralelo
            const promises = itensParaIncluir.map(async (item) => {
                // Construir o corpo da solicitação para cada item
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

                console.log('Enviando POST com o seguinte corpo:');
                console.log(JSON.stringify(body, null, 2));

                try {
                    const response = await fetch('http://protheus-vm:9010/rest/MovPortaria/IncluirItem', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(body)
                    });

                    console.log('Status da resposta do POST:', response.status);

                    const responseData = await response.json();
                    console.log('Dados recebidos da API após o POST:', responseData);

                    if (!response.ok) {
                        console.error('Erro na resposta da API após o POST:', responseData);
                        // Retorna uma mensagem de erro específica para este item
                        return `Erro ao incluir item ${item.D2_ITEM}: ${responseData.mensagem || 'Erro ao enviar a solicitação.'}`;
                    }

                    console.log(`Item ${item.D2_ITEM} incluído com sucesso.`);
                    return null; // Indica sucesso para este item
                } catch (error) {
                    console.error(`Erro ao enviar o item ${item.D2_ITEM}:`, error);
                    return `Erro ao incluir item ${item.D2_ITEM}: ${error.message || 'Erro ao enviar a solicitação.'}`;
                }
            });

            // Aguarda todas as promessas serem resolvidas
            const results = await Promise.all(promises);

            // Verifica se houve erros
            const errors = results.filter(result => result !== null);
            if (errors.length > 0) {
                errorMessage = errors.join('\n');
                
            } else {
                errorMessage = '';
                // Todos os itens foram incluídos com sucesso
                alert('Todos os itens foram incluídos com sucesso!');
                drawerVisible = false; // Fecha o drawer
            }
        } catch (error: any) {
            console.error('Erro na requisição:', error);
            errorMessage = 'Erro ao enviar a solicitação.';
        } finally {
            isLoading = false;
        }
    }

    // Função para fechar o drawer ao clicar fora dele
    function closeDrawer(event: Event) {
        const target = event.target as HTMLElement;
        if (!target.closest('.drawer-content')) {
            drawerVisible = false;
        }
    }

    // Funções para separar documento/serie e cliente/loja
    function separarDocumentoESerie(docCompleto: string): { documento: string; serie: string } {
        if (!docCompleto) return { documento: '', serie: '' };
        const [documento, serie] = docCompleto.split(' - ').map((part) => part.trim());
        return { documento: documento || '', serie: serie || '' };
    }

    function separarClienteLoja(clienteCompleto: string): { cliente: string; loja: string } {
        if (!clienteCompleto) return { cliente: '', loja: '' };
        const [cliente, loja] = clienteCompleto.split(' - ').map((part) => part.trim());
        return { cliente: cliente || '', loja: loja || '' };
    }
</script>

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
                    <input
                        type="radio"
                        name="retiradaPor"
                        value="C"
                        bind:group={retiradaPor}
                        class="radio radio-base-content ml-2"
                    />
                    <span class="label-text text-primary text-base ml-2">Cliente</span>
                </label>

                <label class="flex items-center cursor-pointer">
                    <IconTruckReturn class="mr-2 text-primary" />
                    <input
                        type="radio"
                        name="retiradaPor"
                        value="R"
                        bind:group={retiradaPor}
                        class="radio radio-base-content ml-2"
                    />
                    <span class="label-text text-primary text-base ml-2">Rodoparaná</span>
                </label>
            </div>

            <!-- Inputs para o nome do responsável e a placa -->
            <div class="mb-4 space-y-2">
                <div class="relative">
                    <IconUser class="absolute left-3 top-3 text-primary" />
                    <input
                        type="text"
                        placeholder="Nome do responsável"
                        bind:value={responsible}
                        class="input input-bordered w-full pl-10"
                    />
                </div>
                <div class="relative">
                    <IconCar class="absolute left-3 top-3 text-primary" />
                    <input
                        type="text"
                        placeholder="Placa do carro"
                        bind:value={plate}
                        class="input input-bordered w-full pl-10"
                    />
                </div>
            </div>

            <!-- Campo de Observações -->
            <div class="mb-4">
                <textarea
                    class="textarea textarea-bordered w-full"
                    placeholder="Observações"
                    bind:value={observacao}
                ></textarea>
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
                <button class="btn btn-outline" on:click={() => (drawerVisible = false)}> Cancelar </button>
            </div>

            {#if errorMessage}
                <p class="text-error mt-4 whitespace-pre-line">{errorMessage}</p>
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
