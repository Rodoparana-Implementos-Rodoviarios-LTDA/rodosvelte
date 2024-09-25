<!-- ActionButton.svelte -->
<script lang="ts">
	import { fly } from 'svelte/transition';
	import { circOut } from 'svelte/easing';
	import {
		IconCopyPlus,
		IconClipboard,
		IconUser,
		IconCar
	} from '@tabler/icons-svelte';

	// Importando o componente ProductItens
	import ProductItens from './ProductItens.svelte';
	
	// Recebendo as informações via props
	export let documentoCompleto: string; 
	export let clienteCompleto: string;
	export let filial: string; 
	export let observacao: string | null = '';
	export let usuario: string = '000001';
	export let origem: string = 'E';

	// Variáveis de estado
	let responsible = '';
	let plate = '';
	let retiradaPor = 'C';
	let isLoading = false;
	let errorMessage = '';
	let drawerVisible = false;

	// Variável para armazenar os itens selecionados
	let itensSelecionados = [];

	// Abrir o drawer
	function openDrawer() {
		drawerVisible = true;
	}

	// Função para enviar a solicitação via POST
	async function handleConfirm() {
		console.log('handleConfirm - Início da função');

		// Validar campos obrigatórios
		if (!responsible || !plate) {
			errorMessage = 'Preencha todos os campos obrigatórios.';
			console.error('Erro: Campos obrigatórios faltando - responsible:', responsible, 'plate:', plate);
			return;
		}

		// Construir o corpo da solicitação
		const body = {
			filial: filial || '',
			documento: documentoCompleto,
			cliente: clienteCompleto,
			itens: itensSelecionados,
			retiradopor: retiradaPor || '',
			responsavel: responsible || '',
			placa: plate || '',
			observacoes: observacao || '',
			usuario: usuario || '',
			origem: origem || ''
		};

		console.log('Valores dos campos do body:');
		console.log(JSON.stringify(body, null, 2));

		try {
			isLoading = true;
			const response = await fetch('http://protheus-vm:9010/rest/MovPortaria/IncluirItem', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(body)
			});

			console.log('Status da resposta do POST:', response.status);

			const responseData = await response.json();
			console.log('Dados recebidos da API após o POST:', responseData);

			if (!response.ok || !responseData.success) {
				console.error('Erro na resposta da API após o POST:', responseData);
				errorMessage = `Erro: ${responseData.mensagem || 'Erro ao enviar a solicitação.'}`;
				return;
			}

			console.log('Solicitação enviada com sucesso', responseData);
			errorMessage = '';
			drawerVisible = false; // Fecha o drawer após o envio
		} catch (error) {
			console.error('Erro na requisição:', error);
			errorMessage = 'Erro ao enviar a solicitação.';
		} finally {
			isLoading = false;
		}
	}

	// Função para fechar o drawer ao clicar fora dele
	function closeDrawer(event: Event) {
		const target = event.target as HTMLElement;
		if (!target.closest('.drawer-content')) {
			drawerVisible = false;
		}
	}
</script>

<!-- Botão que abre o Drawer -->
<button class="btn btn-primary" on:click={openDrawer}>
	<IconCopyPlus />
</button>

<!-- Drawer -->
{#if drawerVisible}
	<!-- Overlay que fecha o Drawer ao clicar fora -->
	<div class="fixed inset-0 bg-black bg-opacity-50 z-40" on:click={closeDrawer}></div>

	<!-- Conteúdo do Drawer com animação fly -->
	<div class="drawer drawer-content" transition:fly={{ x: 400, easing: circOut }}>
		<div class="w-full h-screen bg-base-100 text-info p-4">
			<h2 class="text-xl font-bold mb-4 text-neutral-content">Seleção de Produtos</h2>

			<!-- Componente ProductItens que cuida de exibir a lista de produtos -->
			<ProductItens {documentoCompleto} {clienteCompleto} {filial} bind:itensSelecionados />

			<!-- Inputs para o nome do responsável e a placa -->
			<div class="mb-4 space-y-2">
				<div class="relative">
					<IconUser class="absolute left-3 top-3 text-base-content" />
					<input
						type="text"
						placeholder="Nome do responsável"
						bind:value={responsible}
						class="input input-bordered w-full pl-10"
					/>
				</div>
				<div class="relative">
					<IconCar class="absolute left-3 top-3 text-base-content" />
					<input
						type="text"
						placeholder="Placa do carro"
						bind:value={plate}
						class="input input-bordered w-full pl-10"
					/>
				</div>
			</div>

			<div class="mb-4">
				<textarea
					class="textarea textarea-bordered"
					placeholder="Observações"
					bind:value={observacao}
				></textarea>
			</div>

			<!-- Botões de Ação -->
			<div class="flex justify-between mt-6">
				<button class="btn btn-primary" on:click={handleConfirm} disabled={isLoading}>
					{#if isLoading}
						Enviando...
					{:else}
						Confirmar
					{/if}
				</button>
				<button class="btn btn-outline" on:click={() => (drawerVisible = false)}> Cancelar </button>
			</div>

			{#if errorMessage}
				<p class="text-error mt-4">{errorMessage}</p>
			{/if}
		</div>
	</div>
{/if}

<style>
	.drawer {
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
