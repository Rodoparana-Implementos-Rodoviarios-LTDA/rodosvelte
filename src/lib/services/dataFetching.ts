export const dataFetching = async <T>(
	endpoint: string, // Endpoint dinâmico para a API
	sortBy: string,
	sortOrder: string,
	page: number = 1, // Número da página atual
	pageSize: number = 10, // Tamanho da página
	filters: Record<string, any> = {} // Novo parâmetro: filtros opcionais
): Promise<{ data: T[]; hasMore: boolean }> => {
	try {
		const token = sessionStorage.getItem('token');
		if (!token) {
			throw new Error('Token não encontrado no sessionStorage');
		}

		// Configurações da requisição
		const config = {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json',
				'X-Page': page.toString(), // Usa o número da página dinâmico
				'X-Page-Size': pageSize.toString(), // Tamanho da página dinâmico
				'X-Sort-By': sortBy, // Ordenação
				'X-Sort-Order': sortOrder, // Ordem de classificação
				Authorization: `Bearer ${token}` // Token no header
			}
		};

		// Usa o endpoint passado por argumento
		const response = await fetch(`/api/${endpoint}`, config);

		// Verifique se a resposta da API é válida
		if (!response.ok) {
			throw new Error(`Erro na requisição da página ${page}: ${response.status}`);
		}

		const result = await response.json();

		// Retorna os dados e verifica se há mais páginas
		return {
			data: result as T[], // Tipo genérico para os dados
			hasMore: result.length === pageSize // Verifica se há mais páginas
		};
	} catch (error) {
		console.error('Erro ao buscar os dados:', error);
		return { data: [], hasMore: false }; // Retorna sem dados se houver erro
	}
};
