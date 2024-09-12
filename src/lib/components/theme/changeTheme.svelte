<script lang="ts">
	import { onMount } from 'svelte';
	import { IconBrightnessHalf } from '@tabler/icons-svelte';
	import { lightThemes, darkThemes } from './themes';

	let selectedTheme: string = 'light'; // Tema padrão
	let isLightMode: boolean = true; // Controla se está no modo claro ou escuro

	// Função para alterar o tema e salvar no localStorage
	function changeTheme(theme: string): void {
		selectedTheme = theme;
		document.documentElement.setAttribute('data-theme', theme);
		localStorage.setItem('selectedTheme', theme); // Armazena o tema no localStorage
	}

	// Alternar entre claro e escuro
	function toggleTheme(): void {
		isLightMode = !isLightMode;
		const theme = isLightMode ? 'light' : 'dark';
		changeTheme(theme); // Altera o tema dinamicamente
	}

	// Carrega o tema salvo no localStorage ao montar o componente
	onMount(() => {
		const savedTheme = localStorage.getItem('selectedTheme');
		if (savedTheme) {
			selectedTheme = savedTheme;
			isLightMode = savedTheme === 'light'; // Ajusta o checkbox
			document.documentElement.setAttribute('data-theme', savedTheme);
		}
	});
</script>

<!-- Estrutura do drawer -->
<div class="drawer drawer-end">
	<input id="themeDrawer" type="checkbox" class="drawer-toggle" />

	<div class="drawer-content">
		<!-- Botão para abrir o drawer -->
		<label
			for="themeDrawer"
			class="btn btn-primary fixed top-5 right-5 z-0"
			aria-label="Alterar tema"
		>
			<IconBrightnessHalf class="inline-block h-5 w-5" />
		</label>
	</div>

	<!-- Conteúdo do drawer -->
	<div class="drawer-side z-[100]">
		<label for="themeDrawer" class="drawer-overlay" aria-label="Fechar tema"></label>

		<ul class="menu bg-base-200 text-base-content min-h-full w-46 p-4 space-y-4 overflow-y-auto">
			<!-- Alternador de modo claro/escuro -->
			<div class="flex items-center gap-2 p-4">
				<!-- Ícone de Sol (modo claro) -->
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="20"
					height="20"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				>
					<circle cx="12" cy="12" r="5" />
					<path
						d="M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4"
					/>
				</svg>

				<!-- DaisyUI Toggle -->
        <input type="checkbox" class="toggle" checked="checked" on:change={toggleTheme}/>

				<!-- Ícone de Lua (modo escuro) -->
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="20"
					height="20"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				>
					<path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"></path>
				</svg>
			</div>
			<!-- Renderiza temas claros ou escuros com base na seleção -->
			{#if isLightMode}
				<!-- Light Themes -->
				{#each lightThemes as theme}
					<li class="pb-2">
						<button
							class={`btn text-sm flex w-full ${theme.rounded} ${theme.bg}`}
							on:click={() => changeTheme(theme.name)}
						>
							<span class={`${theme.title} font-medium`}>{theme.name}</span>
							<div class="flex space-x-1">
								{#each theme.colors as color}
									<div class={`rounded-full h-5 w-2 ${color}`}></div>
								{/each}
							</div>
						</button>
					</li>
				{/each}
			{:else}
				<!-- Dark Themes -->
				{#each darkThemes as theme}
					<li class="pb-2">
						<button
							class={`btn text-sm flex justify-between w-full ${theme.rounded} ${theme.bg}`}
							on:click={() => changeTheme(theme.name)}
						>
							<span class={`font-medium ${theme.title}`}>{theme.name}</span>
							<div class="flex space-x-1">
								{#each theme.colors as color}
									<div class={`rounded-full h-5 w-2 ${color}`}></div>
								{/each}
							</div>
						</button>
					</li>
				{/each}
			{/if}
		</ul>
	</div>
</div>
