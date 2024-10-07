// src/lib/stores/xmlStore.ts
import { writable } from 'svelte/store';
import type { ConexaoNFE, ConexaoNFEItem } from '$lib/types/ConexaoNFE';
import type { Condicao, CentoCusto, UnidadeMedida } from '$lib/types/CargaInicio';

// Constantes de opções de prioridade
export const prioridadeOptions = [
	{ value: 'alta', label: 'Alta' },
	{ value: 'media', label: 'Média' },
	{ value: 'baixa', label: 'Baixa' }
];

// Constantes de tipos de nota fiscal (NF)
export const tiposNFOptions = [
	{ value: 'revenda', label: 'Revenda' },
	{ value: 'despesa_imobilizado', label: 'Despesa/Imobilizado' },
	{ value: 'materia_prima', label: 'Matéria-Prima' },
	{ value: 'collection', label: 'Collection' }
];

// Você pode usar writable stores para gerenciar o estado, se necessário
export const selectedPriority = writable('');
export const selectedTipoNF = writable('');
export const condicoesStore = writable<Condicao[]>([]);
export const unidadeMedidaStore = writable<UnidadeMedida[]>([]);
export const centoCustoStore = writable<CentoCusto[]>([]);
export const filiaisStore = writable<any[]>([]);
export const xmlItemsStore = writable<ConexaoNFEItem[]>([]);
export const xmlDataStore = writable<ConexaoNFE | null>(null);
export const openComboboxId = writable<string | null>(null);
