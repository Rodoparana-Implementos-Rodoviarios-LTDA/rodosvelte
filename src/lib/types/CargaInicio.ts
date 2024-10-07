// src/lib/types/CargaInicio.ts

export interface UnidadeMedida {
	UM: string;
	DESCRICAO: string;
	DESC: string;
}
export interface CondicaoOption {
	label: string;
	value: string;
}
export interface Condicao {
	Desc: string;
	E4_CODIGO: string;
	E4_DESCRI: string;
}

export interface CentoCusto {
	CTT_CUSTO: string;
	CTT_DESC01: string;
	DESC: string;
}

export interface CargaInicio {
	UnidadeMedida: UnidadeMedida[];
	Condicoes: Condicao[];
	CentoCusto: CentoCusto[];
	// Adicione outros campos conforme necessário
}
