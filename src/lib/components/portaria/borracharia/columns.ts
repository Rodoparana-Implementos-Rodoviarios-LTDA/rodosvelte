// src/lib/components/portaria/borracharia/columns.ts
import type { borracharia, Column } from '$lib/types/tableTypes';
import ActionButton from './ActionButton.svelte';  // Importa o componente Action
import UserAvatar from '$lib/components/ui/UserAvatar.svelte';

export const columns: Column<borracharia>[] = [
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    component: UserAvatar,
    props: (row: borracharia) => {
      const vendedor = row.Vendedor;
      const regex = /^\d+\s*-\s*(.*)$/;  // Captura apenas o nome após o número e traço
      const match = vendedor.match(regex);
      let formattedVendedor = vendedor; // Caso não haja match, mantém o valor original

      if (match) {
        formattedVendedor = match[1]; // Usa apenas o nome
      }

      return {
        username: formattedVendedor // Passa o nome formatado para o componente Avatar
      };
    },
    isFilterable: true,
  },
  
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: borracharia) => row.Filial,
    isFilterable: true
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: borracharia) => row.NF,
    isFilterable: true
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: borracharia) => row.Cliente,
    isFilterable: true
  },
  {
    accessorKey: 'Emissao',
    header: 'Emissão',
    cell: (row: borracharia) => row.Emissao,
    isFilterable: true
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    component: ActionButton,
    props: (row: borracharia) => ({
      documentoCompleto: row.NF,
      clienteCompleto: row.Cliente,
      filial: row.Filial
    }),
    isFilterable: false,
  },

];
