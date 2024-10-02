<!-- src/lib/components/ProdutoSelector.svelte -->
<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { onMount } from 'svelte';
    import type { ProdutoProtheus } from '$lib/types';

    export let selectedProduto: ProdutoProtheus | null = null;
    export let produtosProtheus: ProdutoProtheus[] = [];

    const dispatch = createEventDispatcher();

    let searchTerm: string = '';
    let filteredProdutos: ProdutoProtheus[] = [];

    onMount(() => {
        filteredProdutos = produtosProtheus;
    });

    function handleSearch() {
        if (searchTerm.trim() === '') {
            filteredProdutos = produtosProtheus;
        } else {
            const lowerSearch = searchTerm.toLowerCase();
            filteredProdutos = produtosProtheus.filter(produto =>
                produto.descricao.toLowerCase().includes(lowerSearch) ||
                produto.codigo.toLowerCase().includes(lowerSearch)
            );
        }
    }

    function selectProduto(produto: ProdutoProtheus) {
        selectedProduto = produto;
        dispatch('select', produto);
    }
</script>

<style>
    .dropdown {
        position: relative;
        display: inline-block;
        width: 100%;
    }

    .dropdown-content {
        position: absolute;
        background-color: white;
        min-width: 100%;
        max-height: 200px;
        overflow-y: auto;
        box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
        z-index: 1;
    }

    .dropdown-content div {
        padding: 8px;
        cursor: pointer;
    }

    .dropdown-content div:hover {
        background-color: #f1f1f1;
    }
</style>

<div class="dropdown">
    <input
        type="text"
        placeholder="Buscar produto do Protheus..."
        bind:value={searchTerm}
        on:input={handleSearch}
        class="input input-bordered w-full"
    />
    {#if searchTerm.trim() !== ''}
        <div class="dropdown-content">
            {#each filteredProdutos as produto (produto.codigo)}
                <div on:click={() => selectProduto(produto)}>
                    {produto.codigo} - {produto.descricao}
                </div>
            {/each}
            {#if filteredProdutos.length === 0}
                <div>Nenhum produto encontrado.</div>
            {/if}
        </div>
    {/if}
</div>
