<script lang="ts">
	import { HoverCard } from 'radix-svelte';

	// Recebe os valores como strings
	export let valorIpi: string;
	export let valorIcms: string;
	export let aliqIpi: string;
	export let aliqIcms: string;

	// Calcula o total de imposto
	const imposto = (parseFloat(valorIcms) || 0) + (parseFloat(valorIpi) || 0);

	// Funções auxiliares para verificar se os valores são válidos
	const isValidValue = (value: string) => value && !isNaN(parseFloat(value));

	// Função para formatar o valor com duas casas decimais e remover zeros extras
	const formatValue = (value: string) => parseFloat(value).toFixed(2); // fixa em duas casas decimais
</script>

<!-- Radix-Svelte HoverCard -->
<div class="contents">
	<HoverCard.Root openDelay={30} closeDelay={30}>
		<HoverCard.Trigger asChild>
			<!-- Trigger que exibe o valor total dos impostos -->
			<span>
				{imposto > 0 ? imposto.toFixed(2) : '--'} <!-- Exibe o valor total formatado, ou '--' se for 0 -->
			</span>
		</HoverCard.Trigger>

		<!-- Conteúdo do HoverCard exibindo as informações adicionais -->
		<HoverCard.Portal>
			<HoverCard.Content
				sideOffset={5}
				align={'end'}
				class="w-64 p-4 bg-primary-content text-primary text-sm shadow rounded"
			>
				<HoverCard.Arrow width={15} height={10} class="fill-primary-content" />
				<!-- Conteúdo do hover com as informações detalhadas dos impostos -->
				<div class="flex flex-col gap-2 w-full">
					{#if isValidValue(valorIpi)}
						<p class="flex justify-between truncate">
							Valor do IPI:
							<!-- Remove os zeros extras após a vírgula e exibe a alíquota se disponível -->
							<span class="truncate  flex">{formatValue(valorIpi)} {isValidValue(aliqIpi) ? `(${aliqIpi}%)` : ''}</span>
						</p>
					{/if}

					{#if isValidValue(valorIcms)}
						<p class="flex justify-between truncate">
							Valor do ICMS:
							<!-- Remove os zeros extras após a vírgula e exibe a alíquota se disponível -->
							<span class="truncate flex">{formatValue(valorIcms)} {isValidValue(aliqIcms) ? `(${aliqIcms}%)` : ''}</span>
						</p>
					{/if}
				</div>
			</HoverCard.Content>
		</HoverCard.Portal>
	</HoverCard.Root>
</div>
