// src/lib/services/conexaoNFE.ts
export async function fetchDetalhesXML(xml: string): Promise<any> {
    const token = import.meta.env.VITE_API_TOKEN; // Obtenha o token do .env.local
    const url = `https://api.conexaonfe.com.br/v1/dfes/${xml}/detalhes/xml`;

    try {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${token}`,
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            throw new Error(`Erro na requisição: ${response.statusText}`);
        }

        const data = await response.json();
        console.log('Dados recebidos da API:', data); // Log para verificar os dados recebidos

        return data;
    } catch (error) {
        console.error('Erro ao buscar detalhes do XML:', error);
        throw error;
    }
}
