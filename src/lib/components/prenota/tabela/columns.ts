import type { Column, PreNota } from '$lib/types/tableTypes';
import UserAvatar from '$lib/components/ui/UserAvatar.svelte';

import fililalHover from '$lib/components/ui/fililalHover.svelte';
import StatusIcon from './StatusIcon.svelte';
import DateCell from './DateCell.svelte';
import TipoBadge from './TipoBadge.svelte';
import ActionDropdown from '../revisao/ActionDropdown.svelte';
import CurrencyFormatter from './CurrencyFormatter.svelte';

export const columns: Column<PreNota>[] = [
	{
		accessorKey: 'Status',
		header: 'Status',
		component: StatusIcon,
		props: (row: PreNota) => ({ type: 'Status', value: row.Status }),
		isFilterable: true,
		class: 'w-10 text-center'
	},
	{
		accessorKey: 'Usuario',
		header: 'Usuário',
		component: UserAvatar, // Componente que exibe o avatar do usuário
		props: (row: PreNota) => ({ username: row.Usuario }), // Passa o nome de usuário para o componente
		isFilterable: true, // Coluna filtrável
		class: 'w-10 text-center'
	},
	{
		accessorKey: 'Filial',
		header: 'Filial',
		component: fililalHover,
		props: (row: PreNota) => ({ numeroFilial: row.Filial }),
		isFilterable: true
	},
	{
		accessorKey: 'NF',
		header: 'Nota Fiscal',
		cell: (row: PreNota) => row.NF,
		isFilterable: true,
		class: 'w-36 text-center'
	},
	{
		accessorKey: 'Fornecedor',
		header: 'Fornecedor',
		props: (row: PreNota) => ({ text: row.Fornecedor }),
		isFilterable: true,
		class: 'max-w-sm text-center'
	},
	{
		accessorKey: 'Tipo',
		header: 'Tipo',
		component: TipoBadge,
		props: (row: PreNota) => ({ tipo: row.Tipo }),
		isFilterable: true,
		class: 'w-56 text-center'
	},
	{
		accessorKey: 'Inclusao',
		header: 'Inclusão',
		component: DateCell,
		props: (row: PreNota) => ({
			rawDateInclusion: row.Inclusao,
			rawDateEmission: row.Emissao,
			rawDateDue: row.Vencimento
		}),
		isFilterable: true,
		class: 'w-36 text-center'
	},
	{
		accessorKey: 'Valor',
		header: 'Valor',
		component: CurrencyFormatter,
		props: (row: PreNota) => ({ value: parseFloat(row.Valor) }), // Passa o valor como número para o componente
		isFilterable: true,
		class: 'w-36 text-center'
	},
	{
		accessorKey: 'Prioridade',
		header: 'Prioridade',
		component: StatusIcon,
		props: (row: PreNota) => ({ type: 'Prioridade', value: row.Prioridade }),
		isFilterable: true,
		class: 'w-36 text-center'
	},
	{
		accessorKey: 'Obs',
		header: 'Observações',
		props: (row: PreNota) => ({ text: row.Obs }),
		isFilterable: true,
		class: 'max-w-xs text-center'
	},
	{
		accessorKey: 'Rec',
		header: 'Rec',
		props: (row: PreNota) => ({ text: row.Rec }),
		isFilterable: true,
		class: 'max-w-xs text-center'
	},
	{
		accessorKey: 'actions',
		header: 'Sobre',
		component: ActionDropdown,
		props: (row: PreNota) => ({ rec: row.Rec, status: row.Status }), // Passa tanto o Rec quanto o Status
		isFilterable: false,
		class: 'w-36 text-center'
	}
];
