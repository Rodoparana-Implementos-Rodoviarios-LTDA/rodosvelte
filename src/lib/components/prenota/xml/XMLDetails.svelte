<script lang="ts">
	import { xmlDataStore } from '$lib/stores/xmlStore'; // Importar a store de dados XML
	import { fetchState } from '$lib/services/conexaoNFE'; // Importar a store de estado
	import type { ConexaoNFE } from '$lib/types/ConexaoNFE'; // Importar o tipo correto

	let detalhesXML: ConexaoNFE | null = null; // Dados da store de XML

	// Usa o estado da store fetchState
	$: estadoAtual = $fetchState;

	// Usa o sinal $ para obter a reatividade automática da store xmlDataStore
	$: detalhesXML = $xmlDataStore;
</script>

<!-- Renderização condicional baseada no estado -->
{#if estadoAtual === 'inicial'}
	<div class="h-full flex items-center justify-center">
		<p class="text-center text-info text-lg">
			Por favor, insira uma chave XML para buscar os detalhes.
		</p>
	</div>{:else if estadoAtual === 'carregando'}
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
{:else if estadoAtual === 'erro'}
	<p class="text-red-500">Erro ao buscar os dados da XML. Tente novamente.</p>
{:else if estadoAtual === 'sucesso' && detalhesXML}
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
{/if}
