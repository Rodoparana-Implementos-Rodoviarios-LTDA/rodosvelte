// src/utils/comboboxUtils.ts

import type { ComboboxOption } from '$lib/types/Produtos'; // Corrija o caminho conforme necessário

export function toggleDropdown(
	isOpen: boolean,
	disabled: boolean,
	highlightedIndex: number,
	inputRef: HTMLInputElement | null
) {
	if (disabled) return { isOpen, highlightedIndex };

	isOpen = !isOpen;
	if (isOpen) {
		highlightedIndex = -1;
		if (inputRef) {
			inputRef.focus();
		}
	}

	return { isOpen, highlightedIndex };
}

export function selectOption(
	option: ComboboxOption,
	multiple: boolean,
	selectedOptions: ComboboxOption[] | ComboboxOption | null,
	dispatch: (event: string, detail?: any) => void
) {
	if (multiple) {
		let updatedOptions: ComboboxOption[] = Array.isArray(selectedOptions)
			? [...selectedOptions]
			: [];
		const alreadySelected = updatedOptions.find((selected) => selected.value === option.value);

		if (alreadySelected) {
			updatedOptions = updatedOptions.filter((selected) => selected.value !== option.value);
		} else {
			updatedOptions = [...updatedOptions, option];
		}
		selectedOptions = updatedOptions;
	} else {
		selectedOptions = option;
	}

	dispatch('select', { selected: selectedOptions });

	return selectedOptions;
}

export function isSelected(
	option: ComboboxOption,
	multiple: boolean,
	selectedOptions: ComboboxOption[] | ComboboxOption | null
) {
	if (multiple) {
		return (
			Array.isArray(selectedOptions) &&
			selectedOptions.some((selected) => selected.value === option.value)
		);
	} else {
		return selectedOptions && (selectedOptions as ComboboxOption).value === option.value;
	}
}

export function handleKeyDown(
	e: KeyboardEvent,
	highlightedIndex: number,
	filteredOptions: ComboboxOption[],
	selectOption: (option: ComboboxOption) => void
) {
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
		return false;
	}

	return highlightedIndex;
}

export function handleClickOutside(event: MouseEvent, inputRef: HTMLInputElement | null) {
	if (inputRef && !inputRef.contains(event.target as Node)) {
		return false;
	}
	return true;
}
