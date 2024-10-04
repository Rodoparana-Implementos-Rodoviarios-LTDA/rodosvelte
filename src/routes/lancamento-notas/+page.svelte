<script lang="ts">
	import Table from '$lib/components/ui/tabela/Table.svelte'; // Componente da Tabela
	import { columns } from '$lib/components/portaria/itensConferidos/columns'; // Importa as colunas atualizadas
	import Filtrar from '$lib/components/ui/tabela/Filtrar.svelte';
	import { onMount } from 'svelte';
	import { IconChevronDown } from '@tabler/icons-svelte';
	import type { CarregamentoSaida } from '$lib/types/tableTypes'; // Importa a interface atualizada
	
	let filters = {}; // Inicialmente, nenhum filtro aplicado
	let data: CarregamentoSaida[] = []; // Variável para armazenar os dados da API
	let isLoading = true; // Estado de carregamento
	let errorMessageFetch = ''; // Mensagem de erro
	
	// Função para buscar os dados da API
	async function fetchData() {
	  // Defina os valores de acordo com sua lógica de aplicação
	  const filial = '0101'; // Defina a filial conforme necessário
	  const documentoCompleto = ''; // Defina o documento completo conforme necessário (se aplicável)
	  const clienteCompleto = ''; // Defina o cliente completo conforme necessário (se aplicável)
	  
	  // Funções para separar documento/série e cliente/loja, se necessário
	  function separarDocumentoESerie(docCompleto: string): { documento: string; serie: string } {
		if (!docCompleto) return { documento: '', serie: '' };
		const partes = docCompleto.split(' - ').map((part) => part.trim());
		const documento = partes[0] || '';
		const serie = partes[1] || '';
		return { documento, serie };
	  }
	  
	  function separarClienteLoja(clienteCompleto: string): { cliente: string; loja: string } {
		if (!clienteCompleto) return { cliente: '', loja: '' };
		const partes = clienteCompleto.split(' - ').map((part) => part.trim());
		const cliente = partes[0] || '';
		const loja = partes[1] || '';
		return { cliente, loja };
	  }
	  
	  const { documento, serie } = separarDocumentoESerie(documentoCompleto);
	  const { cliente, loja } = separarClienteLoja(clienteCompleto);
	  
	  const filialTrimmed = filial?.trim();
	  if (!filialTrimmed || !/^\d{4}$/.test(filialTrimmed)) {
		errorMessageFetch = `Filial inválida: Deve conter exatamente 4 dígitos numéricos. Valor atual: '${filialTrimmed}'`;
		console.error(errorMessageFetch);
		isLoading = false;
		return;
	  }
	  
	  // Definir os headers conforme os requisitos da API
	  const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		// 'Authorization': 'Bearer SEU_TOKEN_AQUI', // Inclua o token se necessário
		'X-Page': '1',
		'X-Page-Size': '50', // Ajuste conforme necessário
		'X-Sort-By': '',
		'X-Sort-Order': '',
		'X-Filter-Filial': filialTrimmed,
		// Adicione outros headers de filtro conforme necessário
	  };
	  
	  // Adicionar filtros dinâmicos aos headers, se houver
	  for (const [key, value] of Object.entries(filters)) {
		headers[`X-Filter-${key}`] = value;
	  }
	  
	  console.log('Headers da requisição:', headers);
	  
	  const url = `http://rodoapp:8080/api/pneus/historico`;
	  
	  try {
		isLoading = true;
		const response = await fetch(url, {
		  method: 'GET',
		  headers: headers,
		});
		
		console.log('Status da resposta:', response.status);
		
		if (!response.ok) {
		  const errorText = await response.text();
		  console.error('Erro na resposta da API:', errorText);
		  throw new Error(`Erro ao buscar dados: ${response.status} ${response.statusText} - ${errorText}`);
		}
		
		const responseData = await response.json();
		console.log('Dados recebidos da API:', responseData);
		
		if (!responseData || responseData.length === 0) {
		  console.warn('Nenhum item foi retornado pela API.');
		}
		
		// Processar os dados recebidos e atualizar a variável 'data'
		data = responseData.map((item: any) => ({
		  Filial: item.Filial?.trim(),
		  NF: item.NF?.trim(),
		  Vendedor: item.Vendedor?.trim(),
		  Cliente: item.Cliente?.trim(),
		  Produto: item.Produto?.trim(),
		  DataHora: item.DataHora?.trim(),
		  Responsavel: item.Responsavel?.trim(),
		  Placa: item.Placa?.trim(),
		  Observacao: item.Observacao?.trim() || 'Sem Observação',
		  DataConf: item.DataConf?.trim(),
		  Seq: item.Seq?.trim(),
		  Saldo: Number(item.Saldo), // Garantir que Saldo é um número
		})) || [];
		
	  } catch (error: any) {
		console.error('Erro ao buscar dados da API:', error);
		errorMessageFetch = `Erro ao carregar dados: ${error.message}`;
	  } finally {
		isLoading = false;
	  }
	}
	
	// Função para aplicar filtros
	function applyFilters(event: CustomEvent<Record<string, string>>) {
	  filters = event.detail; // Captura os filtros do evento
	  fetchData(); // Recarrega os dados com os novos filtros
	}
	
	// Função para resetar os filtros
	function resetFilters() {
	  filters = {}; // Reseta os filtros
	  fetchData(); // Recarrega os dados sem filtros
	}
	
	// Carregar os dados quando o componente é montado
	onMount(() => {
	  fetchData(); // Chama a função para obter os dados ao montar o componente
	});
  </script>
  
  <!-- Interface da Tabela -->
  <div class="h-full">
	<!-- Cabeçalho com o título centralizado -->
	<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
	  <h1 class="text-3xl font-bold">Itens Conferidos - Lista de Produtos</h1>
	</div>
  
	<!-- Área dos Filtros e Exportação -->
	<div class="flex flex-col items-end">
	  <div class="flex justify-end space-x-4 pb-5">
		<div class="dropdown dropdown-end">
		  <div tabindex="0" role="button" class="btn btn-outline m-1"><IconChevronDown /></div>
		  <ul
			tabindex="0"
			class="dropdown-content menu bg-base-100 rounded-box z-[1] border border-secondary w-52 p-2 shadow"
		  >
			<li><a class="text-lg" href="/controle-itens">Borracharia</a></li>
			<li>
			  <a class="text-lg" href="/controle-itens/conferencia">Conferência de Saída</a>
			</li>
		  </ul>
		</div>
  
		<!-- Componente de Filtros -->
		<Filtrar {columns} on:applyFilters={applyFilters} on:resetFilters={resetFilters} />
	  </div>
  
	  <!-- Mensagem de erro -->
	  {#if errorMessageFetch}
		<p class="text-error mt-4">{errorMessageFetch}</p>
	  {/if}
  
	  <!-- Mensagem de carregamento -->
	  {#if isLoading}
		<p>Carregando...</p>
	  {:else}
		<!-- Componente da Tabela com dados carregados diretamente -->
		<Table {columns} {data} />
	  {/if}
	</div>
  </div>
  