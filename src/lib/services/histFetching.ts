import type { HistoricoData } from '$lib/types/tableTypes';  // Certifique-se de que o tipo está correto

export async function fetchHistorico(
  page: number = 1,              // Página atual
  pageSize: number = 10,          // Quantidade de itens por página
  sortBy: string = 'DataHora',    // Campo de ordenação
  sortOrder: 'asc' | 'desc' = 'desc' // Ordem de classificação (ascendente/descendente)
): Promise<{ data: HistoricoData[]; hasMore: boolean }> {
  try {
    const response = await fetch(
      `http://rodoapp:8080/api/portaria?page=${page}&pageSize=${pageSize}&sortBy=${sortBy}&sortOrder=${sortOrder}`,
      {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        }
      }
    );

    if (!response.ok) {
      throw new Error(`Erro na requisição: ${response.status}`);
    }

    const result = await response.json();

    // Log da resposta completa da API para verificar o formato
    console.log('Resposta completa da API:', JSON.stringify(result, null, 2));

    // Verifica se o resultado é um array de dados
    if (Array.isArray(result)) {
      return {
        data: result.slice(0, pageSize),  // Limita ao tamanho da página
        hasMore: result.length === pageSize,  // Se tem mais dados para carregar
      };
    }

    // Se o resultado for um objeto que contém 'data' como array
    if (result.data && Array.isArray(result.data)) {
      return {
        data: result.data.slice(0, pageSize),  // Limita ao tamanho da página
        hasMore: result.hasMore || result.data.length === pageSize,  // Se tem mais dados para carregar
      };
    }

    // Caso o formato da resposta seja inesperado, lance um erro
    throw new Error('Formato de resposta inesperado.');
  } catch (error) {
    console.error('Erro ao buscar os dados da API:', error);
    return { data: [], hasMore: false };  // Retorna vazio em caso de erro
  }
}
