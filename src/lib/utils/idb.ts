import { openDB } from 'idb';

let dbPromise: any;

// Verifica se o código está rodando no navegador
if (typeof window !== 'undefined') {
  dbPromise = openDB('login-db', 1, {
    upgrade(db) {
      db.createObjectStore('loginStore', { keyPath: 'id' });
    },
  });
}

// Função para salvar os dados (username e token)
export async function saveLoginData(username: string, token: string): Promise<void> {
  if (!dbPromise) return; // Se não estiver no cliente, sai da função
  const db = await dbPromise;
  await db.put('loginStore', { id: 'auth', username, token });
}

// Função para obter os dados
export async function getLoginData(): Promise<{ username: string; token: string } | undefined> {
  if (!dbPromise) return; // Se não estiver no cliente, sai da função
  const db = await dbPromise;
  return await db.get('loginStore', 'auth');
}

// Função para remover os dados de login
export async function clearLoginData(): Promise<void> {
  if (!dbPromise) return; // Se não estiver no cliente, sai da função
  const db = await dbPromise;
  await db.delete('loginStore', 'auth');
}
