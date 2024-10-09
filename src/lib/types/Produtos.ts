export interface Produto {
	codProduto: string;
	descProduto: string;
	ncmsh: string;
	cst: string;
	origem: string;
	cfop: string;
	unidade: string;
	quantidade: number;
	valorUnitario: number;
	valorTotal: number;
	valorIcms?: number;
	aliqIcms?: number;
	valorIpi?: number;
	aliqIpi?: number;
	selectedProduto?: string; // Adiciona a seleção do produto
}

// Define o tipo das opções de produtos
export interface ProdutoOption {
	label: string;
	value: string;
	B1_ORIGEM: string;
	B1_POSIPI: string;
}
