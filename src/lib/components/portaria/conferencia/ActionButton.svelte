<script lang="ts">
	import { fly } from 'svelte/transition';
	import { circOut } from 'svelte/easing';
	import { IconClockSearch } from '@tabler/icons-svelte';
	import Conferir from './Conferir.svelte';  // Corrige o caminho para Conferir.svelte
	import { onMount } from 'svelte';

	export let documento: string;
	export let produto: string;
	export let saldoConferencia: number;
	export let responsavel: string;
	export let filial: string;
	export let cliente: string;

	let drawerVisible = false;  // Controla a visibilidade do drawer

	// Função para fechar o Drawer ao clicar fora
	function closeDrawer(event: Event) {
		const target = event.target as HTMLElement;
		if (!target.closest('.drawer-content')) {
			drawerVisible = false;
		}
	}

	// Escuta o evento customizado 'closeDrawer' para fechar o drawer
	onMount(() => {
		function closeDrawerEvent() {
			drawerVisible = false;
		}

		document.addEventListener('closeDrawer', closeDrawerEvent);

		return () => {
			document.removeEventListener('closeDrawer', closeDrawerEvent);
		};
	});
</script>

<!-- Botão que abre o Drawer -->
<button class="btn btn-primary" on:click={() => (drawerVisible = true)}>
	<IconClockSearch />
	<span class="sr-only">Conferência</span>
</button>

<!-- Drawer -->
{#if drawerVisible}
	<!-- Overlay que fecha o Drawer ao clicar fora -->
	<div role="close" class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>

	<!-- Conteúdo do Drawer -->
	<div class="drawer drawer-content" transition:fly={{ x: 400, easing: circOut }}>
		<div class="w-full h-screen bg-base-100 text-info p-8">
			<Conferir
				{documento}
				{responsavel}
				{produto}
				{saldoConferencia}
				{filial}
				{cliente}
			/>
		</div>
	</div>
{/if}

<style>
	.drawer {
		background: #012836;
		position: fixed;
		top: 0;
		right: 0;
		height: 100vh;
		width: 400px;
		display: flex;
		justify-content: space-between;
		z-index: 50;
	}
</style>
