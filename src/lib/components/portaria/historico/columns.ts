import type { Column } from '$lib/types/tableTypes';

export const columnsHistorico: Column<any>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: (row: any) => row.Filial || 'N/A', // Exibe 'N/A' se não houver valor
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: (row: any) => row.NF || 'N/A', // Exibe 'N/A' se não houver valor
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: (row: any) => row.Cliente || 'N/A',
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: (row: any) => row.Produto || 'N/A',
  },
  {
    accessorKey: 'TipoMov',
    header: 'Tipo de Movimento',
    cell: (row: any) => row.TipoMov === 'C' ? 'Crédito' : 'Débito', // Exemplo de conversão para crédito/débito
  },
  {
    accessorKey: 'DataHora',
    header: 'Data/Hora',
    cell: (row: any) => new Date(row.DataHora).toLocaleString('pt-BR'), // Formata data/hora
  },
  {
    accessorKey: 'Responsavel',
    header: 'Responsável',
    cell: (row: any) => row.Responsavel || 'N/A',
  },
  {
    accessorKey: 'Placa',
    header: 'Placa',
    cell: (row: any) => row.Placa || 'N/A',
  },
  {
    accessorKey: 'Observacao',
    header: 'Observação',
    cell: (row: any) => row.Observacao || 'Sem observações',
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: (row: any) => row.Saldo.toFixed(2), // Formata saldo com 2 casas decimais
  },
];
