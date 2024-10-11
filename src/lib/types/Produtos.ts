export interface Produto {
	codProduto: string; // Código do produto
	descProduto: string; // Descrição do produto
	ncmsh: string; // NCM
	cst: string; // Código da Situação Tributária
	origem: string; // Origem do produto
	cfop: string; // Código Fiscal de Operações e Prestações
	unidade: string; // Unidade de medida
	quantidade: number; // Quantidade do produto
	valorUnitario: number; // Valor unitário do produto
	valorTotal: number; // Valor total do produto
	valorIcms?: number; // Valor do ICMS (opcional)
	aliqIcms?: number; // Alíquota do ICMS (opcional)
	valorIpi?: number; // Valor do IPI (opcional)
	aliqIpi?: number; // Alíquota do IPI (opcional)
	selectedProduto?: string; // Produto selecionado
}

export interface ComboboxOption {
	label: string; // Nome do produto a ser exibido no Combobox
	value: string; // Valor do produto a ser usado internamente
	campo2: string; // Origem do produto
	campo1: string; // NCM do produto
}
