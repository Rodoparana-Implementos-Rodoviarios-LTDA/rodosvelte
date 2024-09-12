// revisaoFetch.ts
export default async function revisaoFetch(rec: string) {
    const response = await fetch('http://protheus-app:8400/rest/reidoapsdu/getHistorico', {
      method: 'GET',
      headers: {
        'RECSF1': rec,
        'Content-Type': 'application/json'
      }
    });
  
    if (!response.ok) {
      throw new Error('Erro ao buscar histórico');
    }
  
    return await response.json();
  }
  