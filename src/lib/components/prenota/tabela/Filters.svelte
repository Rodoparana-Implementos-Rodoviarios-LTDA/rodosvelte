<script lang="ts">
	import DateRangePicker from '$lib/components/ui/DatePicker.svelte';
	import { createEventDispatcher } from 'svelte';

	// Estados para os filtros
	export let selectedFilial = '';
	export let selectedNF = '';
	export let selectedStatus = '';
	export let selectedFornecedor = '';
	export let selectedTipo = '';
	export let selectedPrioridade = '';
	export let selectedUsuario = '';
	export let selectedObs = '';

	// Estados para os filtros de intervalo de datas
	export let selectedInclusionRange = '';
	export let selectedEmissionRange = '';
	export let selectedDueRange = '';

	const dispatch = createEventDispatcher();

	// Função para aplicar os filtros
	function applyFilters() {
		dispatch('applyFilters', {
			selectedFilial,
			selectedNF,
			selectedStatus,
			selectedFornecedor,
			selectedTipo,
			selectedPrioridade,
			selectedUsuario,
			selectedObs,
			selectedInclusionRange,
			selectedEmissionRange,
			selectedDueRange
		});
	}

	// Função para limpar os filtros
	function resetFilters() {
		selectedFilial = '';
		selectedNF = '';
		selectedStatus = '';
		selectedFornecedor = '';
		selectedTipo = '';
		selectedPrioridade = '';
		selectedUsuario = '';
		selectedObs = '';
		selectedInclusionRange = '';
		selectedEmissionRange = '';
		selectedDueRange = '';
		applyFilters(); // Aplica os filtros após limpar
	}
</script>

<!-- Menu de Filtros dentro do Drawer -->
<div
	class="menu p-4 w-80 bg-base-100 text-base-content h-screen flex flex-col justify-between"
	role="form"
	aria-labelledby="filters-heading"
>
	<div>
		<h2 id="filters-heading" class="text-lg font-bold mb-4 ">Filtros</h2>
		<!-- Filtro de Filial -->
		<div class="form-control mb-4">
			<label class="label" for="filial-input"><span class="label-text">Filial</span></label>
			<input
				id="filial-input"
				class="input input-sm input-bordered"
				type="text"
				bind:value={selectedFilial}
				placeholder="Filial"
				aria-label="Filtrar por Filial"
			/>
		</div>

		<!-- Filtro de NF -->
		<div class="form-control mb-4">
			<label class="label" for="nf-input"><span class="label-text">NF</span></label>
			<input
				id="nf-input"
				class="input input-sm input-bordered"
				type="text"
				bind:value={selectedNF}
				placeholder="Nota Fiscal"
				aria-label="Filtrar por Nota Fiscal"
			/>
		</div>

		<!-- Filtro de Status -->
		<div class="form-control mb-4">
			<label class="label" for="status-select"><span class="label-text">Status</span></label>
			<select
				id="status-select"
				class="select select-sm select-bordered"
				bind:value={selectedStatus}
				aria-label="Filtrar por Status"
			>
				<option value="">Todos</option>
				<option value="Classificado">Classificado</option>
				<option value="Revisar">Revisar</option>
				<option value="Pendente">Pendente</option>
			</select>
		</div>

		<!-- Filtro de Fornecedor -->
		<div class="form-control mb-4">
			<label class="label" for="fornecedor-input"><span class="label-text">Fornecedor</span></label>
			<input
				id="fornecedor-input"
				class="input input-sm input-bordered"
				type="text"
				bind:value={selectedFornecedor}
				placeholder="Fornecedor"
				aria-label="Filtrar por Fornecedor"
			/>
		</div>

		<!-- Filtro de Tipo -->
		<div class="form-control mb-4">
			<label class="label" for="tipo-select"><span class="label-text">Tipo</span></label>
			<select
				id="tipo-select"
				class="select select-sm select-bordered"
				bind:value={selectedTipo}
				aria-label="Filtrar por Tipo"
			>
				<option value="">Todos</option>
				<option value="Despesa/Imobilizado">Despesa/Imobilizado</option>
				<option value="Revenda">Revenda</option>
				<option value="Matéria Prima">Matéria Prima</option>
				<option value="Collection">Collection</option>
			</select>
		</div>

		<!-- Filtro de Prioridade -->
		<div class="form-control mb-4">
			<label class="label" for="prioridade-select"><span class="label-text">Prioridade</span></label
			>
			<select
				id="prioridade-select"
				class="select select-sm select-bordered"
				bind:value={selectedPrioridade}
				aria-label="Filtrar por Prioridade"
			>
				<option value="">Todas</option>
				<option value="Alta">Alta</option>
				<option value="Média">Média</option>
				<option value="Baixa">Baixa</option>
			</select>
		</div>
	</div>

	<!-- Botões de Aplicar e Limpar Filtros -->
	<div class="flex justify-between">
		<button class="btn btn-sm btn-primary" on:click={applyFilters} aria-label="Aplicar Filtros"
			>Aplicar Filtros</button
		>
		<button class="btn btn-sm btn-outline" on:click={resetFilters} aria-label="Limpar Filtros"
			>Limpar Filtros</button
		>
	</div>
</div>
