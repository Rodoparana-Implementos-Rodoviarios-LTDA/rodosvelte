import { writable } from 'svelte/store';

// Estado inicial dos filtros
const initialFilters = {};

// Store de filtros
export const filterStore = writable<Record<string, string>>(initialFilters);

// Função para aplicar um filtro
export function setFilter(key: string, value: string) {
	filterStore.update((filters) => {
		if (value) {
			filters[key] = value;
		} else {
			delete filters[key]; // Remove o filtro se o valor for vazio
		}
		return filters;
	});
}

// Função para resetar os filtros
export function resetFilters() {
	filterStore.set({});
}
