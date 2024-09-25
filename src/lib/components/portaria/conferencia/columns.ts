import type { HistoricoData, Column } from '$lib/types/tableTypes';
import Action from './Action.svelte';  // Corrige o caminho para o Action.svelte

export const columns: Column<HistoricoData>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: HistoricoData) => row.Filial,
    isFilterable: true
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: HistoricoData) => row.NF,
    isFilterable: true
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: HistoricoData) => row.Cliente,
    isFilterable: true
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: HistoricoData) => row.Produto,
    isFilterable: true
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
    isFilterable: true
  },
  {
    accessorKey: 'Placa',
    header: 'Placa',
    cell: (row: HistoricoData) => row.Placa,
    isFilterable: true
  },
  {
    accessorKey: 'Observacao',  // Coluna para observações
    header: 'Observação',
    cell: (row: HistoricoData) => row.Observacao || 'Sem Observação',
    isFilterable: true // Permitir filtro na coluna de observação
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: HistoricoData) => row.Saldo.toFixed(2),
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    component: Action,  // Usa o componente Action para ações
    props: (row: HistoricoData) => ({
      documento: row.NF,  // Passa o NF (nota fiscal)
      produto: row.Produto,  // Passa o Produto
      saldoConferencia: row.Saldo,  // Passa o saldo máximo
      responsavel: row.Responsavel
    }),
    isFilterable: false,  // Ações não são filtráveis
  },
];
