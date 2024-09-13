// src/lib/types/tableTypes.ts

export interface Column<T> {
	header: string;
	accessorKey: keyof T;
	cell?: (row: T) => any;
	component?: any;
	props?: (row: T) => Record<string, any>;
	isFilterable: boolean; // Agora isFilterable é sempre booleano
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
	Rec: string;
	actions?: () => void; // Ações como função opcional
}

export interface borracharia {
	Filial: string;
	NF: string;
	Cliente: string;
	Vendedor: string;
	Produto: string;
	Saldo: string;
	Emissao: string;
}

export interface BorrachariaData {
	Filial: string;
	NF: string;
	Vendedor: string;
	Cliente: string;
	Produto: string;
	Saldo: number;
	Emissao: string;
	Responsavel: string;
	Placa: string;
	Observacao: string;
}

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
