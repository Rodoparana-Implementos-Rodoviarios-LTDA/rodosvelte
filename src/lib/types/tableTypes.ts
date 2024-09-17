// src/lib/types/tableTypes.ts

// Coluna genérica usada em várias tabelas
export interface Column<T> {
	header: string; // O rótulo da coluna (cabeçalho)
	accessorKey: keyof T; // Chave usada para acessar o dado no objeto T
	cell?: (row: T) => any; // Função opcional para customizar a célula
	component?: any; // Componente opcional para renderização de células dinâmicas
	props?: (row: T) => Record<string, any>; // Propriedades opcionais passadas para o componente
	isFilterable?: boolean; // Indica se a coluna pode ser filtrada
}

// Interface genérica para dados da tabela
export interface TableData {
	[key: string]: any; // Dados da tabela com chave genérica
}

// Interface original de borracharia
export interface borracharia {
	Filial: string;
	NF: string;
	Cliente: string;
	Vendedor: string;
	Produto: string;
	Saldo: string;
	Emissao: string;
	actions?: () => void; // Ações como função opcional
}

// Interface original de PreNota
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
	Observacao: string;
	Saldo: number;
	actions?: () => void; // Função de ação personalizada
}
// Coluna genérica usada para a tabela

