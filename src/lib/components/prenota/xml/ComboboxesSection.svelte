<script lang="ts">
	import Combobox from '$lib/components/ui/Combobox.svelte';
	import {
		tiposNFOptions,
		prioridadeOptions,
		condicoesStore,
		selectedTipoNF,
		selectedPriority,
		selectedCondicao
	} from '$lib/stores/xmlStore';
	import { fetchProdutos } from '$lib/services/fetchProdutos';
	import type { CondicaoOption } from '$lib/types/CargaInicio';

	let condicaoOptions: CondicaoOption[] = [];

	// Utilizando o sinal $ para obter os dados da store condicoesStore
	$: condicaoOptions = $condicoesStore.map((condicao) => ({
		label: condicao.E4_DESCRI.trim(),
		value: condicao.E4_CODIGO.trim()
	}));

	// Função para tratar a seleção de Tipo de NF e salvar na store
	async function handleTipoDeNFSelect(event: CustomEvent) {
		const selectedValue = event.detail;
		console.log('Tipo de NF selecionado:', selectedValue);

		// Access the 'label' and 'value' from 'selected'
		const { label } = selectedValue.selected;

		selectedTipoNF.set(label);
		fetchProdutos();
	}
	// Função para tratar a seleção de Prioridade e salvar na store
	function handlePrioridadeSelect(event: CustomEvent) {
		const selectedValue = event.detail;
		selectedPriority.set(selectedValue);
		console.log('Prioridade selecionada:', selectedValue);
	}

	// Função para tratar a seleção de Condição e salvar na store
	function handleCondicaoSelect(event: CustomEvent) {
		const selectedValue = event.detail;
		selectedCondicao.set(selectedValue);
		console.log('Condição selecionada:', selectedValue);
	}
</script>

<!-- Estrutura do componente com os Comboboxes -->
<div id="Inputs" class="card w-[35vw] p-5 gap-5 bg-base-100 justify-evenly">
	<div class="flex justify-center">
		<span class="font-bold text-lg">Adicionar Informações sobre a NF</span>
	</div>

	<div class="flex w-full justify-evenly gap-5">
		<!-- Combobox para Tipo de NF -->
		<div class="w-full">
			<Combobox
				options={tiposNFOptions}
				placeholder="Tipo de NF"
				on:select={handleTipoDeNFSelect}
				buttonClass="bg-base-300"
			/>
		</div>
	</div>

	<div class="flex w-full justify-evenly gap-5">
		<!-- Combobox para Prioridade -->
		<div class="w-full">
			<Combobox
				options={prioridadeOptions}
				placeholder="Prioridade"
				on:select={handlePrioridadeSelect}
				buttonClass="bg-base-300"
			/>
		</div>

		<!-- Combobox para Condição de Pagamento -->
		<div class="w-full">
			<Combobox
				options={condicaoOptions}
				placeholder="Forma de Pagamento"
				on:select={handleCondicaoSelect}
				buttonClass="bg-base-300"
			/>
		</div>
	</div>
</div>
