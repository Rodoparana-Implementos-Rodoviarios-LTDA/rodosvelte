<script lang="ts">
	import { IconHistory } from '@tabler/icons-svelte';
	import Timeline from './Timeline.svelte';
	import SolicitaRevisao from './SolicitaRevisao.svelte';
	import Produtos from './Produtos.svelte';
	let tabs = ['Historico', 'Solicitar Revisão', 'Produtos'];
	let activeTab = 0;

	export let status: string;
	export let rec: string;

	// Função que verifica se a aba é ativa
	function isActive(index: number): boolean {
		return index === activeTab;
	}

	// Função para alterar a aba ativa
	function setActiveTab(index: number) {
		activeTab = index;
	}
</script>

<div class="drawer drawer-end">
	<input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content">
		<label for="my-drawer-4" class="drawer-button btn btn-ghost w-fit">
			<IconHistory class="inline-block h-7 w-7" />
		</label>
	</div>
	<div class="drawer-side z-[100]">
		<label for="my-drawer-4" aria-label="Fechar histórico" class="drawer-overlay"></label>
		<div class="w-[20vw] h-screen bg-base-100 text-info p-8">
			<div class="tabs tabs-boxed">
				{#each tabs as tab, index}
					<button
						class="tab tab-lifted"
						class:tab-active={activeTab == index}
						on:click={() => (activeTab = index)}>{tab}</button
					>
				{/each}
			</div>
			<div class="w-full h-full pt-5" class:hidden={activeTab != 0}>
				<Timeline {rec} {status} /> <!-- Passa o rec e status corretos para Timeline -->
			</div>
			<div class="w-full h-full pt-5" class:hidden={activeTab != 1}>
				<SolicitaRevisao />
			</div>
			<div class="w-full h-full pt-5" class:hidden={activeTab != 2}>
				<Produtos />
			</div>
		</div>
	</div>
</div>
