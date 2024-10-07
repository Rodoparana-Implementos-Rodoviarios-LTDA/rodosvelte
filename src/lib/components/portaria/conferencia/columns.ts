
import type { HistoricoData, Column } from '$lib/types/tableTypes';
import ActionButton from './ActionButton.svelte';  // Componente para a ação de conferência

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
    isFilterable: true
  },
  {
    accessorKey: 'Responsavel',
    header: 'Responsável',
    cell: (row: HistoricoData) => row.Responsavel,
    isFilterable: true
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    component: ActionButton,  // Usa o componente ActionButton para ações
    props: (row: HistoricoData) => ({
      documento: row.NF,            // Passa o NF (nota fiscal)
      produto: row.Produto,         // Passa o Produto
      saldoConferencia: row.Saldo,  // Passa o saldo
      responsavel: row.Responsavel,
      filial: row.Filial,           // Adicionado para uso no componente ActionButton
      cliente: row.Cliente          // Adicionado para uso no componente ActionButton
    }),
<<<<<<< Updated upstream
    isFilterable: false,  // Ações não são filtráveis
  },
=======
    isFilterable: false  // Ações não são filtráveis
  }
];
>>>>>>> Stashed changes
