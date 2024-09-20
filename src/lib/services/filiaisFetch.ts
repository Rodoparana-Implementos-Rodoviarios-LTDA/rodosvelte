// src/lib/fetchAndLoadData.ts
import { getData, saveData } from './idb'; // Ajuste o caminho conforme necessário

// Definição das interfaces
interface Filial {
  M0_CODFIL: string;
  M0_FILIAL: string;
}

interface UnidadeMedida {
  UM: string;
  DESCRICAO: string;
  DESC: string;
}

interface Condicao {
  Desc: string;
  E4_CODIGO: string;
  E4_DESCRI: string;
}

interface CentoCusto {
  CTT_CUSTO: string;
  CTT_DESC01: string;
  DESC: string;
}

interface ApiResponse {
  Filiais: Filial[];
  UnidadeMedida: UnidadeMedida[];
  Condicoes: Condicao[];
  CentoCusto: CentoCusto[]; // Propriedade com erro de tipografia
}

export async function loadData(): Promise<{
  Filiais: Filial[];
  UnidadeMedida: UnidadeMedida[];
  Condicoes: Condicao[];
  centroCusto: CentoCusto[]; // Propriedade corrigida
}> {
  // Verifica se estamos no client-side
  const isClient = typeof window !== 'undefined';
  if (!isClient) {
    throw new Error('IndexedDB não está disponível no server-side.');
  }

  try {
    // Tenta recuperar os dados do IndexedDB
    const [filiais, unidadeMedida, condicoes, centroCusto] = await Promise.all([
      getData('Filiais'),
      getData('UnidadeMedida'),
      getData('Condicoes'),
      getData('centroCusto'), // Recupera a chave corrigida
    ]);

    // Verifica se todos os dados estão presentes
    if (filiais && unidadeMedida && condicoes && centroCusto) {
      console.log('Dados recuperados do IndexedDB.');
      return {
        Filiais: filiais as Filial[],
        UnidadeMedida: unidadeMedida as UnidadeMedida[],
        Condicoes: condicoes as Condicao[],
        centroCusto: centroCusto as CentoCusto[],
      };
    } else {
      console.log('Dados ausentes no IndexedDB. Buscando da API...');
      // Faz o fetch dos dados da API
      const response = await fetch('http://protheus-app:8400/rest/reidoapsdu/consultar/cargaInicio');
      if (!response.ok) {
        throw new Error(`Erro na requisição: ${response.statusText}`);
      }

      const data: ApiResponse = await response.json();

      // Mapeia a propriedade "CentoCusto" para "centroCusto"
      const mappedData = {
        Filiais: data.Filiais,
        UnidadeMedida: data.UnidadeMedida,
        Condicoes: data.Condicoes,
        centroCusto: data.CentoCusto, // Mapeamento aqui
      };

      // Armazena cada conjunto de dados no IndexedDB
      await Promise.all([
        saveData('Filiais', mappedData.Filiais),
        saveData('UnidadeMedida', mappedData.UnidadeMedida),
        saveData('Condicoes', mappedData.Condicoes),
        saveData('centroCusto', mappedData.centroCusto), // Armazenamento com a chave corrigida
      ]);

      console.log('Dados armazenados no IndexedDB.');

      return mappedData;
    }
  } catch (error) {
    console.error('Erro ao carregar os dados:', error);
    throw error;
  }
}
