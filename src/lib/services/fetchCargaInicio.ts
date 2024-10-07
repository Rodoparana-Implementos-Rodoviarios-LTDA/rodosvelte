// src/lib/services/fetchCargaInicio.ts
import {
	unidadeMedidaStore,
	condicoesStore,
	centoCustoStore,
	filiaisStore
} from '$lib/stores/xmlStore';

/**
 * Função para buscar dados da API e atualizar os stores com os dados mais recentes.
 */
export async function fetchAndUpdateData(): Promise<void> {
	const cargaInicioApiUrl = 'http://172.16.99.174:8400/rest/reidoapsdu/consultar/cargaInicio';
	const filiaisApiUrl = 'http://protheus-app:8400/rest/reidoapsdu/consultar/filiais/';

	const requestOptions: RequestInit = {
		method: 'GET',
		headers: { 'Content-Type': 'application/json' },
		redirect: 'follow'
	};

	try {
		// Fetch cargaInicio
		const cargaInicioResponse = await fetch(cargaInicioApiUrl, requestOptions);
		if (!cargaInicioResponse.ok) {
			throw new Error(
				`Erro na requisição de cargaInicio: ${cargaInicioResponse.status} ${cargaInicioResponse.statusText}`
			);
		}
		const cargaInicioResult = await cargaInicioResponse.json();

		// Atualizar stores com os dados recebidos
		unidadeMedidaStore.set(cargaInicioResult.UnidadeMedida || []);
		condicoesStore.set(cargaInicioResult.Condicoes || []);
		centoCustoStore.set(cargaInicioResult.CentoCusto || []);

		// Fetch filiais
		const filiaisResponse = await fetch(filiaisApiUrl, requestOptions);
		if (!filiaisResponse.ok) {
			throw new Error(
				`Erro na requisição de filiais: ${filiaisResponse.status} ${filiaisResponse.statusText}`
			);
		}
		const filiaisResult = await filiaisResponse.json();

		// Atualizar store de filiais
		filiaisStore.set(filiaisResult || []);

		console.log('Dados atualizados com sucesso.');
	} catch (error) {
		console.error('Erro ao buscar ou atualizar os dados:', error);
	}
}
