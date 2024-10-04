// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}
export interface Filial {
	M0_CODFIL: string;
	M0_FILIAL: string;
}

export interface UnidadeMedida {
	UM: string;
	DESCRICAO: string;
	DESC: string;
}

export interface Condicao {
	Desc: string;
	E4_CODIGO: string;
	E4_DESCRI: string;
}

export interface CentroCusto {
	CTT_CUSTO: string;
	CTT_DESC01: string;
	DESC: string;
}

export interface ApiResponse {
	Filiais: Filial[];
	UnidadeMedida: UnidadeMedida[];
	Condicoes: Condicao[];
	CentoCusto: CentoCusto[];
}
export {};
