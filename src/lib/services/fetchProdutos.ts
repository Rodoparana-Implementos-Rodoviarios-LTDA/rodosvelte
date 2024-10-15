import { get } from 'svelte/store';
import { selectedTipoNF, produtosStore, fetchState } from '$lib/stores/xmlStore';
import type { ComboboxOption } from '$lib/types/Produtos';

export async function fetchProdutos(): Promise<void> {
	// Get the token from sessionStorage
	const token = sessionStorage.getItem('token');

	// Check if the token is present
	if (!token) {
		console.error('Token não encontrado na sessão.');
		fetchState.set('erro');
		return;
	}

	// Get the Tipo de NF (object with { value, label }) from the store
	const tipoNFSelecionado = get(selectedTipoNF);

	console.log('Tipo de NF selecionado (value):', tipoNFSelecionado);

	// Configure request headers
	const myHeaders = new Headers();
	myHeaders.append('Content-Type', 'application/json');
	myHeaders.append('Authorization', `Bearer ${token}`);
	myHeaders.append('Tipo', tipoNFSelecionado); // Send the value in the header

	const requestOptions: RequestInit = {
		method: 'GET',
		headers: myHeaders,
		redirect: 'follow'
	};

	try {
		// Update fetch state
		fetchState.set('carregando');

		// Fetch products from the API
		const response = await fetch('http://rodoapp:8080/api/produtos', requestOptions);

		if (!response.ok) {
			throw new Error(`Erro na requisição: ${response.status} ${response.statusText}`);
		}

		const result = await response.json();
		console.log('Produtos recebidos da API:', result);

		// Map the products
		const mappedProdutos: ComboboxOption[] = result.map((produto: any) => ({
			label: produto.B1_DESC,
			value: produto.B1_COD,
			campo1: produto.B1_POSIPI,
			campo2: produto.B1_ORIGEM
		}));

		// Save the products to the store
		produtosStore.set(mappedProdutos);
		console.log('produtos mapeados:', mappedProdutos);
		// Update fetch state to success
		fetchState.set('sucesso');
	} catch (error) {
		console.error('Erro ao buscar os produtos:', error);
		fetchState.set('erro');
	}
}
