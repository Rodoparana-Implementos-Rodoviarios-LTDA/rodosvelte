import { writable } from 'svelte/store';
import { xmlDataStore, xmlItemsStore, fetchState, filiaisStore } from '$lib/stores/xmlStore'; // Importando filiaisStore
import type { ConexaoNFE } from '$lib/types/ConexaoNFE';

// Função para normalizar o CNPJ, removendo pontos, traços e barras
function normalizeCNPJ(cnpj: string): string {
	return cnpj.replace(/[^\d]/g, ''); // Remove tudo que não for número
}

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

		// Normalizar o CNPJ do destinatário para garantir que possamos fazer o cruzamento correto
		const cnpjDestinatario = normalizeCNPJ(result.cnpjDestinatario);

		// Obter as filiais da store
		let filiais = [];
		filiaisStore.subscribe(value => {
			filiais = value;
		})();

		// Encontrar a filial correspondente ao CNPJ do destinatário
		const filialCorrespondente = filiais.find(filial => normalizeCNPJ(filial.cnpjFilial) === cnpjDestinatario);

		// Se encontrar a filial correspondente, substitui o nome do destinatário pelo nome da filial
		if (filialCorrespondente) {
			result.nomeDestinatario = filialCorrespondente.filial.trim(); // Substitui o nome do destinatário
		}

		// Salva os dados gerais no xmlDataStore
		xmlDataStore.set(result);

		// Extrai os itens, adiciona o índice a cada item e salva no xmlItemsStore
		if (result.itens && Array.isArray(result.itens)) {
			const itensComIndice = result.itens.map((item, index) => ({
				...item,
				index: index + 1 // Adiciona o índice, começando em 1
			}));

			// Salva os itens com o índice no xmlItemsStore
			xmlItemsStore.set(itensComIndice);
			console.log('Itens da XML com índice salvos no xmlItemsStore com sucesso.');
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
