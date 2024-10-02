import type { Column, XML } from '$lib/types';
import Impostos from './Impostos.svelte';
import CurrencyFormatter from '$lib/components/ui/tabela/CurrencyFormatter.svelte'; // Importe o componente de formatação de moeda

// Defina as colunas da tabela de itens da nota fiscal
export const columns: Column<XML>[] = [
	{
		accessorKey: 'codProduto',
		header: 'Código',
		class: 'text-center'
	},
	{
		accessorKey: 'descProduto',
		header: 'Descrição Produto',
		class: 'text-left truncate'
	},
	{
		accessorKey: 'ncmsh',
		header: 'NCM/SH',
		class: 'text-center'
	},
	{
		accessorKey: 'cst',
		header: 'CST',
		class: 'text-center'
	},
	{
		accessorKey: 'origem',
		header: 'Origem',
		class: 'text-center'
	},
	{
		accessorKey: 'cfop',
		header: 'CFOP',
		class: 'text-center'
	},
	{
		accessorKey: 'unidade',
		header: 'Unidade',
		class: 'text-center'
	},
	{
		accessorKey: 'quantidade',
		header: 'Quantidade',
		class: 'text-center'
	},
	{
		accessorKey: 'valorUnitario',
		header: 'Valor Unitário',
		component: CurrencyFormatter, // Usa o componente para formatar o valor
		props: (row: any) => ({
			value: parseFloat(row.valorUnitario)
		}),
		class: 'text-right'
	},
	{
		accessorKey: 'valorTotal',
		header: 'Valor Total',
		component: CurrencyFormatter, // Usa o componente para formatar o valor
		props: (row: any) => ({
			value: parseFloat(row.valorTotal)
		}),
		class: 'text-right'
	},
	{
		accessorKey: 'valorIcms',
		header: 'Impostos',
		component: Impostos,
		props: (row: any) => ({
			valorIcms: row.valorIcms,
			aliqIcms: row.aliqIcms,
			valorIpi: row.valorIpi,
			aliqIpi: row.aliqIpi
		}),
		class: 'w-36 text-center'
	}
];
