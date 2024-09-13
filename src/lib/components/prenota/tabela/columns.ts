import type { Column, PreNota } from '$lib/types/tableTypes';
import UserAvatar from '$lib/components/ui/UserAvatar.svelte';
import StatusIcon from './StatusIcon.svelte';
import DateCell from './DateCell.svelte';
import ActionDropdown from './ActionDropdown.svelte';
import HoverText from '$lib/components/ui/HoverText.svelte';


export const columns: Column<PreNota>[] = [
  {
    accessorKey: 'Usuario',
    header: 'Usuário',
    component: UserAvatar,
    props: (row: PreNota) => ({ username: row.Usuario }),
    isFilterable: true
  },
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: PreNota) => row.Filial,
    isFilterable: true
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: PreNota) => row.NF,
    isFilterable: true
  },
  {
    accessorKey: 'Fornecedor',
    header: 'Fornecedor',
    cell: (row: PreNota) => row.Fornecedor,
    isFilterable: true
  },
  {
    accessorKey: 'Inclusao',
    header: 'Inclusão',
    component: DateCell,
    props: (row: PreNota) => ({
      rawDateInclusion: row.Inclusao,
      rawDateEmission: row.Emissao,
      rawDateDue: row.Vencimento
    }),
    isFilterable: true
  },
  {
    accessorKey: 'Valor',
    header: 'Valor',
    cell: (row: PreNota) => row.Valor,
    isFilterable: true
  },
  {
    accessorKey: 'Tipo',
    header: 'Tipo',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Tipo', value: row.Tipo }),
    isFilterable: true
  },
  {
    accessorKey: 'Status',
    header: 'Status',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Status', value: row.Status }),
    isFilterable: true
  },
  {
    accessorKey: 'Prioridade',
    header: 'Prioridade',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Prioridade', value: row.Prioridade }),
    isFilterable: true
  },
  {
    accessorKey: 'Obs',
    header: 'Observações',
    component: HoverText,
    props: (row: PreNota) => ({ text: row.Obs }),
    isFilterable: true
  },
  {
    accessorKey: 'actions',
    header: 'Histórico',
    component: ActionDropdown,
    props: (row: PreNota) => ({ rec: row.Rec }),
    isFilterable: false  // Ações lógicas, portanto, não é filtrável
  }
];
