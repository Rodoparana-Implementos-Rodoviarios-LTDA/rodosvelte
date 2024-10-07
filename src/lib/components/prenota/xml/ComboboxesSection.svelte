<script lang="ts">
	import Combobox from '$lib/components/ui/Combobox.svelte';
	import { onMount } from 'svelte';
	import { tiposNFOptions, prioridadeOptions, condicoesStore } from '$lib/stores/xmlStore';
	import type { CondicaoOption } from '$lib/types/CargaInicio'; // Usando o tipo CondicaoOption

	// Inicializar as opções como um array vazio do tipo CondicaoOption[]
	let condicaoOptions: CondicaoOption[] = [];

	// Utilizando o sinal $ para obter os dados da store condicoesStore
	$: condicaoOptions = $condicoesStore.map((condicao) => ({
		label: condicao.E4_DESCRI.trim(),
		value: condicao.E4_CODIGO.trim()
	}));

	// Função para tratar a seleção de Tipo de NF
	function handleTipoDeNFSelect(event: CustomEvent) {
		console.log('Tipo de NF selecionado:', event.detail);
	}

	// Função para tratar a seleção de Prioridade
	function handlePrioridadeSelect(event: CustomEvent) {
		console.log('Prioridade selecionada:', event.detail);
	}

	// Função para tratar a seleção de Condição (Forma de Pagamento)
	function handleCondicaoSelect(event: CustomEvent) {
		console.log('Condição selecionada:', event.detail);
	}
	// Busca as opções de Condições ao montar o componente
	onMount(async () => {
		// Se necessário, você pode atualizar as stores aqui ao carregar o componente
	});
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
