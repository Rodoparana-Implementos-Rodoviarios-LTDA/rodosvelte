// src/lib/types/ConexaoNFE.ts

export interface ConexaoNFEItem {
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

export interface ConexaoNFE {
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
	itens: ConexaoNFEItem[];
}