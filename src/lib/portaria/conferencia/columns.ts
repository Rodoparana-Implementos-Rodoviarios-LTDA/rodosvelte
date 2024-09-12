import type { Column } from '$lib/types/tableTypes';
import type { Portaria } from '$lib/types/tableTypes';

// Definição das colunas para a tabela de Portaria
export const columns: Column<Portaria>[] = [
  {
    header: 'Filial',
    accessorKey: 'Filial',
  },
  {
    header: 'Nota Fiscal',
    accessorKey: 'NF',
  },
  {
    header: 'Cliente',
    accessorKey: 'Cliente',
    cell: (row: Portaria) => {
      const [codCliente, filialCliente] = row.Cliente.split(' - ', 2);
      return `<strong>${codCliente}</strong> - ${filialCliente}`;
    },
  },
  {
    header: 'Produto',
    accessorKey: 'Produto',
    cell: (row: Portaria) => {
      const [codProduto, filialProduto, descricaoProduto] = row.Produto.split(' - ', 3);
      return `<strong>${codProduto} ${filialProduto}</strong><br>${descricaoProduto}`;
    },
  },
  {
    header: 'Tipo Movimentação',
    accessorKey: 'TipoMov',
    cell: (row: Portaria) => (row.TipoMov === 'C' ? 'Compra' : 'Venda'), // Exibe 'Compra' ou 'Venda' baseado no tipo
  },
  {
    header: 'Data e Hora',
    accessorKey: 'DataHora',
    cell: (row: Portaria) => {
      const dataHora = new Date(row.DataHora);
      return dataHora.toLocaleString('pt-BR'); // Formato completo de data e hora
    },
  },
  {
    header: 'Responsável',
    accessorKey: 'Responsavel',
  },
  {
    header: 'Placa',
    accessorKey: 'Placa',
  },
  {
    header: 'Observação',
    accessorKey: 'Observacao',
    cell: (row: Portaria) => row.Observacao || 'Sem observações',
  },
  {
    header: 'Saldo',
    accessorKey: 'Saldo',
    cell: (row: Portaria) => `<strong>${row.Saldo}</strong>`,
  },
];
