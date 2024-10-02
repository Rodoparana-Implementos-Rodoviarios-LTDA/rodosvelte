<script lang="ts">
	import Combobox from '$lib/components/ui/Combobox.svelte'; // Usando o caminho do combobox
	import { onMount } from 'svelte';
	import { createEventDispatcher } from 'svelte';

	export let ncmsh: string; // O NCM vindo da linha do produto
	export let selectedProduct: { B1_COD: string; B1_DESC: string } | null = null;

	const dispatch = createEventDispatcher();
	let searchTerm = ''; // O termo de busca digitado pelo usuário
	let options: { label: string; value: { B1_COD: string; B1_DESC: string } }[] = []; // Lista de opções formatadas para o combobox
	let isLoading = false; // Controle de carregamento

	// Função que busca os produtos da API com base no termo de busca ou NCM
	async function fetchProducts(query: string): Promise<void> {
		isLoading = true;
		try {
			const response = await fetch(
				`http://172.16.99.174:8400/rest/reidoapsdu/consultar/likeprod?search=${query}`,
				{
					method: 'GET',
					headers: {
						'Content-Type': 'application/json'
					}
				}
			);

			if (!response.ok) {
				throw new Error('Erro ao buscar produtos');
			}

			const data = await response.json();

			// Filtra os resultados para garantir que o NCM (B1_POSIPI) corresponda ao NCM da linha (ncmsh)
			const filteredData = data.filter((item: { B1_POSIPI: string }) => item.B1_POSIPI === ncmsh);

			// Formata as opções para que o label seja o B1_DESC
			options = filteredData.map((item: { B1_COD: string; B1_DESC: string }) => ({
				label: item.B1_DESC, // Define o label como o B1_DESC
				value: item // O valor será o objeto completo
			}));
		} catch (error) {
			console.error('Erro ao buscar produtos:', error);
			options = [];
		} finally {
			isLoading = false;
		}
	}

	// Função chamada quando o valor no combobox muda
	function handleSelect(event: CustomEvent) {
		selectedProduct = event.detail.selectedProduct;
		dispatch('select', { selectedProduct });
	}

	// Busca inicial quando o componente é montado (com NCM)
	onMount(() => {
		if (ncmsh) {
			fetchProducts(ncmsh); // Busca os produtos relacionados ao NCM se estiver definido
		}
	});
</script>

<!-- Componente do Combobox -->
<Combobox
	{options}
	placeholder="Buscar produto por código ou descrição..."
	disabled={isLoading}
	bind:value={searchTerm}
	on:select={handleSelect}
/>

{#if isLoading}
	<p>Carregando...</p>
{:else if !options.length && searchTerm.length > 2}
	<p>Nenhum produto encontrado</p>
{/if}
