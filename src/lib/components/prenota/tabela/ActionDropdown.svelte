<script lang="ts">
	import { IconHistory } from '@tabler/icons-svelte';
	import Timeline from './Timeline.svelte';
	import revisaoFetch from '$lib/services/revisaoFetch'; // Função para buscar os eventos
	import type { Revisao } from '$lib/types/revisao';

	export let rec: string; // Recebe o 'rec' da linha da tabela
	let eventos: Revisao[] = []; // Lista de eventos

	// Função para carregar eventos da API
	async function loadEventos() {
		if (rec) {
			try {
				const response = await revisaoFetch(rec); // Busca eventos pelo rec
				eventos = response.historico; // Define eventos
			} catch (error) {
				console.error('Erro ao buscar eventos:', error);
			}
		}
	}
</script>

<!-- Botão que abre o drawer -->
<div class="drawer drawer-end">
	<input id="my-drawer-4" type="checkbox" class="drawer-toggle" on:change={loadEventos} />
	<div class="drawer-content">
		<!-- Conteúdo da página aqui -->
		<label for="my-drawer-4" class="drawer-button btn btn-outline btn-accent-content w-full">
			<IconHistory class="inline-block h-5 w-5" /> Histórico
		</label>
	</div>
	<div class="drawer-side z-[100]">
		<label for="my-drawer-4" aria-label="Fechar histórico" class="drawer-overlay"></label>
		<ul class="menu bg-base-200 text-base-content min-h-full w-[25vw] p-4">
			<Timeline {eventos} />
		</ul>
	</div>
</div>
