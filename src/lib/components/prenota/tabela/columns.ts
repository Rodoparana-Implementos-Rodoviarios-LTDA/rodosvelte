// src/lib/components/prenota/tabela/columns.ts
import type { Column, PreNota } from '$lib/types/tableTypes';
import UserAvatar from '$lib/components/ui/tabela/UserAvatar.svelte';
import StatusIcon from './StatusIcon.svelte';
import DateCell from './DateCell.svelte';
import HoverText from '$lib/components/ui/tabela/HoverText.svelte';
export const columns: Column<PreNota>[] = [
  {
    accessorKey: 'Usuario',  // Usa o nome exato da chave no objeto PreNota
    header: 'Usuário',
    component: UserAvatar,
    props: (row: PreNota) => ({ username: row.Usuario })
  },
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: PreNota) => row.Filial
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: PreNota) => row.NF
  },
  {
    accessorKey: 'Fornecedor',
    header: 'Fornecedor',
    cell: (row: PreNota) => row.Fornecedor
  },
  {
    accessorKey: 'Vencimento',
    header: 'Vencimento',
    component: DateCell,
    props: (row: PreNota) => ({
      rawDateInclusion: row.Inclusao,
      rawDateEmission: row.Emissao,
      rawDateDue: row.Vencimento
    })
  },
  {
    accessorKey: 'Valor',
    header: 'Valor',
    cell: (row: PreNota) => row.Valor
  },
  {
    accessorKey: 'Tipo',
    header: 'Tipo',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Tipo', value: row.Tipo })
  },
  {
    accessorKey: 'Status',
    header: 'Status',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Status', value: row.Status })
  },
  {
    accessorKey: 'Prioridade',
    header: 'Prioridade',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Prioridade', value: row.Prioridade })
  },
  {
    accessorKey: 'Obs',
    header: 'Observações',
    component: HoverText,  // Usando HoverText para exibir observação completa ao passar o mouse
    props: (row: PreNota) => ({ text: row.Obs })
  }
];
