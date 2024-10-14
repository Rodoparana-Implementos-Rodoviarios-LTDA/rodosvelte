// src/lib/stores/pedidoStore.ts
import { writable } from 'svelte/store';
import type { Pedido } from '$lib/types/Pedidos';

// Estrutura para armazenar os pedidos selecionados (usando um objeto para fácil acesso por número de pedido)
export const pedidosStore = writable<{ [numeroPedido: string]: Pedido }>({});

// Função para adicionar um pedido ao store
export const adicionarPedido = (novoPedido: Pedido) => {
	pedidosStore.update((pedidosAtuais) => {
		return {
			...pedidosAtuais,
			[novoPedido.pedido]: novoPedido // Adiciona ou atualiza o pedido com base no número do pedido
		};
	});
};

// Função para remover um pedido pelo número
export const removerPedido = (numeroPedido: string) => {
	pedidosStore.update((pedidosAtuais) => {
		// Usamos destructuring para remover o pedido específico, preservando o resto
		const { [numeroPedido]: _, ...restoPedidos } = pedidosAtuais; // Remove o pedido
		return restoPedidos;
	});
};

// Função para obter todos os pedidos
export const obterPedidos = () => {
	let currentPedidos;
	pedidosStore.subscribe((value) => (currentPedidos = value))(); // Assina a store para pegar o valor atual
	return currentPedidos;
};

// Função para limpar todos os pedidos
export const limparPedidos = () => {
	pedidosStore.set({}); // Reseta a store de pedidos
};

// Função para buscar um pedido específico pelo número
export const buscarPedidoPorNumero = (numeroPedido: string): Pedido | undefined => {
	let currentPedidos = obterPedidos();
	return currentPedidos[numeroPedido]; // Retorna o pedido se encontrado
};
