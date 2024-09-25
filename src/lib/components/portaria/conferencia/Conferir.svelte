<script lang="ts">
	import { dataFetching } from '$lib/services/dataFetching';
	import { IconClipboard, IconTagsFilled, IconUser } from '@tabler/icons-svelte';
	import { onMount } from 'svelte';

	export let documento: string; // Documento da linha clicada
	export let produto: string; // Produto da linha clicada
	export let saldoConferencia: number; // Saldo conferido pelo usuário
	export let responsavel: string; // "C" ou "R" para identificar Cliente ou Motorista

	let modalVisible = false;
	let isCheckboxChecked = false; // Controla o estado do checkbox
	let saldoMaximo: number | null = null; // Saldo máximo obtido da API

	let pageData: any[] = [];
	let endpoint: string = 'api/borracharia';
	let filters = { 'NF': documento, 'Produto': produto }; // Inicialmente, nenhum filtro aplicado
	let sortBy = 'Emissao'; // Coluna padrão para ordenação
	let sortOrder: 'asc' | 'desc' = 'desc'; // Ordem padrão
	let pageSize: 10;
	let isLoading = true;

	// Função para abrir o modal
	function openModal() {
		if (isCheckboxChecked) {
			modalVisible = true;
		}
	}
	async function loadPage(page: number = 1) {

		try {
			const result = await dataFetching(endpoint, sortBy, sortOrder, page, pageSize, filters);
			pageData = result.data;
			saldoMaximo = pageData[0].Saldo
		} catch (error) {
			console.error(`Erro ao carregar dados de ${endpoint}:`, error);
		} finally {
			isLoading = false;
		}
	}
	// Função para fechar o modal e drawer
	function closeModalAndDrawer() {
		modalVisible = false;
		const closeEvent = new CustomEvent('closeDrawer', { bubbles: true });
		document.dispatchEvent(closeEvent);
	}

	// Função para carregar o saldo quando o componente for montado
	onMount(() => {
		if (documento && produto) {
			loadPage(); // Carrega o saldo da API ao montar o componente
		} else {
			console.error('Documento ou Produto está ausente na montagem.');
		}
	});
</script>

<!-- Conteúdo principal -->
<div class="w-full h-full bg-base-100 text-base-content p-6 rounded-lg shadow-lg">
	<h2 class="text-2xl font-bold mb-6 text-primary">Conferência de Produto</h2>

	<!-- Informações do Documento -->
	<p class="mb-4">
		<span class="inline-flex items-center">
			<IconClipboard class="mr-2 text-primary" />
			<strong class="text-primary text-lg">Documento:</strong>
		</span>
		<br />
		<span class="text-base-content text-lg">{documento}</span>
	</p>

	<!-- Informações do Produto -->
	<p class="mb-4">
		<span class="inline-flex items-center">
			<IconTagsFilled class="mr-2 text-primary" />
			<strong class="text-primary text-lg">Produto:</strong>
		</span>
		<br />
		<span class="text-base-content text-lg">{produto}</span>
	</p>

	<!-- Retirado por -->
	<p class="mb-4">
		<span class="inline-flex items-center">
			<IconUser class="mr-2 text-primary" />
			<strong class="text-primary text-lg">Retirado por:</strong>
		</span>
		<br />
		<span class="text-base-content text-lg">{responsavel}</span>
	</p>

	<!-- Comparação entre saldo original e saldo conferido -->
	<div class="mb-6 p-4 bg-base-300 border-l-4 border-primary rounded-lg shadow-md">
		<div class="flex justify-between items-center mb-4">
			<div>
				<span class="text-lg font-semibold text-primary">Documentação</span>
				<p class="text-sm text-warning">Quantidade registrada</p>
			</div>

			<!-- Mostrando o saldo máximo -->
			<span class="text-2xl font-bold text-base-content">
				{#if saldoMaximo !== null}
					{saldoMaximo}
				{:else}
					Buscando saldo...
				{/if}
			</span>
		</div>

		<hr class="my-2" />
		<div class="flex justify-between items-center">
			<div>
				<span class="text-lg font-semibold text-primary">Apanhada</span>
				<p class="text-sm text-warning">Quantidade real</p>
			</div>
			<span class="text-2xl font-bold text-base-content">{saldoConferencia}</span>
		</div>
	</div>

	<!-- Checkbox de confirmação -->
	<div class="form-control mb-8">
		<label class="cursor-pointer flex items-center space-x-4">
			<input type="checkbox" class="checkbox checkbox-primary" bind:checked={isCheckboxChecked} />
			<span class="label-text text-primary text-lg">Confirmar conferência</span>
		</label>
	</div>

	<!-- Botão para abrir o modal, só habilitado se o checkbox estiver marcado -->
	<button class="btn btn-primary w-full" on:click={openModal} disabled={!isCheckboxChecked}>
		Confirmar Conferência
	</button>
</div>

<!-- Modal de confirmação -->
{#if modalVisible}
	<dialog class="modal modal-open" id="confirm-modal">
		<div class="modal-box">
			<h3 class="font-bold text-primary text-lg">Tem certeza?</h3>
			<p class="py-4 text-neutral-content">
				Você está prestes a confirmar a conferência. Deseja continuar?
			</p>
			<div class="modal-action">
				<button class="btn btn-primary-content" on:click={() => (modalVisible = false)}>
					Cancelar
				</button>
				<button class="btn btn-primary" on:click={closeModalAndDrawer}>Sim, Confirmar</button>
			</div>
		</div>
	</dialog>
{/if}
