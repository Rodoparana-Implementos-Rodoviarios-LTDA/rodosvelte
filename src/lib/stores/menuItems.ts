// src/lib/stores/menuItems.ts
import type { MenuSection } from '$lib/types/menu';
import {
	IconAlignBoxLeftBottomFilled,
	IconHelpCircleFilled,
	IconHome,
	IconWheel,
	IconWritingSign
} from '@tabler/icons-svelte';

export const menuItems: MenuSection[] = [
	{
		title: 'Início', // Nome da seção
		icon: IconHome, // Ícone de início
		link: '/inicio', // Link direto para "Início"
		items: [] // Deixa o array vazio, indicando que não tem sub-itens
	},
	{
		title: 'Lançamento de Notas',
		icon: IconAlignBoxLeftBottomFilled,
		items: [
			{
				name: 'Lista de Pre Notas',
				iconUrl: 'https://cdn.lordicon.com/xpgofwru.json',
				link: '/lancamento-notas/'
			},
			{
				name: 'Incluir Manualmente',
				iconUrl: 'https://cdn.lordicon.com/hrjifpbq.json',
				link: '/lancamento-notas/incluir'
			},
			{
				name: 'Incluir XML',
				iconUrl: 'https://cdn.lordicon.com/fowixcuo.json',
				link: '/lancamento-notas/xml'
			}
		]
	},
	{
		title: 'Controle de Itens',
		icon: IconWheel,
		items: [
			{
				name: 'Borracharia',
				iconUrl: 'https://cdn.lordicon.com/ifsxxxte.json',
				link: '/controle-itens'
			},
			{
				name: 'Conferência',
				iconUrl: 'https://cdn.lordicon.com/qnpnzlkk.json',
				link: '/controle-itens/conferencia'
			},
			{
				name: 'Histórico de Conferência',
				iconUrl: 'https://cdn.lordicon.com/dafdkyyt.json',
				link: '/controle-itens/conferidos'
			}
		]
	},
	{
		title: 'Assinatura de Email',
		icon: IconWritingSign,
		items: [
			{
				name: 'Timber',
				iconUrl: 'https://cdn.lordicon.com/boyoxams.json',
				link: '/assinatura/timber'
			},
			{
				name: 'Rodoparaná',
				iconUrl: 'https://cdn.lordicon.com/lomfljuq.json',
				link: '/assinatura/rodo'
			}
		]
	},
	{
		title: 'Suporte e Intranet',
		icon: IconHelpCircleFilled,
		items: [
			{
				name: 'Suporte',
				iconUrl: 'https://cdn.lordicon.com/iazmohzf.json',
				link: 'https://hesk.rodoparana.com.br',
				external: true
			},
			{
				name: 'Intranet',
				iconUrl: 'https://cdn.lordicon.com/rwtswsap.json',
				link: 'https://sites.google.com/site/baserodoparana/home?authuser=1',
				external: true
			}
		]
	}
];
