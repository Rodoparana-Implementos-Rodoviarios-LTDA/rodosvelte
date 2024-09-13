import type { borracharia, Column } from '$lib/types/tableTypes';
import ActionButton from '$lib/components/portaria/borracharia/ActionButton.svelte'; // Componente de Ações
import Filtrar from '$lib/components/prenota/tabela/Filtrar.svelte'; // Componente de Filtro

export const columns: Column<borracharia>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: borracharia) => row.Filial,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: borracharia) => row.NF,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: borracharia) => row.Cliente,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    cell: (row: borracharia) => row.Vendedor,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: borracharia) => row.Produto,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: borracharia) => row.Saldo,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'Emissao',
    header: 'Emissão',
    cell: (row: borracharia) => row.Emissao,
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    header: Filtrar, // Componente de Filtro
    accessorKey: 'actions',
    component: ActionButton, // Componente que renderiza as ações
    props: (row: borracharia) => ({
      saldoMaximo: row.Saldo // Passa o saldo para o componente ActionButton
    }),
    isFilterable: false // Não aplicável para filtro
  }
];
