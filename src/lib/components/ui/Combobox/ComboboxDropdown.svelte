<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let filteredOptions = [];
	export let highlightedIndex = -1;
	export let isSelected;
	export let tit1: string;
	export let tit2: string;

	const dispatch = createEventDispatcher();

	const selectOption = (option) => {
		dispatch('optionSelect', option);
	};

	const highlightIndexChange = (index) => {
		dispatch('highlightIndexChange', index);
	};
</script>

<div class="absolute z-10 w-full bg-base-100 border border-base-300 rounded-md mt-1 max-h-60 overflow-auto shadow-lg">
	<ul role="listbox" aria-labelledby="combobox-label">
		{#if filteredOptions.length > 0}
			{#each filteredOptions as option, index}
				<li
					class={`px-3 py-2 cursor-pointer text-start flex flex-col justify-start ${
						highlightedIndex === index ? 'bg-base-200 text-primary' : ''
					} ${isSelected(option) ? 'font-bold' : ''}`}
					on:click={() => selectOption(option)}
					on:mouseenter={() => highlightIndexChange(index)}
					role="option"
					aria-selected={isSelected(option)}
				>
					<!-- Mostra os campos dinâmicos -->
					{#if option.campo1 || option.campo2}
						<div class="flex justify-start items-center gap-2">
							{#if option.campo1}
								<span class="text-xs text-base-content/60">{tit1} {option.campo1}</span>
							{/if}
							{#if option.campo1 && option.campo2}
								<span class="text-xs text-base-content/60">|</span>
							{/if}
							{#if option.campo2}
								<span class="text-xs text-base-content/60">{tit2} {option.campo2}</span>
							{/if}
						</div>
					{/if}
					<span class="text-base font-bold">{option.label}</span>
				</li>
			{/each}
		{:else}
			<li class="px-3 py-2 text-base-content/60">Nenhuma opção encontrada</li>
		{/if}
	</ul>
</div>
