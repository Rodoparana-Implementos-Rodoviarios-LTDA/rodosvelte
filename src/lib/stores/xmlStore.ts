// src/lib/stores/xmlStore.ts
import { writable } from 'svelte/store';
import type { ConexaoNFE, ConexaoNFEItem } from '$lib/types/ConexaoNFE';
import type { Condicao, CentoCusto, UnidadeMedida } from '$lib/types/CargaInicio';
import type { ProdutoOption } from '$lib/types/Produtos';

// Constantes de opções de prioridade
export const prioridadeOptions = [
	{ value: 'alta', label: 'Alta' },
	{ value: 'media', label: 'Média' },
	{ value: 'baixa', label: 'Baixa' }
];

// Constantes de tipos de nota fiscal (NF)
export const tiposNFOptions = [
	{ value: 'Revenda', label: 'Revenda' },
	{ value: 'Despesa/Imobilizado', label: 'Despesa/Imobilizado' },
	{ value: 'Matéria Prima', label: 'Matéria Prima' },
	{ value: 'Collection', label: 'Collection' }
];

// Writable stores para gerenciar estados selecionados
export const selectedPriority = writable<string>(''); // Prioridade selecionada
export const selectedTipoNF = writable<string>(''); // Tipo de NF selecionado
export const selectedCondicao = writable<string>(''); // Condição de pagamento selecionada
export const condicoesStore = writable<Condicao[]>([]); // Condições de pagamento
export const unidadeMedidaStore = writable<UnidadeMedida[]>([]); // Unidade de medida
export const centoCustoStore = writable<CentoCusto[]>([]); // Centro de custo
export const filiaisStore = writable<{ value: string; label: string }[]>([]); // Filiais

// Writable stores para os dados da XML
export const xmlItemsStore = writable<ConexaoNFEItem[]>([]); // Itens da XML
export const xmlDataStore = writable<ConexaoNFE | null>(null); // Dados gerais da XML

export const fetchState = writable<'inicial' | 'carregando' | 'sucesso' | 'erro'>('inicial');

// Writable store para produtos
export const produtosStore = writable<ProdutoOption[]>([]); // Armazenar os produtos após o fetch
