<script>
	import Combobox from '$lib/components/ui/Combobox.svelte';
	import { getOptionsData } from '$lib/services/idb';
	import { onMount } from 'svelte';

	let tipoDeNFOptions = [];
	let prioridadeOptions = [];
	let condicaoOptions = []; // Para o combobox de Forma de Pagamento

	// Função para tratar a seleção de Tipo de NF
	function handleTipoDeNFSelect(event) {
		console.log('Tipo de NF selecionado:', event.detail);
	}

	// Função para tratar a seleção de Filial
	function handleFilialSelect(event) {
		console.log('Filial selecionada:', event.detail);
	}

	// Função para tratar a seleção de Forma de Pagamento (Condição)
	function handleCondicaoSelect(event) {
		console.log('Condição selecionada:', event.detail);
	}

	// Função para tratar a seleção de Prioridade
	function handlePrioridadeSelect(event) {
		console.log('Prioridade selecionada:', event.detail);
	}

	// Busca os dados do IndexedDB ao montar o componente
	onMount(async () => {
		try {
			// Recupera as opções de Tipo de NF e Prioridade do IndexedDB
			tipoDeNFOptions = (await getOptionsData('tipoDeNFOptions')) || [];
			prioridadeOptions = (await getOptionsData('prioridadeOptions')) || [];

			// Recupera os dados de Condições (se necessário)
			const rawCondicaoOptions = (await getOptionsData('condicoesOptions')) || [];
			condicaoOptions = rawCondicaoOptions.map((condicao) => ({
				label: condicao.E4_DESCRI.trim(), // Descrição da Condição
				value: condicao.E4_CODIGO.trim() // Código da Condição
			}));
		} catch (error) {
			console.error('Erro ao buscar dados do IndexedDB:', error);
		}
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
				options={tipoDeNFOptions}
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
