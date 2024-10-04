import Dexie from 'dexie';
const DB_NAME = 'rodoappDB';

interface PreNota {
	id: string;
	data: any;
}
interface OptionData {
	id: string;
	data: any;
}
class RodoAppDB extends Dexie {
	prenotasStore: Dexie.Table<PreNota, string>;
	optionsStore: Dexie.Table<OptionData, string>;

	constructor() {
		super(DB_NAME);

		// Versão 1: cria a prenotasStore
		this.version(1).stores({
			prenotasStore: 'id'
		});

		// Versão 2: adiciona optionsStore
		this.version(2).stores({
			prenotasStore: 'id',
			optionsStore: 'id'
		});

		this.prenotasStore = this.table('prenotasStore');
		this.optionsStore = this.table('optionsStore');
	}

	// Método para adicionar dinamicamente uma tabela
	addDynamicTable(tableName: string) {
		if (!this[tableName]) {
			this.version(this.verno + 1).stores({
				[tableName]: 'id'
			});
		}
		return this.table(tableName);
	}
}

// Instância do banco de dados
export const db = new RodoAppDB();

/**
 * Função para salvar dados em uma tabela dinâmica
 * @param {string} tableName - Nome da tabela a ser criada ou usada
 * @param {string} key - Chave dos dados
 * @param {any} data - Dados a serem salvos
 * @returns {Promise<void>}
 */

export async function saveOptionsData(tableName: string, key: string, data: any): Promise<void> {
	try {
		const table = db.addDynamicTable(tableName);
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
 * @returns {Promise<any | null>} - Retorna os dados encontrados ou null
 */
export async function getOptionsData(tableName: string, key: string): Promise<any | null> {
	try {
		const table = db.addDynamicTable(tableName);
		const record = await table.get(key);

		// Retorna o registro completo se encontrado
		return record ? record : null;
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
		console.log(`Dados com a chave '${key}' excluídos da tabela '${tableName}'.`);
	} catch (error) {
		console.error(`Erro ao excluir dados na tabela: ${tableName}, chave: ${key}`, error);
	}
}
