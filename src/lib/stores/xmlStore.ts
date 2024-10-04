// src/lib/stores/xmlStore.ts
import { writable } from 'svelte/store';
import type { DetalhesXML } from '$lib/types';

export const xmlKeyStore = writable<string>('');
export const xmlDataStore = writable<DetalhesXML | null>(null);
