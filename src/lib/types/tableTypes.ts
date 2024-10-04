// src/lib/types/tableTypes.ts

// Coluna genérica usada em várias tabelas
// src/lib/types/tableTypes.ts

// Coluna genérica usada em várias tabelas
// src/lib/types/tableTypes.ts

// Coluna genérica usada em várias tabelas
export interface Column<T> {
	header: string; // O rótulo da coluna (cabeçalho)
	accessorKey: keyof T; // Chave usada para acessar o dado no objeto T
	cell?: (row: T) => any; // Função opcional para customizar a célula
	component?: any; // Componente opcional para renderização de células dinâmicas
	props?: (row: T) => Record<string, any>; // Propriedades opcionais passadas para o componente
	isFilterable: boolean; // Indica se a coluna pode ser filtrada
	class?: string; // Classe opcional para estilos
}

export interface HistoricoConferido {
    Filial: string;
    NF: string;           // Número da Nota Fiscal e série
    Vendedor: string;     // Vendedor (inclui nome e status)
    Cliente: string;      // Código, loja e nome do cliente
    Produto: string;      // Código do produto e descrição
    DataHora: string;     // Data e hora da movimentação
    Responsavel: string;  // Responsável pela conferência
    Placa: string;        // Placa do veículo
    Observacao: string;   // Observações
    DataConf: string;     // Data de confirmação
    Seq: string;          // Sequência
    Saldo: number;        // Quantidade/Saldo
}
  
// Interface original de borracharia
export interface borracharia {
	Filial: string;
	NF: string;
	Cliente: string;
	Produto: string; // Exemplo: "PA00534-P - 01 - PNEU 295/80 STRADA R"
	TipoMov: string;
	DataHora: string;
	Responsavel?: string; // Tornar opcional
	Placa: string;
	Observacao: string;
	Saldo: number;
	Emissao: string;
	actions?: () => void; // Ações como função opcional

	// **Adicionadas as novas propriedades como opcionais**
	CodigoProduto?: string;
	SubCodigoProduto?: string;
	DescricaoProduto?: string;
}

// Interface para itens da Nota Fiscal
export interface ItemNF {
	D2_ITEM: string;
	D2_COD: string; // Código do Produto conforme API
	B1_DESC: string; // Descrição do Produto conforme API
	SALDO: number;
	quantity: number; // Campo para armazenar a quantidade inserida pelo usuário
}

// Interface original de PreNota
export interface PreNota {
	Filial: string;
	NF: string;
	Status: string;
	Fornecedor: string;
	Emissao: string; // Corrigido o erro de digitação removendo 'a'
	Inclusao: string;
	Vencimento: string;
	Valor: string;
	Tipo: string;
	Prioridade: string;
	Usuario: string;
	Obs: string;
	Rec: string;
	actions?: () => void; // Ações como função opcional
}

export interface HistoricoData {
	Filial: string;
	NF: string;
	Cliente: string;
	Produto: string;
	TipoMov: string;
	DataHora: string;
	Responsavel: string;
	Placa: string;
	Observacao?: string;
	Saldo: number;
  }

// Interface para PreNotaTabela
export interface PreNotaTabela {
	'X-Filter-Filial': string;
	'X-Filter-NF': string;
	'X-Filter-Status': string;
	'X-Filter-Fornecedor': string;
	'X-Filter-Emissao': string;
	'X-Filter-Inclusao': string;
	'X-Filter-Vencimento': string;
	'X-Filter-Valor': string;
	'X-Filter-Tipo': string;
	'X-Filter-Prioridade': string;
	'X-Filter-Usuario': string;
	'X-Filter-Obs': string;
	'X-Filter-Rec': string;
}
