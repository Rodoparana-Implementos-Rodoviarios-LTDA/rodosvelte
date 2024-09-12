import type { Column } from '$lib/types/tableTypes';
import ActionButton from '$lib/components/portaria/borracharia/ActionButton.svelte'; // Corrija o caminho para ActionButton

export const columns: Column<any>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: any) => row.Filial,
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: any) => row.NF,
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: any) => row.Cliente,
  },
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    cell: (row: any) => row.Vendedor,
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: any) => row.Produto,
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: any) => row.Saldo,
  },
  {
    accessorKey: 'Emissao',
    header: 'Emissão',
    cell: (row: any) => row.Emissao,
  },
  {
    header: 'Ações',
    accessorKey: 'actions',
    cell: (row: any) => ({
      $$component: ActionButton, // Usa o ActionButton diretamente na célula de Ações
      $$props: {
        saldoMaximo: row.Saldo, // Passa o saldo da linha para o componente
      }
    }),
  }
];
