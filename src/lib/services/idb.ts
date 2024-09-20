  const DB_NAME = 'rodoappDB';
  const STORE_NAME = 'prenotasStore';

  // Verifica se estamos no client-side
  const isClient = typeof window !== 'undefined';

  // Função para abrir ou criar o IndexedDB
  const openDB = (): Promise<IDBDatabase> => {
    return new Promise((resolve, reject) => {
      if (!isClient) {
        return reject('IndexedDB não está disponível no server-side.');
      }

      const request = indexedDB.open(DB_NAME, 1);

      // Cria a store no primeiro acesso
      request.onupgradeneeded = function () {
        const db = request.result;
        if (!db.objectStoreNames.contains(STORE_NAME)) {
          db.createObjectStore(STORE_NAME, { keyPath: 'id' });
        }
      };

      request.onsuccess = () => resolve(request.result);
      request.onerror = (event) => reject(`Erro ao abrir o IndexedDB: ${event.target.error}`);
    });
  };

  // Função para buscar dados do IndexedDB
  export async function getData(key: string): Promise<any> {
    if (!isClient) {
      console.warn('IndexedDB não está disponível no server-side');
      return null;
    }

    const db = await openDB();
    const transaction = db.transaction(STORE_NAME, 'readonly');
    const store = transaction.objectStore(STORE_NAME);
    const request = store.get(key);

    return new Promise((resolve, reject) => {
      request.onsuccess = () => {
        resolve(request.result?.data || null);
      };
      request.onerror = () => {
        reject(`Erro ao buscar dados com a chave: ${key}`);
      };
    });
  }

  // Função para salvar dados no IndexedDB
  export async function saveData(key: string, data: any): Promise<void> {
    const db = await openDB();
    const transaction = db.transaction(STORE_NAME, 'readwrite');
    const store = transaction.objectStore(STORE_NAME);

    return new Promise((resolve, reject) => {
      const request = store.put({ id: key, data });
      
      request.onsuccess = () => {
        transaction.oncomplete = () => resolve();
      };
      
      request.onerror = () => {
        reject(`Erro ao salvar dados com a chave: ${key}`);
      };
    });
  }

  // Função para deletar dados do IndexedDB
  export async function deleteData(key: string): Promise<void> {
    const db = await openDB();
    const transaction = db.transaction(STORE_NAME, 'readwrite');
    const store = transaction.objectStore(STORE_NAME);

    return new Promise((resolve, reject) => {
      const request = store.delete(key);
      
      request.onsuccess = () => {
        transaction.oncomplete = () => resolve();
      };
      
      request.onerror = () => {
        reject(`Erro ao deletar dados com a chave: ${key}`);
      };
    });
  }
