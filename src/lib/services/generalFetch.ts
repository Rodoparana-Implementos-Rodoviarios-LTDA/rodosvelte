// src/lib/fetchAndLoadData.ts
import type { Filial, UnidadeMedida, Condicao, CentoCusto, ApiResponse } from '../../app';
import { getData, saveData } from './idb'; // Ajuste o caminho conforme necessário

export async function fetchAndLoadData(): Promise<{
	Filiais: Filial[];
	UnidadeMedida: UnidadeMedida[];
	Condicoes: Condicao[];
	centroCusto: CentoCusto[];
	tipoDeNFOptions: { label: string; value: string }[];
	prioridadeOptions: { label: string; value: string }[];
}> {
	const isClient = typeof window !== 'undefined';
	if (!isClient) {
		throw new Error('IndexedDB não está disponível no server-side.');
	}
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
	try {
		const [filiais, unidadeMedida, condicoes, centroCusto, tipoNFOptions, prioridadeOptions] =
			await Promise.all([
				getData('Filiais'),
				getData('UnidadeMedida'),
				getData('Condicoes'),
				getData('centroCusto'),
				getData('tipoDeNFOptions'),
				getData('prioridadeOptions')
			]);

		// Verifica se todos os dados estão presentes no IndexedDB
		if (
			filiais &&
			unidadeMedida &&
			condicoes &&
			centroCusto &&
			tipoNFOptions &&
			prioridadeOptions
		) {
			console.log('Dados recuperados do IndexedDB.');
			console.log('Filiais:', filiais);
			console.log('UnidadeMedida:', unidadeMedida);
			console.log('Condicoes:', condicoes);
			console.log('centroCusto:', centroCusto);
			console.log('tipoDeNFOptions:', tipoNFOptions);
			console.log('prioridadeOptions:', prioridadeOptions);

			return {
				Filiais: filiais as Filial[],
				UnidadeMedida: unidadeMedida as UnidadeMedida[],
				Condicoes: condicoes as Condicao[],
				centroCusto: centroCusto as CentoCusto[],
				tipoDeNFOptions: tipoNFOptions as { label: string; value: string }[],
				prioridadeOptions: prioridadeOptions as { label: string; value: string }[]
			};
		} else {
			console.log('Dados ausentes no IndexedDB. Buscando da API...');
			// Faz o fetch dos dados da API
			const response = await fetch(
				'http://protheus-app:8400/rest/reidoapsdu/consultar/cargaInicio'
			);
			if (!response.ok) {
				throw new Error(`Erro na requisição: ${response.statusText}`);
			}

			const data: ApiResponse = await response.json();

			const mappedData = {
				Filiais: data.Filiais,
				UnidadeMedida: data.UnidadeMedida,
				Condicoes: data.Condicoes,
				centroCusto: data.CentoCusto,
				tipoDeNFOptions,
				prioridadeOptions
			};

			// Armazena cada conjunto de dados no IndexedDB
			await Promise.all([
				saveData('Filiais', mappedData.Filiais),
				saveData('UnidadeMedida', mappedData.UnidadeMedida),
				saveData('Condicoes', mappedData.Condicoes),
				saveData('centroCusto', mappedData.centroCusto),
				saveData('tipoDeNFOptions', mappedData.tipoDeNFOptions),
				saveData('prioridadeOptions', mappedData.prioridadeOptions)
			]);

			console.log('Dados armazenados no IndexedDB.');
			console.log('Filiais salvas:', mappedData.Filiais);
			console.log('UnidadeMedida salvas:', mappedData.UnidadeMedida);
			console.log('Condicoes salvas:', mappedData.Condicoes);
			console.log('centroCusto salvas:', mappedData.centroCusto);
			console.log('tipoDeNFOptions salvas:', mappedData.tipoDeNFOptions);
			console.log('prioridadeOptions salvas:', mappedData.prioridadeOptions);

			return mappedData;
		}
	} catch (error) {
		console.error('Erro ao carregar os dados:', error);
		throw error;
	}
}
