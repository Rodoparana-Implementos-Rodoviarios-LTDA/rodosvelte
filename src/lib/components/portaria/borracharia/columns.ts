import type { borracharia, Column } from '$lib/types/tableTypes';
import ActionButton from '$lib/components/portaria/borracharia/ActionButton.svelte'; // Corrija o caminho para ActionButton

export const columns: Column<borracharia>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: borracharia) => row.Filial,
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: borracharia) => row.NF,
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: borracharia) => row.Cliente,
  },
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    cell: (row: borracharia) => row.Vendedor,
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: borracharia) => row.Produto,
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: borracharia) => row.Saldo,
  },
  {
    accessorKey: 'Emissao',
    header: 'Emissão',
    cell: (row: borracharia) => row.Emissao,
  },
  {
    header: 'Ações',
    accessorKey: 'actions',
    component: ActionButton, // Componente que renderiza o dropdown de ações
    props: (row: borracharia) => ({ saldo: row.Saldo }) // Passa o saldo para o ActionButton
  }
];
