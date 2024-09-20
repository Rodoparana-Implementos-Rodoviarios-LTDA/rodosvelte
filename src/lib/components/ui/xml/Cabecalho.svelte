<script>
	import Combobox from '$lib/components/ui/Combobox.svelte';
	import {
		IconDownload,
		IconReportSearch,
		IconGitFork,
		IconPaperclip,
		IconArrowNarrowLeft,
		IconDeviceFloppy,
		IconPencilUp
	} from '@tabler/icons-svelte';

	let tipoDeNFOptions = [
		{ label: 'Revenda', value: 'revenda' },
		{ label: 'Despesa/Imobilizado', value: 'despesa_imobilizado' },
		{ label: 'Matéria Prima', value: 'materia_prima' },
		{ label: 'Collection', value: 'collection' }
	];

	let prioridadeOptions = [
		{ label: 'Baixa', value: 'baixa' },
		{ label: 'Média', value: 'media' },
		{ label: 'Alta', value: 'alta' }
	];
	let selectedTipoDeNF = null;
	let selectedPrioridade = null;

	const handleTipoDeNFSelect = (event) => {
		selectedTipoDeNF = event.detail.selected;
		console.log('Tipo de NF selecionado:', selectedTipoDeNF);
	};

	const handlePrioridadeSelect = (event) => {
		selectedPrioridade = event.detail.selected;
		console.log('Prioridade selecionada:', selectedPrioridade);
	};
</script>

<main class="h-full w-full flex justify-start items-start bg-base-300 p-4 card gap-5">
	<div id="firstLine" class="flex items-center justify-between gap-3 w-full">
		<div id="leftInput" class="flex gap-5">
			<label class="input input-bordered flex items-center gap-2 w-[30vw]">
				<input type="text" class="grow" placeholder="Buscar XML" />
				<IconReportSearch />
			</label>
			<div class=" join border border-primary">
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Baixar DANFE"
				>
					<IconDownload /></button
				>
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Rateio"
				>
					<IconGitFork /></button
				>
				<button
					class="tooltip tooltip-info tooltip-bottom btn hover:bg-primary hover:text-primary-content"
					data-tip="Anexos"
				>
					<IconPaperclip /></button
				>
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
				<IconArrowNarrowLeft /></button
			>
			<button
				class="tooltip tooltip-success tooltip-bottom btn btn-success hover:text-success-content"
				data-tip="Salvar"
			>
				<IconDeviceFloppy /></button
			>
		</div>
	</div>
	<div id="secondLine" class="flex w-full h-full gap-5">
		<div id="detailsXML" class="card w-[30vw] h-full bg-base-100 p-5 gap-5">
			<div class="flex justify-between">
				<div class="flex gap-2 items-center justify-start">
					<span class="text-lg font-medium">FORNECEDOR:</span><span class="text-sm opacity-70"
						>(071197-10) -SYMA COMPUTADORES LTDA - 04912543000136</span
					>
				</div>
				<div class="flex gap-2 items-center justify-start">
					<span class="text-lg font-medium">NF:</span><span class="text-sm opacity-70"
						>1069022 - 1</span
					>
				</div>
			</div>
			<div class="flex justify-between">
				<div class="flex gap-2 items-center justify-start">
					<span class="text-lg font-medium">FILIAL:</span><span class="text-sm opacity-70"
						>RODOPARANA Curitiba</span
					>
				</div>
				<div class="flex gap-2 items-center justify-start">
					<span class="text-lg font-medium">DATA:</span><span class="opacity-70"
						>16/09/2024 - 14:14:39</span
					>
				</div>
			</div>
			<div class="collapse collapse-arrow bg-base-300">
				<input type="checkbox" />
				<div class="collapse-title text-lg font-medium w-full flex justify-center">
					Detalhes da XML
				</div>
				<div class="collapse-content">
					<p>
						";R.G. (Destinatario): 1019924897;OC 013852;Valor aproximado tributos R$4.379,43
						(25,73%) Fonte: IBPT (Lei Federal 12741/2012)"
					</p>
				</div>
			</div>
		</div>
		<div id="Inputs" class="card w-[35vw] p-5 gap-5 bg-base-100 justify-evenly">
			<div class="flex justify-center"><span class="font-bold text-lg">Adicionar Informações sobre a NF</span></div>
			<div class="flex w-full justify-evenly gap-5">
				<div class="w-full">
					<Combobox
						buttonClass="bg-base-300"
						options={pedidoOptions}
						placeholder="Pedido"
						on:select={handlePedidoSelect}
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
						options={tipoDeNFOptions}
						placeholder="Tipo de NF"
						on:select={handleTipoDeNFSelect}
						buttonClass="bg-base-300"
					/>
				</div>
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
