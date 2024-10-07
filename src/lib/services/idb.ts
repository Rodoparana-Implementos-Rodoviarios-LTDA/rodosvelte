import Dexie, { type Table } from 'dexie';
const DB_NAME = 'rodoappDB';

interface OptionData {
	id: string;
	data: any;
}

class RodoAppDB extends Dexie {
	optionsStore: Table<OptionData, string>;

	constructor() {
		super(DB_NAME);

		// Versão 2: Adiciona optionsStore
		this.version(2).stores({
			optionsStore: 'id'
		});

		this.optionsStore = this.table('optionsStore');
	}

	// Método para adicionar dinamicamente uma tabela
	addDynamicTable<T>(tableName: string): Table<{ id: string; data: T }, string> {
		if (!this[tableName as keyof this]) {
			this.version(this.verno + 1).stores({
				[tableName]: 'id'
			});
		}
		// Retorna a tabela com a tipagem correta
		return this.table(tableName);
	}
}

// Instância do banco de dados
export const db = new RodoAppDB();

/**
 * Função para salvar dados em uma tabela dinâmica
 * @param {string} tableName - Nome da tabela a ser criada ou usada
 * @param {string} key - Chave dos dados
 * @param {T} data - Dados a serem salvos
 * @returns {Promise<void>}
 */
export async function saveOptionsData<T>(tableName: string, key: string, data: T): Promise<void> {
	try {
		const table = db.addDynamicTable<T>(tableName);
		// Salva um objeto com 'id' e 'data'
		await table.put({ id: key, data });
		console.log(`Dados salvos com sucesso na tabela: ${tableName}, chave: ${key}`);
	} catch (error) {
		console.error(`Erro ao salvar dados na tabela: ${tableName}, chave: ${key}`, error);
	}
}

/**
 * Função para buscar dados de uma tabela dinâmica
 * @param {string} tableName - Nome da tabela a ser usada
 * @param {string} key - Chave dos dados a serem buscados
 * @returns {Promise<T | null>} - Retorna os dados encontrados ou null
 */
export async function getOptionsData<T>(tableName: string, key: string): Promise<T | null> {
	try {
		const table = db.addDynamicTable<T>(tableName);

		// Se a chave for 'all', retorna todos os itens da tabela
		if (key === 'all') {
			const allRecords = await table.toArray();
			return allRecords as unknown as T;
		}

		// Retorna o registro completo se encontrado
		const record = await table.get(key);
		return record ? (record as unknown as T) : null;
	} catch (error) {
		console.error(`Erro ao buscar dados na tabela: ${tableName}, chave: ${key}`, error);
		return null;
	}
}

/**
 * Função para excluir um registro específico de uma tabela dinâmica
 * @param {string} tableName - Nome da tabela
 * @param {string} key - Chave do registro a ser excluído
 * @returns {Promise<void>}
 */
export async function deleteData(tableName: string, key: string): Promise<void> {
	try {
		const table = db.addDynamicTable(tableName);
		await table.delete(key);
		console.log(`Registro com chave ${key} excluído da tabela ${tableName}`);
	} catch (error) {
		console.error(`Erro ao excluir o registro ${key} da tabela ${tableName}`, error);
	}
}

/**
 * Função para excluir todo o banco de dados IndexedDB
 */
export async function clearIndexedDB(): Promise<void> {
	try {
		await db.delete(); // Deleta o banco de dados inteiro
		console.log('IndexedDB excluído com sucesso.');
	} catch (error) {
		console.error('Erro ao excluir o IndexedDB:', error);
	}
}
