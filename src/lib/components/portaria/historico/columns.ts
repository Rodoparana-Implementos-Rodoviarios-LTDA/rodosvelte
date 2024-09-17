import type { HistoricoData, Column } from '$lib/types/tableTypes';
import ConfButton from '$lib/components/portaria/historico/ConfButton.svelte'

export const columns: Column<HistoricoData>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: HistoricoData) => row.Filial,
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: HistoricoData) => row.NF,
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: HistoricoData) => row.Cliente,
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: HistoricoData) => row.Produto,
  },
  {
    accessorKey: 'DataHora',
    header: 'Data/Hora',
    cell: (row: HistoricoData) => row.DataHora,
  },
  {
    accessorKey: 'Responsavel',
    header: 'Responsável',
    cell: (row: HistoricoData) => row.Responsavel,
  },
  {
    accessorKey: 'Placa',
    header: 'Placa',
    cell: (row: HistoricoData) => row.Placa,
  },
  {
    accessorKey: 'Observacao',
    header: 'Observação',
    cell: (row: HistoricoData) => row.Observacao || 'Sem Observação',
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: HistoricoData) => row.Saldo.toFixed(2),
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    component: ConfButton, // Usa o componente ConfButton
    props: (row: HistoricoData) => ({
      documento: row.NF,  // Passa o NF (nota fiscal) para o componente ConfButton
      produto: row.Produto,  // Passa o Produto para o componente ConfButton
      saldoMaximo: row.Saldo,  // Passa o saldo para o componente ConfButton
      responsavel: row.Responsavel, // Passa o responsável
    }),
    isFilterable: false, // Ações não são filtráveis
  },
];