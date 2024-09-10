// src/lib/types/tableTypes.ts
// src/lib/types/tableTypes.ts
// src/lib/types/tableTypes.ts
export interface Column<T> {
  header: string;                     // O rótulo da coluna (cabeçalho)
  accessorKey: keyof T;               // Chave usada para acessar o dado no objeto T
  cell?: (row: T) => any;             // Função opcional para customizar a célula
  component?: any;                    // Componente opcional para renderização de células dinâmicas
  props?: (row: T) => Record<string, any>; // Propriedades opcionais passadas para o componente
}


  export interface TableData {
    [key: string]: any; // Dados da tabela com chave genérica
  }
  export interface PreNota {
    Filial: string;
    NF: string;
    Status: string;
    Fornecedor: string;
    Emissao: string;
    Inclusao: string;
    Vencimento: string;
    Valor: string;
    Tipo: string;
    Prioridade: string;
    Usuario: string;
    Obs: string;

  }
  