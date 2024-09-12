import type { ColumnDef } from '@tanstack/svelte-table';
import type { Borracharia } from '$lib/types/tableTypes';


export const columns: ColumnDef<Borracharia>[] = [
  {
    accessorKey: 'Filial',
    header: 'Filial',
    cell: info => info.getValue(),
  },
  {
    accessorKey: 'NF',
    header: 'Nota Fiscal',
    cell: info => info.getValue(),
  },
  {
    accessorKey: 'Vendedor',
    header: 'Vendedor',
    cell: ({ row }) => {
      const [codVendedor, nomeVendedor] = row.original.Vendedor.split(' ', 2);
      return `<strong>${codVendedor}</strong> ${nomeVendedor}`;
    }
  },
  {
    accessorKey: 'Cliente',
    header: 'Cliente',
    cell: ({ row }) => {
      const [codCliente, filialCliente, nomeCliente] = row.original.Cliente.split(' ', 3);
      return `
        <strong>${codCliente} ${filialCliente}</strong><br>
        ${nomeCliente}
      `;
    }
  },
  {
    accessorKey: 'Produto',
    header: 'Produto',
    cell: ({ row }) => {
      const [codProduto, filialProduto, descricaoProduto] = row.original.Produto.split(' ', 3);
      return `
        <strong>${codProduto} ${filialProduto}</strong><br>
        ${descricaoProduto}
      `;
    }
  },
  {
    accessorKey: 'Emissao',
    header: 'Data de Emissão',
    cell: ({ row }) => {
      const dataEmissao = new Date(row.original.Emissao);
      return dataEmissao.toLocaleDateString('pt-BR'); // Formata a data para DD/MM/YYYY
    }
  },
  {
    accessorKey: 'Saldo',
    header: 'Saldo',
    cell: info => `<strong>${info.getValue()}</strong>`,
  }
];
