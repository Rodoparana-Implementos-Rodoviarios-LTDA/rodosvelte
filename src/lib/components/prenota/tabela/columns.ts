import type { Column, PreNota } from '$lib/types/tableTypes';
import UserAvatar from '$lib/components/ui/tabela/UserAvatar.svelte';
import StatusIcon from './StatusIcon.svelte';
import DateCell from './DateCell.svelte';
import ActionDropdown from './ActionDropdown.svelte';
import HoverText from '$lib/components/ui/tabela/HoverText.svelte';
import Filtrar from './Filtrar.svelte';

export const columns: Column<PreNota>[] = [
  {
    accessorKey: 'Usuario',
    header: 'Usuário',
    component: UserAvatar,
    props: (row: PreNota) => ({ username: row.Usuario }),
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: PreNota) => row.Filial,
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: PreNota) => row.NF,
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Fornecedor',
    header: 'Fornecedor',
    cell: (row: PreNota) => row.Fornecedor,
    isFilterable: true  // Indica que esta coluna pode ser filtrada
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
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Valor',
    header: 'Valor',
    cell: (row: PreNota) => row.Valor,
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Tipo',
    header: 'Tipo',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Tipo', value: row.Tipo }),
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Status',
    header: 'Status',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Status', value: row.Status }),
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Prioridade',
    header: 'Prioridade',
    component: StatusIcon,
    props: (row: PreNota) => ({ type: 'Prioridade', value: row.Prioridade }),
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'Obs',
    header: 'Observações',
    component: HoverText,
    props: (row: PreNota) => ({ text: row.Obs }),
    isFilterable: true  // Indica que esta coluna pode ser filtrada
  },
  {
    accessorKey: 'actions',
    header: Filtrar,  // Passa o componente Filtrar no lugar do cabeçalho
    component: ActionDropdown,
    props: (row: PreNota) => ({ rec: row.Rec }),
    isFilterable: false  // Não é filtrável, já que é um campo de ações
  }
];
