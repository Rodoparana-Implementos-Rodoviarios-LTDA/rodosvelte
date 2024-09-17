import type { borracharia, Column } from '$lib/types/tableTypes';
import ActionButton from '$lib/components/portaria/borracharia/ActionButton.svelte';
import UserAvatar from '$lib/components/ui/UserAvatar.svelte' // Componente de Ações
  
export const columns: Column<borracharia>[] = [
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    component: UserAvatar, // Usa o componente UserAvatar para exibir o vendedor
    props: (row: borracharia) => ({ username: row.Vendedor }), // Passa o nome de usuário do vendedor para o componente
    isFilterable: true // Permitir filtro nessa coluna
  },
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
    cell: (row: borracharia) => new Date(row.Emissao).toLocaleDateString('pt-BR'), // Formatação direta de data
    isFilterable: true // Permitir filtro nessa coluna
  },
  {
    accessorKey: 'actions',
    header: 'Ações', // Cabeçalho de ações
    component: ActionButton, // Componente de ação
    props: (row: borracharia) => ({
      saldoMaximo: row.Saldo // Passa o saldo para o ActionButton
    }),
    isFilterable: false // Não é aplicável para filtro
  }
];
