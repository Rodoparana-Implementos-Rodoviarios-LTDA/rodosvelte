import { saveData } from "./idb";

// Função para buscar as filiais da API e salvar no IndexedDB
export async function fetchAndSaveFiliais(): Promise<void> {
    const url = 'http://protheus-app:8400/rest/reidoapsdu/consultar/filiais/';
    
    try {
      const response = await fetch(url);
      if (!response.ok) {
        throw new Error(`Erro ao buscar filiais da API: ${response.status}`);
      }
  
      const filiais = await response.json();

      // Aplica trim() aos campos relevantes (número, filial, cnpjFilial)
      const filiaisTratadas = filiais.map((filial: any) => ({
        numero: filial.numero.trim(),
        filial: filial.filial.trim(),
        cnpjFilial: filial.cnpjFilial.trim(),
      }));

      // Salva as filiais tratadas no IndexedDB
      await saveData('filiais', filiaisTratadas);
      console.log('Filiais tratadas e salvas no IndexedDB com sucesso');
    } catch (error) {
      console.error(`Erro ao buscar ou salvar filiais: ${error}`);
    }
}
