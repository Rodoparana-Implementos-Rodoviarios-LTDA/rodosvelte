import type { PreNota } from "$lib/types/tableTypes";

// Função de fetch paginada
export const dataFetching = async (
  page: number = 1,             // Número da página atual
  pageSize: number = 10,        // Tamanho da página
  sortBy: string = '',          // Campo para ordenar
  sortOrder: string = 'asc',    // Ordem de classificação
  filters: Record<string, string> = {}  // Filtros opcionais
): Promise<{ data: PreNota[], hasMore: boolean }> => {
  try {
    const token = sessionStorage.getItem('token');
    if (!token) {
      throw new Error("Token não encontrado no sessionStorage");
    }

    // Construa a URL com os parâmetros de paginação e filtros
    const params = new URLSearchParams({
      page: String(page),
      pageSize: String(pageSize),
      sortBy,
      sortOrder,
      ...filters,  // Filtros dinâmicos
    });

    const config = {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,  // Token no header
      },
    };

    const response = await fetch(`http://localhost:8080/api/prenotas?${params.toString()}`, config);

    // Verifique se a resposta da API é válida
    if (!response.ok) {
      throw new Error(`Erro na requisição da página ${page}: ${response.status}`);
    }

    const result = await response.json();

    // Adicione um log para depurar a resposta da API
    console.log('Resposta da API:', result);

    // Verifica se a resposta é um array
    if (!Array.isArray(result)) {
      throw new Error('Resposta inesperada: o resultado não é um array.');
    }

    // Retorna os dados da página atual e verifica se há mais páginas a carregar
    return {
      data: result,         // Os registros são o próprio array retornado
      hasMore: result.length === pageSize  // Se há mais dados a serem carregados
    };

  } catch (error) {
    console.error("Erro ao buscar os dados:", error);
    return { data: [], hasMore: false };  // Retorna sem dados se houver erro
  }
};
