<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	// Cria o dispatcher para enviar eventos para o componente pai
	const dispatch = createEventDispatcher();

	export let placeholder: string;
	export let searchTerm: string;
	export let disabled: boolean;

	// Função para lidar com a mudança no input
	const handleInputChange = (e) => {
		const inputValue = e.target.value;
		// Dispara o evento 'inputChange' com o valor atualizado
		dispatch('inputChange', inputValue);
	};
</script>

<!-- Componente de input -->
<input
	type="text"
	class="bg-transparent  flex-1 focus:outline-none"
	{placeholder}
	bind:value={searchTerm}
	on:input={handleInputChange}
	aria-label={placeholder}
	{disabled}
/>
