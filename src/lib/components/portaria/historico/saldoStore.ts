// src/lib/stores/saldoStore.ts
import { writable } from 'svelte/store';

// Cria uma store para armazenar o saldo da borracharia
export const saldoStore = writable<number>(0);
