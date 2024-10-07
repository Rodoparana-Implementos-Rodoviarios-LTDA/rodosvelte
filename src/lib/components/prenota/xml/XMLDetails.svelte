<script lang="ts">
	import { xmlDataStore } from '$lib/stores/xmlStore'; // Importar a store de dados XML
	import type { DetalhesXML } from '$lib/types';

	export let loadingDetalhes: boolean = false; // Prop para controle de loading
	let detalhesXML: DetalhesXML | null = null; // Dados da store de XML
	let errorDetalhes: string | null = null;

	// Usando o sinal `$` para reatividade automática da store
	$: {
		if ($xmlDataStore) {
			detalhesXML = $xmlDataStore;
			errorDetalhes = null;
		} else if (!loadingDetalhes && !detalhesXML) {
			errorDetalhes = 'Nenhum dado encontrado.';
		}
	}
</script>

{#if loadingDetalhes}
	<div class="flex w-full flex-col gap-4">
		<div class="flex items-center gap-4">
			<div class="skeleton h-16 w-16 shrink-0 rounded-full"></div>
			<div class="flex flex-col gap-4 w-full">
				<div class="skeleton h-4 w-full"></div>
				<div class="skeleton h-4 w-full"></div>
			</div>
		</div>
		<div class="skeleton h-32 w-full"></div>
	</div>
{:else if errorDetalhes}
	<p class="text-red-500">Erro: {errorDetalhes}</p>
{:else if detalhesXML}
	<!-- Renderiza os detalhes da XML -->
	<div class="flex justify-between">
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">FORNECEDOR:</span>
			<div class="flex flex-col">
				<span class="text-sm opacity-70">{detalhesXML.nomeEmitente}</span>
				<span class="text-xs opacity-70"> {detalhesXML.cnpjEmitente}</span>
			</div>
		</div>
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">NF:</span>
			<span class="text-sm opacity-70 truncate">{detalhesXML.numero} - {detalhesXML.serie}</span>
		</div>
	</div>
	<div class="flex justify-between">
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">FILIAL:</span>
			<div class="flex flex-col">
				<span class="text-sm opacity-70"
					>{detalhesXML.filialName || detalhesXML.nomeDestinatario}</span
				>
				<span class="text-xs opacity-70"> {detalhesXML.cnpjDestinatario}</span>
			</div>
		</div>
		<div class="flex gap-2 items-center justify-start">
			<span class="text-lg font-medium">DATA:</span>
			<span class="opacity-70">{detalhesXML.dataEmissao}</span>
		</div>
	</div>
	<div class="collapse collapse-arrow bg-base-300">
		<input type="checkbox" />
		<div class="collapse-title text-lg font-medium w-full flex justify-center">Detalhes da XML</div>
		<div class="collapse-content">
			<p>{detalhesXML.informacoesAdicionais}</p>
		</div>
	</div>
{:else}
	<p>Insira uma chave XML para buscar os detalhes.</p>
{/if}
