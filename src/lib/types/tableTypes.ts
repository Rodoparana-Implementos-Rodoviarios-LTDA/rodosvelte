// src/lib/types/tableTypes.ts

<<<<<<< HEAD
// Coluna genérica usada em várias tabelas
export interface Column<T> {
	header: string; // O rótulo da coluna (cabeçalho)
	accessorKey: keyof T; // Chave usada para acessar o dado no objeto T
	cell?: (row: T) => any; // Função opcional para customizar a célula
	component?: any; // Componente opcional para renderização de células dinâmicas
	props?: (row: T) => Record<string, any>; // Propriedades opcionais passadas para o componente
	isFilterable?: boolean; // Indica se a coluna pode ser filtrada
=======
export interface Column<T> {
	header: string;
	accessorKey: keyof T;
	cell?: (row: T) => any;
	component?: any;
	props?: (row: T) => Record<string, any>;
	isFilterable: boolean; // Agora isFilterable é sempre booleano
	class?:string;
>>>>>>> fe242afe9d703777fb5b91b9ff3a57e2868c792b
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
<<<<<<< HEAD
// Coluna genérica usada para a tabela

=======

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
>>>>>>> fe242afe9d703777fb5b91b9ff3a57e2868c792b
