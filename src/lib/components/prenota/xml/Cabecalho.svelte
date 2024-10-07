<script lang="ts">
	import { fetchAndSaveConexaoNFE } from '$lib/services/conexaoNFE';
	import { createEventDispatcher } from 'svelte';

	import XMLDetails from './XMLDetails.svelte'; // Novo componente que exibe detalhes
	import { IconReportSearch, IconDownload, IconGitFork, IconPaperclip, IconPencilUp, IconArrowNarrowLeft, IconDeviceFloppy } from '@tabler/icons-svelte';
	import ComboboxesSection from './ComboboxesSection.svelte';
	const dispatch = createEventDispatcher();

	// Variáveis para a busca do XML
	let xmlKey: string = '';
	let loadingDetalhes: boolean = false;
	let errorDetalhes: string | null = null;
	let xmlData: any = null; // Armazenamento local dos detalhes do XML

	const handleBuscarXML = async () => {
		if (!xmlKey.trim()) {
			console.warn('A chave XML está vazia.');
			return;
		}

		loadingDetalhes = true; // Inicia o estado de carregamento
		errorDetalhes = null;

		try {
			// Busca os detalhes do XML diretamente da API
			const data = await fetchAndSaveConexaoNFE(xmlKey.trim());
			dispatch('chaveXML', { xmlKey });

			// Armazena os dados localmente para exibição
			xmlData = data;
		} catch (error) {
			errorDetalhes = (error as Error).message;
			console.error('Erro ao buscar detalhes do XML:', errorDetalhes);
		} finally {
			loadingDetalhes = false; // Finaliza o estado de carregamento
		}
	};
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
				<button on:click={handleBuscarXML} disabled={loadingDetalhes}>
					<IconReportSearch />
				</button>
			</label>
			<div class="join border border-primary">
				<!-- Botões adicionais -->
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
			<!-- Botões à direita -->
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
				disabled={!xmlData}
			>
				<IconDeviceFloppy />
			</button>
		</div>
	</div>

	<div id="secondLine" class="flex w-full h-full gap-5">
		<!-- Usamos o componente XMLDetails e passamos os dados -->
		<div id="detailsXML" class="card w-[40vw] h-full bg-base-100 p-5 gap-5">
			<XMLDetails {loadingDetalhes} />
			<!-- Passando a prop de loading para XMLDetails -->
		</div>

		<!-- Incluímos o componente ComboboxesSection -->
		<ComboboxesSection />

		<div>
			<textarea
				placeholder="Adicione suas Observações..."
				class="p-5 text-lg textarea w-[25vw] h-full bg-base-100"
			/>
		</div>
	</div>
</main>
