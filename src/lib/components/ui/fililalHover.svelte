<script lang="ts">
	import { onMount } from 'svelte';
	import { getOptionsData } from '$lib/services/idb'; // Função para buscar dados no IndexedDB
	import { HoverCard } from 'radix-svelte';

	export let numeroFilial: string; // O número da filial exibido na tabela

	let filialInfo: string = 'Carregando...'; // Placeholder enquanto carrega os dados

	// Função para buscar a filial no IndexedDB
	async function fetchFilial() {
		try {
			// Busca os dados da tabela 'filiaisOptions' usando a chave 'filiais'
			const filiais = await getOptionsData('filiaisOptions', 'filiais');
			if (filiais && filiais.data) {
				const filialEncontrada = filiais.data.find(
					(filial: any) => filial.numero.trim() === numeroFilial.trim()
				);
				filialInfo = filialEncontrada ? filialEncontrada.filial : 'Filial não encontrada';
			} else {
				filialInfo = 'Nenhuma filial encontrada no banco de dados';
			}
		} catch (error) {
			console.error('Erro ao buscar filial:', error);
			filialInfo = 'Erro ao carregar dados da filial';
		}
	}

	// Chama a função ao montar o componente
	onMount(() => {
		fetchFilial();
	});
</script>

<!-- Componente Radix-Svelte HoverCard -->
<HoverCard.Root openDelay={30} closeDelay={30}>
	<HoverCard.Trigger asChild>
		<span>{numeroFilial.trim()}</span>
	</HoverCard.Trigger>

	<!-- Conteúdo do HoverCard exibindo o nome da filial -->
	<HoverCard.Portal>
		<HoverCard.Content
			sideOffset={5}
			align={'start'}
			class="p-4 bg-primary-content text-primary shadow rounded"
		>
			<HoverCard.Arrow width={15} height={10} class="fill-primary-content" />
			<p class="text-base font-medium">{filialInfo}</p>
		</HoverCard.Content>
	</HoverCard.Portal>
</HoverCard.Root>

<style>
	.contents {
		display: inline-block;
	}
</style>
