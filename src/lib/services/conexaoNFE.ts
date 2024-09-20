// src/lib/services/conexaoNFE.ts
import type { DetalhesXML, Item, Filial } from '$lib/types/tableTypes'; // Ajuste o caminho conforme necessário

export async function fetchDetalhesXML(
	xml: string
): Promise<DetalhesXML & { filialName?: string }> {
	const token = import.meta.env.VITE_API_TOKEN; // Obtenha o token do .env.local
	const filiaisApiUrl = import.meta.env.VITE_FILIAIS_API_URL; // Obtenha a URL da API de filiais do .env.local
	const url = `https://api.conexaonfe.com.br/v1/dfes/${xml}/detalhes/xml`;

	try {
		// 1. Fetch detalhes do XML
		const response = await fetch(url, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${token}`,
				'Content-Type': 'application/json'
			}
		});

		if (!response.ok) {
			throw new Error(`Erro na requisição de detalhes do XML: ${response.statusText}`);
		}

		const data: DetalhesXML = await response.json();
		console.log('Dados recebidos da API de detalhes do XML:', data);

		// 2. Extrai o CNPJ do destinatário
		const cnpjDestinatario = data.cnpjDestinatario?.trim();
		if (!cnpjDestinatario) {
			console.warn('CNPJ do destinatário não encontrado nos detalhes do XML.');
			return data;
		}
		console.log('CNPJ Destinatário:', cnpjDestinatario);

		// 3. Fetch lista de filiais
		const filiaisResponse = await fetch(filiaisApiUrl, {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		});

		if (!filiaisResponse.ok) {
			throw new Error(`Erro na requisição de filiais: ${filiaisResponse.statusText}`);
		}

		const filiais: Filial[] = await filiaisResponse.json();
		console.log('Filiais recebidas da API:', filiais);

		// 4. Triangular o CNPJ do destinatário com a lista de filiais
		const filialEncontrada = filiais.find((filial) => {
			const cnpjFilial = filial.cnpjFilial?.trim();
			return cnpjFilial === cnpjDestinatario;
		});

		if (filialEncontrada) {
			data.filialName = filialEncontrada.filial.trim();
			console.log('Filial encontrada:', filialEncontrada.filial.trim());
		} else {
			console.warn(`Nenhuma filial encontrada para o CNPJ: ${cnpjDestinatario}`);
		}

		return data;
	} catch (error) {
		console.error('Erro ao buscar detalhes do XML:', error);
		throw error;
	}
}
