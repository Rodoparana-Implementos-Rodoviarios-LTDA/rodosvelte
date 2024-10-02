import { saveOptionsData } from './idb';

// Arrays estáticos que serão salvos no IndexedDB
const tipoDeNFOptions = [
	{ label: 'Revenda', value: 'revenda' },
	{ label: 'Despesa/Imobilizado', value: 'despesa_imobilizado' },
	{ label: 'Matéria Prima', value: 'materia_prima' },
	{ label: 'Collection', value: 'collection' }
];

const prioridadeOptions = [
	{ label: 'Baixa', value: 'baixa' },
	{ label: 'Média', value: 'media' },
	{ label: 'Alta', value: 'alta' }
];

// Função para salvar os arrays estáticos no IndexedDB
async function saveStaticOptions(): Promise<void> {
	try {
		await saveOptionsData('tipoDeNFOptions', 'tipoDeNF', tipoDeNFOptions); // Atualizado com a chave e tabela
		await saveOptionsData('prioridadeOptions', 'prioridade', prioridadeOptions); // Atualizado com a chave e tabela
		console.log('Dados estáticos salvos com sucesso no IndexedDB.');
	} catch (error) {
		console.error('Erro ao salvar dados estáticos no IndexedDB:', error);
	}
}

// Função para buscar dados da API nova e salvar no IndexedDB
async function fetchAndSaveCargaInicio(): Promise<void> {
	try {
		const myHeaders = new Headers();
		myHeaders.append('Content-Type', 'application/json');

		const requestOptions: RequestInit = {
			method: 'GET',
			headers: myHeaders,
			redirect: 'follow'
		};

		const response = await fetch(
			'http://172.16.99.174:8400/rest/reidoapsdu/consultar/cargaInicio',
			requestOptions
		);

		if (!response.ok) {
			throw new Error(`Erro na requisição: ${response.status} ${response.statusText}`);
		}

		const data = await response.json();

		// Salva cada seção dos dados no IndexedDB
		if (data.Filiais) {
			await saveOptionsData('filiaisOptions', 'filiais', data.Filiais); // Atualizado com a chave e tabela
			console.log('Dados de Filiais salvos com sucesso no IndexedDB.');
		}

		if (data.UnidadeMedida) {
			await saveOptionsData('unidadeMedidaOptions', 'unidadeMedida', data.UnidadeMedida); // Atualizado com a chave e tabela
			console.log('Dados de Unidade de Medida salvos com sucesso no IndexedDB.');
		}

		if (data.Condicoes) {
			await saveOptionsData('condicoesOptions', 'condicoes', data.Condicoes); // Atualizado com a chave e tabela
			console.log('Dados de Condições salvos com sucesso no IndexedDB.');
		}

		if (data.CentoCusto) {
			await saveOptionsData('centroCustoOptions', 'centroCusto', data.CentoCusto); // Atualizado com a chave e tabela
			console.log('Dados de Centro de Custo salvos com sucesso no IndexedDB.');
		}
	} catch (error) {
		console.error('Erro ao buscar ou salvar dados de cargaInicio:', error);
	}
}

// Função principal que executa as outras funções
async function initializeData(): Promise<void> {
	await saveStaticOptions(); // Salva os dados estáticos
	await fetchAndSaveCargaInicio(); // Busca e salva os dados da nova API
}

// Chama a função de inicialização ao carregar o módulo
initializeData().catch((error) => console.error('Erro ao inicializar os dados:', error));

export { initializeData };
