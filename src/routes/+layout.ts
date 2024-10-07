// src/routes/+layout.ts
import { fetchAndUpdateData } from '$lib/services/fetchCargaInicio';
import { condicoesStore, unidadeMedidaStore, centoCustoStore } from '$lib/stores/xmlStore';

export const load = async () => {
  let condicoes = [];
  let unidadeMedida = [];
  let centoCusto = [];

  // Verifica o estado atual das stores
  condicoesStore.subscribe(value => condicoes = value);
  unidadeMedidaStore.subscribe(value => unidadeMedida = value);
  centoCustoStore.subscribe(value => centoCusto = value);

  // Se as stores estão vazias, faz o fetch e atualiza os dados
  if (condicoes.length === 0 || unidadeMedida.length === 0 || centoCusto.length === 0) {
    await fetchAndUpdateData();
  }

  return {}; // Retorna qualquer dado necessário para a página
};
