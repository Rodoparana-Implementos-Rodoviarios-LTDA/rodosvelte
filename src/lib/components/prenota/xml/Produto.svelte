<script lang="ts">
	import TableCombobox from './XmlCombobox.svelte';
	import { produtosStore, fetchState } from '$lib/stores/xmlStore';

	export let index: number;
	export let origem: string; // Should be a string
	export let ncmsh: string; // Should be a string

	let options = [];
	let selectedOption = null;
	let isLoading = true;

	// Reactive statements
	$: isLoading = $fetchState === 'carregando';
	$: options = ($produtosStore || []).map(({ label, value, B1_ORIGEM, B1_POSIPI }) => ({
		label,
		value,
		origem: String(B1_ORIGEM || '').trim(),
		ncm: String(B1_POSIPI || '').trim()
	}));

	function handleSelect(event) {
		selectedOption = event.detail.selected;
		console.log(`Produto no índice ${index} selecionado:`, selectedOption);
	}
</script>

<!-- Combobox Component -->
{#if isLoading}
	<div class=""><IconRotateClockwise2 class="animate-spin" /></div>
{:else}
	<TableCombobox
		{options}
		placeholder="Buscar produto por código ou descrição..."
		disabled={isLoading}
		bind:selectedOption
		on:select={handleSelect}
		{origem}
		{ncmsh}
	/>
{/if}
