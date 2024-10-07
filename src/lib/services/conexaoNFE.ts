import { writable } from 'svelte/store';
import { xmlDataStore, xmlItemsStore } from '$lib/stores/xmlStore';
import type { ConexaoNFE } from '$lib/types/ConexaoNFE';

// Store de estado para o carregamento
export const fetchState = writable<'inicial' | 'carregando' | 'sucesso' | 'erro'>('inicial');

/**
 * Função para buscar detalhes do ConexãoNFE da API e salvar no store.
 *
 * @param {string} xml - O identificador XML a ser utilizado na URL da API.
 * @returns {Promise<void>}
 */
export async function fetchAndSaveConexaoNFE(xml: string): Promise<void> {
	// Configuração dos Headers
	const myHeaders = new Headers();
	myHeaders.append(
		'Authorization',
		'Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJjb25leGFvbmZlLmNvbS5iciIsIm5hbWUiOiJtYXRldXMiLCJpZCI6NTYwNTF9.SDMelkA6zQz0BFtLb-bCH4y6t2pTxWyuI5Lr2Bu_YUo'
	);
	myHeaders.append('Content-Type', 'application/json');

	// Configuração das Opções da Requisição
	const requestOptions: RequestInit = {
		method: 'GET',
		headers: myHeaders,
		redirect: 'follow'
	};

	// Inicializa o estado como "carregando"
	fetchState.set('carregando');

	// Construção da URL com o parâmetro XML
	const conexaoNFEApiUrl = `https://api.conexaonfe.com.br/v1/dfes/${xml}/detalhes/xml`;

	try {
		// Realiza a requisição à API
		const response = await fetch(conexaoNFEApiUrl, requestOptions);

		if (!response.ok) {
			throw new Error(
				`Erro na requisição de ConexãoNFE: ${response.status} ${response.statusText}`
			);
		}

		// Extrai os dados da resposta
		const result: ConexaoNFE = await response.json();
		console.log('ConexãoNFE recebida da API:', result);

		// Salva os dados gerais no xmlDataStore
		xmlDataStore.set(result);

		// Extrai os itens e salva no xmlItemsStore
		if (result.itens && Array.isArray(result.itens)) {
			xmlItemsStore.set(result.itens);
			console.log('Itens da XML salvos no xmlItemsStore com sucesso.');
		} else {
			throw new Error('Nenhum item encontrado na resposta.');
		}

		// Define o estado como "sucesso"
		fetchState.set('sucesso');
	} catch (error) {
		console.error('Erro ao buscar ou salvar ConexãoNFE:', error);
		// Define o estado como "erro"
		fetchState.set('erro');
	}
}
