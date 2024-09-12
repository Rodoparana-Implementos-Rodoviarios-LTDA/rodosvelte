import type { Column } from '$lib/types/tableTypes';

export const columnsHistorico: Column<any>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: any) => row.Filial
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: any) => row.NF
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: any) => row.Cliente
  },
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    cell: (row: any) => row.Vendedor
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: any) => row.Produto
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: any) => row.Saldo
  },
  {
    accessorKey: 'DataHora',
    header: 'Data/Hora',
    cell: (row: any) => row.DataHora
  },
  {
    accessorKey: 'Responsavel',
    header: 'Responsável',
    cell: (row: any) => row.Responsavel
  },
  {
    accessorKey: 'Placa',
    header: 'Placa',
    cell: (row: any) => row.Placa
  },
  {
    accessorKey: 'Observacao',
    header: 'Observação',
    cell: (row: any) => row.Observacao
  },
];
