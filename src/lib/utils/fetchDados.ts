export interface DadosResponse<T> {
  data: T[];
  total: number;
  page: number;
  pageSize: number;
}

export async function fetchDados<T>(
  endpoint: string,
  token?: string, // Token opcional
  page: number = 1,
  pageSize: number = 20,
  sortBy: string = 'Inclusao',
  sortOrder: 'asc' | 'desc' = 'desc'
): Promise<DadosResponse<T>> {
  console.log('Iniciando fetchDados...');
  console.log('Endpoint:', endpoint);
  console.log('Token:', token ? 'Token presente' : 'Token ausente');
  console.log('Página:', page, 'Tamanho da Página:', pageSize);
  console.log('Ordenar por:', sortBy, 'Ordem:', sortOrder);

  const url = `/api/${endpoint}`;
  console.log('URL da requisição:', url);

  const headers = new Headers({
    "Content-Type": "application/json",
  });

  if (token && endpoint.includes('prenotas')) {
    console.log('Adicionando token ao cabeçalho de autorização...');
    headers.append("Authorization", `Bearer ${token}`);
  }

  // Adiciona headers de paginação e ordenação
  headers.append("X-Page", page.toString());
  headers.append("X-Page-Size", pageSize.toString());
  headers.append("X-Sort-By", sortBy);
  headers.append("X-Sort-Order", sortOrder);
  
  console.log('Headers montados:', headers);

  const requestOptions: RequestInit = {
    method: 'GET',
    headers,
    redirect: 'follow',
  };

  try {
    console.log('Fazendo requisição para:', url, 'com requestOptions:', requestOptions);

    const response = await fetch(url, requestOptions);
    console.log('Resposta recebida:', response);

    const contentType = response.headers.get('content-type');
    console.log('Content-Type da resposta:', contentType);

    if (!response.ok) {
      let errorDetails;
      console.log('Erro na resposta. Status:', response.status, response.statusText);

      if (contentType && contentType.includes('application/json')) {
        errorDetails = await response.json();
        console.log('Detalhes do erro (JSON):', errorDetails);
      } else {
        errorDetails = await response.text();
        console.log('Detalhes do erro (Texto):', errorDetails);
      }

      console.error('Erro completo:', response.status, response.statusText, errorDetails);
      throw new Error(`Erro ao buscar dados: ${response.status} - ${response.statusText}`);
    }

    const result: DadosResponse<T> = await response.json();
    console.log('Dados recebidos:', result);

    return result;

  } catch (error) {
    console.error('Erro ao realizar a requisição:', error);
    throw error;
  }
}
