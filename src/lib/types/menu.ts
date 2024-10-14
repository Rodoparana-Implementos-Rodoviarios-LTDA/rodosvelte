import type { ComponentType } from 'svelte';
export interface MenuItem {
	name: string;
	iconUrl?: string;
	link: string;
	external?: boolean;
}

export interface MenuSection {
	title: string;
	icon: ComponentType;
	link?: string; // Alguns itens podem ter apenas um link, como "Início"
	items?: MenuItem[]; // Deixa `items` opcional ou vazio para seções como "Início"
}
