import type { BorrachariaData } from "$lib/types/tableTypes";

// Função de fetch paginada
export const dataFetching = async (
  endpoint: string,             // Endpoint dinâmico para a API
  page: number = 1,             // Número da página atual
  pageSize: number = 10,        // Tamanho da página
  sortBy: string = '',          // Campo para ordenar
  sortOrder: string = 'asc',    // Ordem de classificação
  filters: Record<string, string> = {}  // Filtros opcionais
): Promise<{ data: BorrachariaData[], hasMore: boolean }> => {
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

    // Usa o endpoint passado por argumento
    const response = await fetch(`http://rodoapp:8080/api/${endpoint}?${params.toString()}`, config);

    // Verifique se a resposta da API é válida
    if (!response.ok) {
      throw new Error(`Erro na requisição da página ${page}: ${response.status}`);
    }

    const result = await response.json();

    // Adicione um log para depurar a resposta da API
    console.log('Resposta da API:', result);

    // Mapear o resultado para BorrachariaData[]
    const mappedData: BorrachariaData[] = result.map((item: any) => ({
      Filial: item.Filial,
      NF: item.NF,
      Vendedor: item.Vendedor || '',
      Cliente: item.Cliente || '',
      Produto: item.Produto || '',
      Saldo: item.Saldo || 0,
      Emissao: item.Emissao || item.DataHora ||'', // assuming 'DataHora' is the emission date
      Responsavel: item.Responsavel || '',
      Placa: item.Placa || '',
      Observacao: item.Observacao || ''
    }));

    // Retorna os dados mapeados e verifica se há mais páginas
    return {
      data: mappedData,
      hasMore: result.length === pageSize  // Se há mais dados a serem carregados
    };

  } catch (error) {
    console.error("Erro ao buscar os dados:", error);
    return { data: [], hasMore: false };  // Retorna sem dados se houver erro
  }
};
