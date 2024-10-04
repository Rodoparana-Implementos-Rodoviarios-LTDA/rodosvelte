// columns.ts
import type { Column } from '$lib/types/tableTypes';
import type { HistoricoConferido } from '$lib/types/tableTypes';


export const columns: Column<HistoricoConferido>[] = [

	{
		accessorKey: 'Filial',
		header: 'Filial',
		cell: (row: HistoricoConferido) => row.Filial,
		isFilterable: true,
	},
	{
		accessorKey: 'NF',
		header: 'Nota Fiscal',
		cell: (row: HistoricoConferido) => row.NF,
		isFilterable: true,
	},
	{
		accessorKey: 'Cliente',
		header: 'Cliente',
		cell: (row: HistoricoConferido) => row.Cliente,
		isFilterable: true,
	},
	{
		accessorKey: 'Produto',
		header: 'Produto',
		cell: (row: HistoricoConferido) => row.Produto,
		isFilterable: true,
	},
	{
		accessorKey: 'DataHora',
		header: 'Data/Hora',
		cell: (row: HistoricoConferido) => row.DataHora,
		isFilterable: true,
	},
	{
		accessorKey: 'Responsavel',
		header: 'Responsável',
		cell: (row: HistoricoConferido) => row.Responsavel,
		isFilterable: true,
	},
	{
		accessorKey: 'Placa',
		header: 'Placa',
		cell: (row: HistoricoConferido) => row.Placa,
		isFilterable: true,
	},
	{
		accessorKey: 'Observacao',
		header: 'Observação',
		cell: (row: HistoricoConferido) => row.Observacao || 'Sem Observação',
		isFilterable: true,
	},
	{
		accessorKey: 'DataConf',
		header: 'Data de Confirmação',
		cell: (row: HistoricoConferido) => row.DataConf,
		isFilterable: true,
	},
	{
		accessorKey: 'Saldo',
		header: 'Saldo',
		cell: (row: HistoricoConferido) => row.Saldo.toString(),
		isFilterable: true,
	},
];
