<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import { IconChevronDown } from '@tabler/icons-svelte';

	export let options: {
		label: string;
		value: string;
		campo1?: string;
		campo2?: string;
	}[] = [];

	export let value: string | string[] = ''; // Valor programático para seleção
	export let placeholder: string = 'Selecione uma opção';
	export let disabled: boolean = false;
	export let multiple: boolean = false; // Seleção múltipla
	export let additionalClass: string = '';
	export let buttonClass: string = '';
	export let tit1: string = '';
	export let tit2: string = '';

	const dispatch = createEventDispatcher();
	let isOpen = false;
	let selectedOptions = multiple ? [] : null; // Múltiplas seleções ou único
	let searchTerm = '';
	let highlightedIndex = -1;
	let inputRef: HTMLInputElement | null = null;

	// Reatividade para atualizar `selectedOptions` programaticamente
	$: if (value) {
		if (multiple && Array.isArray(value)) {
			// Para múltiplas seleções, preenche o array de opções selecionadas
			selectedOptions = options.filter((option) => value.includes(option.value));
		} else if (!multiple && typeof value === 'string') {
			// Para seleção única
			selectedOptions = options.find((option) => option.value === value) || null;
		}
	}

	// Filtrar opções com base no termo de busca
	$: filteredOptions = searchTerm
		? options.filter(
				(option) =>
					option.label.toLowerCase().includes(searchTerm.toLowerCase()) ||
					option.value.toLowerCase().includes(searchTerm.toLowerCase())
			)
		: options;

	// Abrir e fechar o dropdown
	const toggleDropdown = () => {
		if (disabled) return;
		isOpen = !isOpen;
		if (isOpen) {
			highlightedIndex = -1;
			if (inputRef) {
				inputRef.focus();
			}
		}
	};

	// Selecionar uma ou múltiplas opções
	const selectOption = (option) => {
		if (multiple) {
			const alreadySelected = selectedOptions.find((selected) => selected.value === option.value);
			if (alreadySelected) {
				selectedOptions = selectedOptions.filter((selected) => selected.value !== option.value);
			} else {
				selectedOptions = [...selectedOptions, option];
			}
			searchTerm = ''; // Limpar o campo de busca após selecionar
			inputRef?.focus(); // Focar no input novamente para permitir nova busca
		} else {
			selectedOptions = option;
			isOpen = false;
		}
		dispatch('select', { selected: selectedOptions });
	};

	// Verificar se a opção está selecionada
	const isSelected = (option) => {
		if (multiple) {
			return selectedOptions.some((selected) => selected.value === option.value);
		} else {
			return selectedOptions && selectedOptions.value === option.value;
		}
	};

	// Controlar a navegação com teclado (setas e Enter)
	const handleKeyDown = (e) => {
		if (e.key === 'ArrowDown') {
			e.preventDefault();
			highlightedIndex = (highlightedIndex + 1) % filteredOptions.length;
		} else if (e.key === 'ArrowUp') {
			e.preventDefault();
			highlightedIndex = (highlightedIndex - 1 + filteredOptions.length) % filteredOptions.length;
		} else if (e.key === 'Enter') {
			e.preventDefault();
			if (highlightedIndex >= 0 && highlightedIndex < filteredOptions.length) {
				selectOption(filteredOptions[highlightedIndex]);
			}
		} else if (e.key === 'Escape') {
			isOpen = false;
		}
	};

	// Fechar o dropdown ao clicar fora
	const handleClickOutside = (event) => {
		if (inputRef && !inputRef.contains(event.target)) {
			isOpen = false;
		}
	};

	onMount(() => {
		document.addEventListener('click', handleClickOutside);
		return () => {
			document.removeEventListener('click', handleClickOutside);
		};
	});
</script>

<!-- Estrutura do Combobox -->
<div class={`relative ${additionalClass}`}>
	<!-- Botão para abrir o dropdown -->
	<button
		type="button"
		class={`flex items-center justify-start border ${
			disabled ? 'bg-base-200 cursor-not-allowed' : 'bg-base-100'
		} border-base-300 rounded-md p-2 ${buttonClass}`}
		on:click={toggleDropdown}
		on:keydown={handleKeyDown}
		aria-haspopup="listbox"
		aria-expanded={isOpen}
	>
		<div class="flex flex-wrap gap-1 w-full">
			{#if multiple}
				<!-- Mostrar múltiplas opções selecionadas -->
				{#each selectedOptions as option}
					<span class="badge badge-neutral text-primary font-medium text-base">
						{option.label}
						<button
							class="ml-1 text-error"
							on:click|stopPropagation={() => selectOption(option)}
							aria-label={`Remover ${option.label}`}
						>
							&times;
						</button>
					</span>
				{/each}
				{#if selectedOptions.length === 0}
					<!-- Input de busca -->
					<input
						type="text"
						class="bg-transparent flex-1 focus:outline-none"
						{placeholder}
						bind:value={searchTerm}
						on:keydown={handleKeyDown}
						aria-label={placeholder}
						bind:this={inputRef}
						{disabled}
					/>
				{/if}
			{:else if selectedOptions}
				<!-- Mostrar única opção selecionada -->
				<span class="flex-1">{selectedOptions.label}</span>
			{:else}
				<!-- Input de busca -->
				<input
					type="text"
					class="bg-transparent flex-1 focus:outline-none w-full"
					{placeholder}
					bind:value={searchTerm}
					on:keydown={handleKeyDown}
					aria-label={placeholder}
					bind:this={inputRef}
					{disabled}
				/>
			{/if}
		</div>
		<IconChevronDown
			class={`transition-transform duration-200 ${isOpen ? 'transform rotate-180' : ''}`}
		/>
	</button>

	{#if isOpen}
		<div
			class="absolute z-10 w-full bg-base-100 border border-base-300 rounded-md mt-1 max-h-60 overflow-auto shadow-lg"
		>
			<ul role="listbox" aria-labelledby="combobox-label">
				{#if filteredOptions.length > 0}
					{#each filteredOptions as option, index}
						<li
							class={`px-3 py-2 cursor-pointer text-start flex flex-col justify-start ${
								highlightedIndex === index ? 'bg-base-200 text-primary' : ''
							} ${isSelected(option) ? 'font-bold' : ''}`}
							on:click={() => selectOption(option)}
							on:mouseenter={() => (highlightedIndex = index)}
							on:keydown={(e) => {
								if (e.key === 'Enter' || e.key === ' ') {
									selectOption(option);
								}
							}}
							role="option"
							aria-selected={isSelected(option)}
							tabindex="0"
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
	{/if}
</div>
