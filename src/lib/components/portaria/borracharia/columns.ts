import type { borracharia, Column } from '$lib/types/tableTypes';
import ActionButton from './ActionButton.svelte';  // Importa o componente Action
import UserAvatar from '$lib/components/ui/UserAvatar.svelte'

export const columns: Column<borracharia>[] = [
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    component: UserAvatar,
    props: (row: borracharia) => {
      // Expressão regular para remover números e traços antes do nome
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
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: borracharia) => row.Produto,
    isFilterable: true
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: borracharia) => row.Saldo,
    isFilterable: true
  },
  {
    accessorKey: 'Emissao',
    header: 'Emissão',
    cell: (row: borracharia) => new Date(row.Emissao).toLocaleDateString('pt-BR'),
    isFilterable: true
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    component: ActionButton,
    props: (row: borracharia) => ({
      documentoCompleto: `${row.NF} - ${row.Produto}`,  // Exemplo de como juntar as informações para o ActionButton
      produtoCompleto: row.Produto,
      clienteCompleto: `${row.Cliente} - Nome do Cliente`,
      saldoMaximo: row.Saldo,
      filial: row.Filial
    }),
    isFilterable: false,
  },
];
export function customVendedorFilter(rows: borracharia[], filterValue: string) {
  return rows.filter((row) => {
    const vendedor = row.Vendedor;
    const regex = /[a-zA-Z]+/; // Apenas letras
    const match = vendedor.match(regex);
    if (match) {
      return match[0].toLowerCase().startsWith(filterValue.toLowerCase());
    }
    return false;
  });
}