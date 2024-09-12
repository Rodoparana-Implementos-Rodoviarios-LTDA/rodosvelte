export interface Column<T> {
	header: string; // O rótulo da coluna (cabeçalho)
	accessorKey: keyof T; // Chave usada para acessar o dado no objeto T
	cell?: (row: T) => any; // Função opcional para customizar a célula
	component?: any; // Componente opcional para renderização de células dinâmicas
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
	avatarUrl: string;
	Obs: string;
	ObsRev: string;
	Revisao: string;
	Rec: string;
	actions?: () => void; // Ações como função opcional
}

export interface Borracharia {
	Filial: string;
	NF: string;
	Cliente: string;
	Vendedor: string;
	Produto: string;
	Saldo: string;
	Emissao: string;
}

export interface Portaria {
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
}
