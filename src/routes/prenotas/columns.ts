import type { PreNota } from '$lib/types/tableTypes';
import HoverCard from '$lib/components/ui/HoverCard/HoverCard.svelte';
import UserAvatar from '$lib/components/ui/table/UserAvatar.svelte';
import Badge from '$lib/components/prenota/tabela/Badge.svelte';
import Fornecedor from '$lib/components/prenota/tabela/Fornecedor.svelte';
import ActionDropdown from '$lib/components/prenota/tabela/ActionDropdown.svelte';
import Dates from '$lib/components/prenota/tabela/Dates.svelte';

// Definindo as colunas para a tabela de PreNota
export const columns = [
  {
    accessorKey: 'Usuario',
    header: 'Usuário',
    cell: ({ row }: { row: PreNota }) => ({
      component: UserAvatar,
      props: {
        name: row.Usuario,
        imageUrl: row.avatarUrl ?? null,
      },
    }),
  },
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: ({ row }: { row: PreNota }) => row.Filial,
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: ({ row }: { row: PreNota }) => row.NF,
  },
  {
    accessorKey: 'Fornecedor',
    header: 'Fornecedor',
    cell: ({ row }: { row: PreNota }) => ({
      component: Fornecedor,
      props: {
        fornecedorInfo: row.Fornecedor,
        hoverCardId: 'hovercard-' + Math.random().toString(36).substr(2, 9),
      },
    }),
  },
  {
    accessorKey: 'Datas',
    header: 'Datas',
    cell: ({ row }: { row: PreNota }) => ({
      component: Dates,
      props: {
        inclusao: row.Inclusao,
        emissao: row.Emissao,
        vencimento: row.Vencimento,
      },
    }),
  },
  {
    accessorKey: 'Valor',
    header: 'Valor',
    cell: ({ row }: { row: PreNota }) => row.Valor,
  },
  {
    accessorKey: 'Tipo',
    header: 'Tipo',
    cell: ({ row }: { row: PreNota }) => ({
      component: Badge,
      props: {
        label: row.Tipo,
        category: 'tipo',
      },
    }),
  },
  {
    accessorKey: 'Status',
    header: 'Status',
    cell: ({ row }: { row: PreNota }) => ({
      component: Badge,
      props: {
        label: row.Status,
        category: 'status',
      },
    }),
  },
  {
    accessorKey: 'Prioridade',
    header: 'Prioridade',
    cell: ({ row }: { row: PreNota }) => ({
      component: Badge,
      props: {
        label: row.Prioridade,
        category: 'prioridade',
      },
    }),
  },
  {
    accessorKey: 'Obs',
    header: 'Observações',
    cell: ({ row }: { row: PreNota }) => ({
      component: HoverCard,
      props: {
        content: row.Obs,
        truncatedContent: row.Obs.slice(0, 20) + '...',
        hoverCardId: 'hovercard-obs-' + Math.random().toString(36).substr(2, 9),
      },
    }),
  },
  {
    accessorKey: 'actions',
    header: 'Ações',
    cell: ({ row }: { row: PreNota }) => ({
      component: ActionDropdown,
      props: { rec: row.Rec },
    }),
  },
];
