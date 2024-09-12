<script lang="ts">
	import { IconHistory } from '@tabler/icons-svelte';
	import AddPneu from './AddPneu.svelte';

	// Recebe o saldo como prop do componente pai
	export let saldo: number = 0;
</script>

<!-- Botão que abre o drawer -->
<div class="drawer drawer-end">
	<input id="my-drawer-4" type="checkbox" class="drawer-toggle" />
	<div class="drawer-content">
		<!-- Conteúdo da página aqui -->
		<label for="my-drawer-4" class="drawer-button btn btn-primary">
			<IconHistory class="inline-block h-5 w-5" /> Histórico
		</label>
	</div>
	<div class="drawer-side z-[100]">
		<label for="my-drawer-4" aria-label="Fechar histórico" class="drawer-overlay"></label>
		<ul class="menu bg-base-200 text-base-content min-h-full w-[25vw] p-4">
			<!-- Passa o saldo como prop para AddPneu -->
			<AddPneu saldo={saldo} />
		</ul>
	</div>
</div>
