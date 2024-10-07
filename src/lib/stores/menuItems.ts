import type { MenuSection } from '$lib/types/menu';
import {
	IconList,
	IconPlus,
	IconUpload,
	IconDashboard,
	IconClockCheck,
	IconShieldCheck,
	IconBulldozer,
	IconTruck,
	IconPhone,
	IconInfoCircle
} from '@tabler/icons-svelte';

export const menuItems: MenuSection[] = [
	{
		title: 'Lançamento de Notas',
		items: [
			{ name: 'Lista de Pre Notas', icon: IconList, link: '/lancamento-notas/' },
			{ name: 'Incluir Manualmente', icon: IconPlus, link: '/lancamento-notas/incluir' },
			{ name: 'Incluir XML', icon: IconUpload, link: '/lancamento-notas/xml' }
		]
	},
	{
		title: 'Controle de Itens',
		items: [
			{ name: 'Borracharia', icon: IconDashboard, link: '/controle-itens' },
			{ name: 'Conferência', icon: IconClockCheck, link: '/controle-itens/conferencia' },
			{ name: 'Conferidos', icon: IconShieldCheck, link: '/controle-itens/conferidos' }
		]
	},
	{
		title: 'Assinatura de Email',
		items: [
			{
				name: 'Timber',
				icon: IconBulldozer,
				link: '/assinatura/timber'
			},
			{
				name: 'Rodoparaná',
				icon: IconTruck,
				link: '/assinatura/rodo'
			}
		]
	},
	{
		title: 'Suporte e Intranet',
		items: [
			{
				name: 'Suporte',
				icon: IconPhone,
				link: 'https://hesk.rodoparana.com.br',
				external: true
			},
			{
				name: 'Intranet',
				icon: IconInfoCircle,
				link: 'https://sites.google.com/site/baserodoparana/home?authuser=1',
				external: true
			}
		]
	}
];
