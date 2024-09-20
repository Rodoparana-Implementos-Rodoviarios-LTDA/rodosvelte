// src/lib/types/tableTypes.ts

// src/lib/types.ts
export interface DetalhesXML {
    numero: string;
    serie: string;
    dataEmissao: string;
    valorTotalDaNota: string;
    nomeEmitente: string;
    cnpjEmitente: string;
    ufEmitente: string;
    nomeDestinatario: string;
    cnpjDestinatario: string;
    ufDestinatario: string;
    informacoesAdicionais: string;
    itens: Item[];
    filialName?: string; // Nome da filial após triangulação
}

export interface Item {
    codProduto: string;
    descProduto: string;
    ncmsh: string;
    cst: string;
    origem: string;
    cfop: string;
    unidade: string;
    quantidade: string;
    valorUnitario: string;
    valorTotal: string;
    bcIcms: string;
    valorIcms: string;
    valorIpi: string;
    aliqIcms: string;
    aliqIpi: string;
}

export interface Filial {
    numero: string;
    filial: string;
    cnpjFilial: string;
}

export interface Column<T> {
	header: string;
	accessorKey: keyof T;
	cell?: (row: T) => any;
	component?: any;
	props?: (row: T) => Record<string, any>;
	isFilterable: boolean; // Agora isFilterable é sempre booleano
	class?:string;
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
