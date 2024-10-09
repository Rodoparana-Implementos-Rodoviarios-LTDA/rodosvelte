<script lang="ts">
	import { createEventDispatcher, onMount } from 'svelte';
	import { IconChevronDown } from '@tabler/icons-svelte';
  
	export let options: { label: string; value: string }[] = [];
	export let placeholder: string = 'Selecione uma opção';
	export let disabled: boolean = false;
	export let multiple: boolean = false;
	export let additionalClass: string = '';
	export let buttonClass: string = '';
  
	const dispatch = createEventDispatcher();
  
	let isOpen = false;
	export let selectedOption = multiple ? [] : null;
	let searchTerm = '';
	let highlightedIndex = -1;
	let inputRef: HTMLInputElement | null = null;
  
	// Computed store for filtered options
	$: filteredOptions = searchTerm
	  ? options.filter((option) => option.label.toLowerCase().includes(searchTerm.toLowerCase()))
	  : options;
  
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
  
	const selectOption = (option) => {
	  if (multiple) {
		if (selectedOption.find((selected) => selected.value === option.value)) {
		  selectedOption = selectedOption.filter((selected) => selected.value !== option.value);
		} else {
		  selectedOption = [...selectedOption, option];
		}
	  } else {
		selectedOption = option;
		isOpen = false;
	  }
	  dispatch('select', { selected: selectedOption });
	};
  
	const isSelected = (option) => {
	  if (multiple) {
		return selectedOption.find((selected) => selected.value === option.value);
	  } else {
		return selectedOption && selectedOption.value === option.value;
	  }
	};
  
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
  
  <div class={`relative ${additionalClass}`}>
	<div
	  class={`flex items-center justify-start border ${
		disabled ? 'bg-base-200 cursor-not-allowed' : 'bg-base-100'
	  } border-base-300 rounded-md p-2 cursor-pointer ${buttonClass}`}
	  on:click={toggleDropdown}
	  aria-haspopup="listbox"
	  aria-expanded={isOpen}
	  role="combobox"
	>
	  <div class="flex flex-wrap gap-1 w-full">
		{#if multiple}
		  {#each selectedOption as option}
			<span class="badge badge-primary badge-outline">
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
		  {#if selectedOption.length === 0}
			<input
			  type="text"
			  class="bg-transparent flex-1 focus:outline-none"
			  {placeholder}
			  bind:value={searchTerm}
			  on:click|stopPropagation={() => (isOpen = true)}
			  on:keydown={handleKeyDown}
			  on:input={(e) => {
				searchTerm = e.target.value;
				isOpen = true;
			  }}
			  aria-label={placeholder}
			  bind:this={inputRef}
			  {disabled}
			/>
		  {/if}
		{:else if selectedOption}
		  <span class="flex-1">{selectedOption.label}</span>
		{:else}
		  <input
			type="text"
			class="bg-transparent flex-1 focus:outline-none"
			{placeholder}
			bind:value={searchTerm}
			on:click|stopPropagation={() => (isOpen = true)}
			on:keydown={handleKeyDown}
			on:input={(e) => {
			  searchTerm = e.target.value;
			  isOpen = true;
			}}
			aria-label={placeholder}
			bind:this={inputRef}
			{disabled}
		  />
		{/if}
	  </div>
	  <IconChevronDown
		class={`transition-transform duration-200 ${isOpen ? 'transform rotate-180' : ''}`}
	  />
	</div>
  
	{#if isOpen}
	  <div
		class="absolute z-10 w-full bg-base-100 border border-base-300 rounded-md mt-1 max-h-60 overflow-auto shadow-lg"
	  >
		<ul role="listbox" aria-labelledby="combobox-label">
		  {#if filteredOptions.length > 0}
			{#each filteredOptions as option, index}
			  <li
				class={`px-3 py-2 cursor-pointer text-start ${
				  highlightedIndex === index ? 'bg-base-200 text-primary' : ''
				} ${isSelected(option) ? 'font-bold' : ''}`}
				on:click={() => selectOption(option)}
				on:mouseenter={() => (highlightedIndex = index)}
				role="option"
				aria-selected={isSelected(option)}
			  >
				<span>{option.label}</span>
			  </li>
			{/each}
		  {:else}
			<li class="px-3 py-2 text-base-content/60">Nenhuma opção encontrada</li>
		  {/if}
		</ul>
	  </div>
	{/if}
  </div>
  
  <style>
	.badge button {
	  background: transparent;
	  border: none;
	  cursor: pointer;
	}
  </style>
  