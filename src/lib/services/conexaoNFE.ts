import type { DetalhesXML, Filial } from '$lib/types';
import { getOptionsData, saveOptionsData, deleteData } from '$lib/services/idb'; // Corrigido para usar getOptionsData

export async function fetchDetalhesXML(
	xml: string
): Promise<DetalhesXML & { filialName?: string }> {
	const token = import.meta.env.VITE_API_TOKEN; // Obtenha o token do .env.local
	const filiaisApiUrl = import.meta.env.VITE_FILIAIS_API_URL; // Obtenha a URL da API de filiais do .env.local
	const url = `https://api.conexaonfe.com.br/v1/dfes/${xml}/detalhes/xml`;

	try {
		// Exclui os dados antigos com base na chave 'xml'
		console.log(`Excluindo dados antigos da chave '${xml}' na tabela 'xml'...`);
		await deleteData('xml', xml); // Exclui os dados anteriores da tabela 'xml'

		// 1. Fetch detalhes do XML da API
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
			// Salva os dados mesmo que a filial não seja encontrada
			await saveOptionsData('xml', xml, data);
			
			// Emite o evento de que os dados foram atualizados
			const event = new CustomEvent('xmlDataUpdated', { detail: { xml } });
			window.dispatchEvent(event); // Emite o evento globalmente
			console.log('Evento xmlDataUpdated emitido!');
			
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

		// 5. Estruturar um novo objeto com os dados detalhados
		const structuredData = {
			...data,
			filialName: filialEncontrada ? filialEncontrada.filial : null
		};

		// 6. Salvar os dados estruturados na tabela fixa `xml`
		await saveOptionsData('xml', xml, structuredData); // Salva na tabela 'xml' com a chave xmlKey
		console.log(`Dados salvos na tabela 'xml' com a chave ${xml}`);

		// Emite o evento de que os dados foram atualizados
		const event = new CustomEvent('xmlDataUpdated', { detail: { xml } });
		window.dispatchEvent(event); // Emite o evento globalmente
		console.log('Evento xmlDataUpdated emitido!');

		return structuredData;
	} catch (error) {
		console.error('Erro ao buscar detalhes do XML:', error);

		// Se falhar a chamada de API, verifica se os dados estão no IndexedDB usando getOptionsData
		const cachedData = await getOptionsData('xml', xml);
		if (cachedData) {
			console.warn('Retornando dados do cache devido ao erro.');
			return cachedData;
		}

		throw error; // Se não houver dados em cache, lança o erro
	}
}
