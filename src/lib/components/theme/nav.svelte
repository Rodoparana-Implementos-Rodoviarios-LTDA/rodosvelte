<script lang="ts">
	import Logo from '$lib/Components/theme/Logo.svelte';
	import { goto } from '$app/navigation';
	import { menuItems } from '$lib/stores/menuItems';
	import { clearIndexedDB } from '$lib/services/idb';

	let username = 'Usuário';
	let isLoading = true;

	// Carregar os dados do usuário do sessionStorage
	async function loadUserData() {
		if (typeof window !== 'undefined') {
			try {
				const storedUsername = sessionStorage.getItem('username');
				if (storedUsername) {
					username = storedUsername;
				}
			} catch (error) {
				console.error('Erro ao buscar o nome de usuário do sessionStorage:', error);
			} finally {
				isLoading = false;
			}
		}
	}

	// Função de logout que limpa o IndexedDB e redireciona para a tela de login
	async function handleLogout() {
		try {
			sessionStorage.removeItem('username');
			sessionStorage.removeItem('token');
			await clearIndexedDB();
			goto('/');
			console.log('Usuário deslogado e redirecionado para /');
		} catch (error) {
			console.error('Erro ao fazer logout:', error);
		}
	}

	loadUserData();
</script>

<!-- Drawer Layout do DaisyUI -->
<div class="drawer">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" />

	<!-- Conteúdo Principal -->
	<div class="drawer-content flex flex-col">
		<label for="my-drawer" class="btn btn-square btn-outline fixed top-5 left-5 z-10">
			<lord-icon
				src="https://cdn.lordicon.com/hqymfzvj.json"
				trigger="hover"
				state="hover-file-2"
				style="width: 24px; height: 24px; filter: brightness(0) saturate(100%) invert(57%) sepia(8%) saturate(7470%) hue-rotate(175deg) brightness(90%) contrast(87%);"
			>
			</lord-icon>
		</label>
		<slot />
	</div>

	<!-- Drawer Lateral -->
	<div class="drawer-side h-screen z-20">
		<label for="my-drawer" class="drawer-overlay"></label>
		<ul class="menu p-4 w-80 bg-base-200 text-base-content h-screen flex flex-col overflow-y-auto">
			<div class="pb-14">
				<Logo size="w-14 h-14" textSize="text-4xl" />
			</div>
			<!-- Accordion para as seções -->
			{#each menuItems as section, index}
				{#if section.items && section.items.length > 0}
					<!-- Renderiza com colapso se tiver sub-itens -->
					<div class="collapse bg-base-300 mb-6 hover:bg-base-100">
						<input type="radio" name="menu-accordion" id="section-{index}" class="peer" />
						<label
							for="section-{index}"
							class="collapse-title text-lg font-medium flex items-center gap-3"
						>
							<svelte:component this={section.icon} class="h-6 w-6 text-blue-500" />
							{section.title}
						</label>
						<div class="collapse-content">
							{#each section.items as item}
								<li class="mb-2">
									<a
										href={item.link}
										class="flex items-center space-x-3 hover:bg-base-300 rounded-lg transition duration-200 ease-in-out text-lg"
										target={item.external ? '_blank' : '_self'}
										tabindex="0"
										role="menuitem"
										aria-label={item.name}
									>
										<lord-icon
											src={item.iconUrl}
											trigger="hover"
											state="hover-pinch"
											style="width: 35px; height: 35px; filter: brightness(0) saturate(100%) invert(57%) sepia(8%) saturate(7470%) hue-rotate(175deg) brightness(90%) contrast(87%);"
										>
										</lord-icon>

										<span>{item.name}</span>
									</a>
								</li>
							{/each}
						</div>
					</div>
				{:else}
					<!-- Renderiza como link direto se não houver sub-itens -->
					<div class="bg-base-300 rounded-xl  mb-6 hover:bg-base-100">
						<a
							href={section.link}
							class="collapse-title text-lg font-medium flex items-center gap-3"
						>
							<svelte:component this={section.icon} class="h-6 w-6 text-blue-500" />
							{section.title}
						</a>
					</div>
				{/if}
			{/each}

			<!-- Logout no fim da página -->
			<div class="mt-auto mb-4">
				<button
					class="btn btn-outline btn-error flex items-center justify-center w-full h-full relative group"
					on:click={handleLogout}
				>
					<span class="text-xl transition-opacity duration-300 ease-in-out group-hover:opacity-0"
						>{username}</span
					>
					<span
						class="w-full  text-xl flex items-center justify-evenly gap-3 absolute opacity-0 transition-opacity duration-400 ease-in-out group-hover:opacity-100"
					>
						Sair do RodoApp
						<lord-icon
							src="https://cdn.lordicon.com/nqtddedc.json"
							trigger="hover"
							colors="primary:#ffffff"
							style="width: 35px; height: 35px"
						>
						</lord-icon>
					</span>
				</button>
			</div>
		</ul>
	</div>
</div>
