// src/lib/stores/pedidoStore.ts
import { writable } from 'svelte/store';
import type { Pedido } from '$lib/types/Pedidos';

// Estrutura para armazenar os pedidos selecionados
export const pedidosStore = writable<{ [numeroPedido: string]: Pedido }>({});

// Função para adicionar um pedido ao store
export const adicionarPedido = (novoPedido: Pedido) => {
	pedidosStore.update((pedidosAtuais) => {
		return {
			...pedidosAtuais,
			[novoPedido.pedido]: novoPedido // Adiciona ou atualiza o pedido
		};
	});
};

// Função para remover um pedido pelo número
export const removerPedido = (numeroPedido: string) => {
	pedidosStore.update((pedidosAtuais) => {
		const { [numeroPedido]: _, ...restoPedidos } = pedidosAtuais; // Remove o pedido
		return restoPedidos;
	});
};

// Função para obter todos os pedidos
export const obterPedidos = () => {
	let currentPedidos;
	pedidosStore.subscribe((value) => (currentPedidos = value))();
	return currentPedidos;
};

// Limpar todos os pedidos
export const limparPedidos = () => {
	pedidosStore.set({});
};
