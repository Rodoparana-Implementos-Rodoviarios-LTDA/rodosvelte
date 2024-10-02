<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { getOptionsData } from '$lib/services/idb'; // Função para buscar dados do IndexedDB
	import type { DetalhesXML } from '$lib/types';

	export let xmlKey: string;
	let detalhesXML: DetalhesXML | null = null;
	let loadingDetalhes: boolean = false;
	let errorDetalhes: string | null = null;

	// Função para buscar os dados do IndexedDB
	async function fetchDetalhesFromIndexedDB(): Promise<void> {
		try {
			loadingDetalhes = true;
			errorDetalhes = null;

			// Busca os dados na tabela 'xml' usando a chave fornecida
			const result = await getOptionsData('xml', xmlKey);

			if (result && result.data) {
				// Acessa os dados dentro do campo 'data'
				detalhesXML = result.data;
				console.log('Dados obtidos do IndexedDB:', detalhesXML);
				errorDetalhes = null; // Nenhum erro
			} else {
				throw new Error('Nenhum dado encontrado no IndexedDB.');
			}
		} catch (err) {
			errorDetalhes = (err as Error).message;
			console.error('Erro ao buscar dados no IndexedDB:', err);
		} finally {
			loadingDetalhes = false;
		}
	}

	// Escutar o evento 'xmlDataUpdated' para buscar os dados novamente
	function handleXmlDataUpdated(event: Event) {
		const customEvent = event as CustomEvent; // Faz o cast para CustomEvent
		console.log(`Evento 'xmlDataUpdated' capturado com detalhe:`, customEvent.detail);

		// Busca os dados do IndexedDB
		fetchDetalhesFromIndexedDB().then(() => {
			if (!errorDetalhes) {
				console.log('Busca de dados do XML concluída com sucesso!');
			} else {
				console.log('Erro ao buscar os dados do XML:', errorDetalhes);
			}
		});
	}

	// Hook que será executado quando o componente for montado
	onMount(() => {
		// Tenta recuperar a chave XML do LocalStorage se não for fornecida
		if (!xmlKey) {
			const storedKey = localStorage.getItem('xmlKey');
			if (storedKey) {
				xmlKey = storedKey;
			}
		}

		// Se houver uma chave XML, busca os dados
		if (xmlKey) {
			fetchDetalhesFromIndexedDB();
		}

		// Adiciona o listener para o evento customizado
		window.addEventListener('xmlDataUpdated', handleXmlDataUpdated);

		// Limpeza do evento quando o componente for desmontado
		onDestroy(() => {
			window.removeEventListener('xmlDataUpdated', handleXmlDataUpdated);
		});
	});
</script>

{#if loadingDetalhes}
	<div class="flex w-52 flex-col gap-4">
		<div class="flex items-center gap-4">
			<div class="skeleton h-16 w-16 shrink-0 rounded-full"></div>
			<div class="flex flex-col gap-4">
				<div class="skeleton h-4 w-20"></div>
				<div class="skeleton h-4 w-28"></div>
			</div>
		</div>
		<div class="skeleton h-32 w-full"></div>
	</div>
{:else if errorDetalhes}
	<p class="text-red-500">Erro: {errorDetalhes}</p>
{:else if detalhesXML}
	<div class="flex justify-between">
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">FORNECEDOR:</span>
			<div class="flex flex-col">
				<span class="text-sm opacity-70">{detalhesXML.nomeEmitente}</span>
				<span class="text-xs opacity-70"> {detalhesXML.cnpjEmitente}</span>
			</div>
		</div>
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">NF:</span>
			<span class="text-sm opacity-70 truncate">{detalhesXML.numero} - {detalhesXML.serie}</span>
		</div>
	</div>
	<div class="flex justify-between">
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">FILIAL:</span>
			<div class="flex flex-col">
				<span class="text-sm opacity-70"
					>{detalhesXML.filialName || detalhesXML.nomeDestinatario}</span
				>
				<span class="text-xs opacity-70"> {detalhesXML.cnpjDestinatario}</span>
			</div>
		</div>
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">DATA:</span>
			<span class="opacity-70">{detalhesXML.dataEmissao}</span>
		</div>
	</div>
	<div class="collapse collapse-arrow bg-base-300">
		<input type="checkbox" />
		<div class="collapse-title text-lg font-medium w-full flex justify-center">Detalhes da XML</div>
		<div class="collapse-content">
			<p>{detalhesXML.informacoesAdicionais}</p>
		</div>
	</div>
{:else}
	<p>Insira uma chave XML para buscar os detalhes.</p>
{/if}
