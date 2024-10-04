export async function itensFetching(
	endpoint: string,
	headers: Record<string, string>
  ): Promise<{ data: any[] }> {
	try {
	  const url = `http://rodoapp:8080/api/pneus/borracharia/listanfs`;
  
	  console.log('Fazendo requisição para:', url);
	  console.log('Headers:', headers);
  
	  const response = await fetch(url, {
		method: 'GET',
		headers: headers,
	  });
  
	  if (!response.ok) {
		const errorText = await response.text();
		console.error('Erro na resposta da API:', errorText);
		throw new Error(`Erro ao buscar itens: ${response.status} ${response.statusText} - ${errorText}`);
	  }
  
	  const data = await response.json();
	  console.log('Dados recebidos da API:', data);
  
	  return { data };
	} catch (error) {
	  console.error('Erro ao buscar os itens:', error);
	  throw error;
	}
  }
  