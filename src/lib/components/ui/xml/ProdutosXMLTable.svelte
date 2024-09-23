<!-- src/lib/components/ui/xml/ProdutosXMLTable.svelte -->
<script lang="ts">
    import type { DetalhesXML, Item } from '$lib/types'; // Ajuste o caminho conforme necessário
    import { onMount, createEventDispatcher } from 'svelte';
    import { getData } from '$lib/services/idb'; // Função para buscar dados do IndexedDB

    // Props recebidas
    export let xmlKey: string;

    // Define as colunas da tabela
    const columns = [
        { header: 'Código XML', accessor: 'codProduto' },
        { header: 'Descrição XML', accessor: 'descProduto' },
        { header: 'Quantidade', accessor: 'quantidade' },
        { header: 'Valor Unitário', accessor: 'valorUnitario' },
        { header: 'Valor Total', accessor: 'valorTotal' },
        { header: 'Código Protheus', accessor: 'codigoProtheus' },
        { header: 'Descrição Protheus', accessor: 'descricaoProtheus' }
    ];

    // Estado local para armazenar os detalhes da XML
    let detalhesXML: DetalhesXML | null = null;
    let error: string | null = null;
    let isLoading: boolean = true;

    const dispatch = createEventDispatcher();

    // Função para carregar os dados da tabela a partir do IndexedDB
    async function loadXMLData() {
        if (!xmlKey.trim()) {
            detalhesXML = null;
            isLoading = false;
            return;
        }

        isLoading = true;
        error = null;

        try {
            const data: DetalhesXML = await getData(xmlKey);
            if (data) {
                detalhesXML = data;
            } else {
                console.warn('Nenhum dado encontrado para a chave XML fornecida.');
                detalhesXML = null;
            }
        } catch (err) {
            error = (err as Error).message;
            console.error('Erro ao recuperar dados do IndexedDB:', error);
            detalhesXML = null;
        } finally {
            isLoading = false;
        }
    }

    // Recarregar os dados sempre que a chave XML mudar
    $: if (xmlKey) {
        loadXMLData();
    }
</script>

<style>
    .empty-state {
        text-align: center;
        padding: 2rem;
        color: #9ca3af; /* Cor de texto padrão para estados vazios */
    }
</style>

<div class="overflow-x-auto h-full">
    <table class="table table-lg w-full h-full bg-base-300">
        <!-- Cabeçalho da Tabela -->
        <thead>
            <tr class="text-center text-lg">
                {#each columns as column}
                    <th>{column.header}</th>
                {/each}
            </tr>
        </thead>

        <!-- Corpo da Tabela -->
        <tbody>
            {#if isLoading}
                <tr>
                    <td colspan={columns.length} class="text-center">Carregando dados...</td>
                </tr>
            {:else if error}
                <tr>
                    <td colspan={columns.length} class="text-center text-red-500">Erro: {error}</td>
                </tr>
            {:else if detalhesXML && detalhesXML.itens.length > 0}
                {#each detalhesXML.itens as item, index}
                    <tr class="hover">
                        <td>{item.codProduto}</td>
                        <td>{item.descProduto}</td>
                        <td>{item.quantidade}</td>
                        <td>{item.valorUnitario}</td>
                        <td>{item.valorTotal}</td>
                        <td>{item.codigoProtheus || '-'}</td>
                        <td>{item.descricaoProtheus || '-'}</td>
                    </tr>
                {/each}
            {:else}
                <tr>
                    <td colspan={columns.length} class="empty-state">Nenhum dado disponível.</td>
                </tr>
            {/if}
        </tbody>
    </table>
</div>
