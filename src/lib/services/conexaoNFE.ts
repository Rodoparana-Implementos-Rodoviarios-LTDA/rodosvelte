// src/lib/services/conexaoNFE.ts
import type { DetalhesXML, Filial } from '$lib/types'; // Ajuste o caminho conforme necessário
import { saveData } from '$lib/services/idb'; // Ajuste o caminho para onde a função saveData está localizada

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
            // Salva os dados mesmo que a filial não seja encontrada
            await saveData(xml, data);
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

        // 5. Salvar os dados no IndexedDB usando saveData
        await saveData(xml, data);
        console.log('Dados salvos no IndexedDB com sucesso.');

        return data;
    } catch (error) {
        console.error('Erro ao buscar detalhes do XML:', error);
        throw error;
    }
}
