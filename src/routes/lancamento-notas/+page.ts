// src/routes/lancamento-notas/+page.ts
import { db } from '$lib/services/idb';

/** @type {import('./$types').PageLoad} */
export async function load() {
    // Inicializar o IndexedDB e garantir que as tabelas existam
    try {
        // Verifica se o IndexedDB já está aberto
        if (!db.isOpen()) {
            // Abre e inicializa o banco de dados
            await db.open();
            console.log('IndexedDB inicializado com sucesso.');
        } else {
            console.log('IndexedDB já está aberto.');
        }

        // Se necessário, você pode carregar ou inicializar dados aqui
        // Exemplo: Buscar dados iniciais da tabela optionsStore (se existir lógica para isso)

        return {
            status: 'success',
        };
    } catch (error) {
        console.error('Erro ao inicializar o IndexedDB:', error);
        return {
            status: 'error',
        };
    }
}
