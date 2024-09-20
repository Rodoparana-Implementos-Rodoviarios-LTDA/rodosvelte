<script lang="ts">
	import { onMount } from 'svelte';
	import Combobox from '$lib/components/ui/Combobox.svelte';
	import { getData } from '$lib/services/idb'; // Importa a função getData do idb
	import { fetchAndLoadData } from '$lib/services/generalFetch'; // Importa a função de fetch
	import { fetchDetalhesXML } from '$lib/services/conexaoNFE'; // Importa a função de conexão com a API
	import {
		IconDownload,
		IconReportSearch,
		IconGitFork,
		IconPaperclip,
		IconArrowNarrowLeft,
		IconDeviceFloppy,
		IconPencilUp
	} from '@tabler/icons-svelte';
	
	import type { Condicao } from '$lib/../app'; // Importa a interface Condicao
	
	// Interface para os detalhes do XML
	interface DetalhesXML {
		numero: string;
		serie: string;
		dataEmissao: string;
		valorTotalDaNota: string;
		nomeEmitente: string;
		cnpjEmitente: string;
		ufEmitente: string;
		nomeDestinatario: string;
		cnpjDestinatario: string;
		ufDestinatario: string;
		informacoesAdicionais: string;
		itens: Item[];
	}
	
	interface Item {
		codProduto: string;
		descProduto: string;
		ncmsh: string;
		cst: string;
		origem: string;
		cfop: string;
		unidade: string;
		quantidade: string;
		valorUnitario: string;
		valorTotal: string;
		bcIcms: string;
		valorIcms: string;
		valorIpi: string;
		aliqIcms: string;
		aliqIpi: string;
	}
	
	let selectedTipoDeNF: string | null = null;
	let selectedPrioridade: string | null = null;
	let selectedFormaPagamento: string | null = null; // Estado para a seleção do Combobox de Forma de Pagamento
	
	let condicaoOptions: { label: string; value: string }[] = []; // Opções para o Combobox de Forma de Pagamento
	let tipoDeNFOptions: { label: string; value: string }[] = []; // Opções para o Combobox de Tipo de NF
	let prioridadeOptions: { label: string; value: string }[] = []; // Opções para o Combobox de Prioridade
	
	let loadingCondicoes: boolean = true;
	let errorCondicoes: string | null = null;
	
	// Variáveis para a busca do XML
	let xmlKey: string = '';
	let detalhesXML: DetalhesXML | null = null;
	let loadingDetalhes: boolean = false;
	let errorDetalhes: string | null = null;
	
	const handleTipoDeNFSelect = (event: CustomEvent<{ selected: string }>) => {
		selectedTipoDeNF = event.detail.selected;
		console.log('Tipo de NF selecionado:', selectedTipoDeNF);
	};
	
	const handlePrioridadeSelect = (event: CustomEvent<{ selected: string }>) => {
		selectedPrioridade = event.detail.selected;
		console.log('Prioridade selecionada:', selectedPrioridade);
	};
	
	const handleCondicaoSelect = (event: CustomEvent<{ selected: string }>) => {
		selectedFormaPagamento = event.detail.selected;
		console.log('Forma de Pagamento selecionada:', selectedFormaPagamento);
	};
	
	const handleBuscarXML = async () => {
		if (!xmlKey.trim()) {
			console.warn('A chave XML está vazia.');
			return;
		}
		
		loadingDetalhes = true;
		errorDetalhes = null;
		detalhesXML = null;
		
		try {
			detalhesXML = await fetchDetalhesXML(xmlKey.trim());
			console.log('Detalhes do XML:', detalhesXML);
		} catch (error) {
			errorDetalhes = (error as Error).message;
			console.error('Erro ao buscar detalhes do XML:', errorDetalhes);
		} finally {
			loadingDetalhes = false;
		}
	};
	
	onMount(async () => {
		try {
			// Tenta recuperar os dados de Condicoes, tipoDeNFOptions e prioridadeOptions do IndexedDB
			let [condicoes, tipoNF, prioridade] = await Promise.all([
				getData('Condicoes'),
				getData('tipoDeNFOptions'),
				getData('prioridadeOptions')
			]);
			
			// Se algum dado estiver ausente, faz o fetch e salva
			if (!condicoes || !tipoNF || !prioridade) {
				await fetchAndLoadData();
				condicoes = await getData('Condicoes');
				tipoNF = await getData('tipoDeNFOptions');
				prioridade = await getData('prioridadeOptions');
			}
			
			// Mapeia os dados de Condicoes para as opções do Combobox
			if (condicoes && Array.isArray(condicoes)) {
				condicaoOptions = condicoes.map((condicao: Condicao) => ({
					label: condicao.Desc.trim(),
					value: condicao.E4_CODIGO.trim()
				}));
				console.log('Condicoes recuperadas:', condicaoOptions);
			} else {
				console.warn('Nenhum dado de Condicoes encontrado.');
			}
			
			// Atribui as opções de Tipo de NF
			if (tipoNF && Array.isArray(tipoNF)) {
				tipoDeNFOptions = tipoNF;
				console.log('Tipo de NF recuperado:', tipoDeNFOptions);
			} else {
				console.warn('Nenhum dado de Tipo de NF encontrado.');
			}
			
			// Atribui as opções de Prioridade
			if (prioridade && Array.isArray(prioridade)) {
				prioridadeOptions = prioridade;
				console.log('Prioridade recuperada:', prioridadeOptions);
			} else {
				console.warn('Nenhuma prioridade encontrada.');
			}
		} catch (error) {
			errorCondicoes = (error as Error).message;
			console.error('Erro ao recuperar Condicoes:', errorCondicoes);
		} finally {
			loadingCondicoes = false;
		}
	});
</script>

<main class="h-full w-full flex justify-start items-start bg-base-300 p-4 card gap-5">
	<div id="firstLine" class="flex items-center justify-between gap-3 w-full">
		<div id="leftInput" class="flex gap-5">
			<label class="input input-bordered flex items-center gap-2 w-[30vw]">
				<input
					type="text"
					class="grow"
					placeholder="Buscar XML"
					bind:value={xmlKey}
					on:keydown={(e) => e.key === 'Enter' && handleBuscarXML()}
				/>
				<button on:click={handleBuscarXML}>
					<IconReportSearch />
				</button>
			</label>
			<div class="join border border-primary">
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Baixar DANFE"
				>
					<IconDownload />
				</button>
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Rateio"
				>
					<IconGitFork />
				</button>
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Anexos"
				>
					<IconPaperclip />
				</button>
			</div>
		</div>
		
		<div id="rightButtons" class="flex gap-2">
			<button class="tooltip tooltip-bottom btn btn-outline" data-tip="Incluir Nota Manual">
				<a href="/lancamento-notas/incluir"><IconPencilUp /></a>
			</button>
			<button
				class="tooltip tooltip-error tooltip-bottom btn btn-error hover:text-error-content"
				data-tip="Voltar"
			>
				<IconArrowNarrowLeft />
			</button>
			<button
				class="tooltip tooltip-success tooltip-bottom btn btn-success hover:text-success-content"
				data-tip="Salvar"
			>
				<IconDeviceFloppy />
			</button>
		</div>
	</div>
	<div id="secondLine" class="flex w-full h-full gap-5">
		<div id="detailsXML" class="card w-[30vw] h-full bg-base-100 p-5 gap-5">
			{#if loadingDetalhes}
				<p>Carregando detalhes do XML...</p>
			{:else if errorDetalhes}
				<p class="text-red-500">Erro: {errorDetalhes}</p>
			{:else if detalhesXML}
				<div class="flex justify-between">
					<div class="flex gap-2 items-center justify-start">
						<span class="text-lg font-medium">FORNECEDOR:</span>
						<span class="text-sm opacity-70">{detalhesXML.nomeEmitente} - {detalhesXML.cnpjEmitente}</span>
					</div>
					<div class="flex gap-2 items-center justify-start">
						<span class="text-lg font-medium">NF:</span>
						<span class="text-sm opacity-70">{detalhesXML.numero} - {detalhesXML.serie}</span>
					</div>
				</div>
				<div class="flex justify-between">
					<div class="flex gap-2 items-center justify-start">
						<span class="text-lg font-medium">FILIAL:</span>
						<span class="text-sm opacity-70">{detalhesXML.nomeDestinatario} - {detalhesXML.cnpjDestinatario}</span>
					</div>
					<div class="flex gap-2 items-center justify-start">
						<span class="text-lg font-medium">DATA:</span>
						<span class="opacity-70">{detalhesXML.dataEmissao}</span>
					</div>
				</div>
				<div class="collapse collapse-arrow bg-base-300">
					<input type="checkbox" />
					<div class="collapse-title text-lg font-medium w-full flex justify-center">
						Detalhes da XML
					</div>
					<div class="collapse-content">
						<p>{detalhesXML.informacoesAdicionais}</p>
					</div>
				</div>
			{:else}
				<p>Insira uma chave XML para buscar os detalhes.</p>
			{/if}
		</div>
		<div id="Inputs" class="card w-[35vw] p-5 gap-5 bg-base-100 justify-evenly">
			<div class="flex justify-center">
				<span class="font-bold text-lg">Adicionar Informações sobre a NF</span>
			</div>
			<div class="flex w-full justify-evenly gap-5">
				<div class="w-full">
					<Combobox
						buttonClass="bg-base-300"
						options={tipoDeNFOptions}
						placeholder="Tipo de NF"
						on:select={handleTipoDeNFSelect}
					/>
				</div>
				<div class="w-full">
					<Combobox
						options={condicaoOptions}
						placeholder="Forma de Pagamento"
						on:select={handleCondicaoSelect}
						buttonClass="bg-base-300"
					/>
				</div>
			</div>
			<div class="flex w-full justify-evenly gap-5">
				<div class="w-full">
					<Combobox
						options={prioridadeOptions}
						placeholder="Prioridade"
						on:select={handlePrioridadeSelect}
						buttonClass="bg-base-300"
					/>
				</div>
			</div>
		</div>
		
		<div>
			<textarea
				placeholder="Adicione suas Observações..."
				class="p-5 text-lg textarea w-[35vw] h-full bg-base-100"
			/>
		</div>
	</div>
</main>
