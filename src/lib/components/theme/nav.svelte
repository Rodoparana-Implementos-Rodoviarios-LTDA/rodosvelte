<script lang="ts">
	import {
		InfoCircle,
		Phone,
		Logout,
		Dashboard,
		List,
		Plus,
		Upload,
		Businessplan,
		ShieldCheck,
		Category
	} from 'tabler-icons-svelte'; // Ícones adicionais
	import Logo from './logo.svelte'; // Componente de logo
	import ThemeChanger from './changeTheme.svelte'; // Componente de troca de tema
	import { goto } from '$app/navigation';
	import { getData, deleteData } from '$lib/services/idb'; // Funções para manipular IndexedDB

	// Definindo a tipagem para os itens do menu
	interface MenuItem {
		name: string;
		icon: typeof Dashboard | typeof Phone; // Tipagem para os ícones
		link: string;
		external?: boolean; // Tornar opcional para itens que não precisam de "external"
	}

	interface MenuSection {
		title: string;
		items: MenuItem[];
	}

	// Definição do menu usando um objeto com tipagem
	const menuItems: MenuSection[] = [
		{
			title: 'Pre Notas',
			items: [
				{ name: 'Dashboard', icon: Dashboard, link: '/prenotas/dashboard' },
				{ name: 'Lista de Pre Notas', icon: List, link: '/prenotas/' },
				{ name: 'Incluir Manualmente', icon: Plus, link: '/prenotas/incluir' },
				{ name: 'Incluir XML', icon: Upload, link: '/prenotas/xml' }
			]
		},
		{
			title: 'Portaria',
			items: [
				{ name: 'Borracharia', icon: Dashboard, link: '/portaria' },
				{ name: 'Dashboard', icon: Businessplan, link: '/portaria/dashboard' },
				{ name: 'Conferência', icon: ShieldCheck, link: '/portaria/conferencia' }
			]
		},
		{
			title: 'Suporte e Intranet',
			items: [
				{ name: 'Suporte', icon: Phone, link: 'https://hesk.rodoparana.com.br', external: true },
				{
					name: 'Intranet',
					icon: InfoCircle,
					link: 'https://sites.google.com/site/baserodoparana/home?authuser=1',
					external: true
				}
			]
		}
	];

	let username = 'Usuário'; // Valor padrão caso não haja dados
	let isLoading = true; // Variável para mostrar carregamento

	// Função para obter o username do IndexedDB
	async function loadUserData() {
		if (typeof window !== 'undefined') {
			// Verifica se está no client-side
			try {
				const storedUsername = sessionStorage.getItem('username'); // Pega o username da sessionStorage
				if (storedUsername) {
					username = storedUsername; // Atualiza o username
				}
			} catch (error) {
				console.error('Erro ao buscar o nome de usuário do sessionStorage:', error);
			} finally {
				isLoading = false; // Indica que o carregamento terminou
			}
		}
	}

	// Função para logout
	async function handleLogout() {
		try {
			// Remove o token e o nome de usuário do IndexedDB
			await deleteData('username');
			await deleteData('authToken');

			// Redireciona para a página inicial ou página de login
			goto('/');
			console.log('Usuário deslogado e redirecionado para /');
		} catch (error) {
			console.error('Erro ao fazer logout:', error);
		}
	}

	// Carregar o nome do usuário ao montar o componente
	loadUserData();
</script>

<!-- Drawer Layout do DaisyUI -->
<div class="drawer">
	<input id="my-drawer" type="checkbox" class="drawer-toggle" />

	<!-- Conteúdo Principal -->
	<div class="drawer-content flex flex-col">
		<label for="my-drawer" class="btn btn-square btn-ghost fixed top-5 left-5 z-10">
			<Category class="h-6 w-6" />
		</label>
		<slot />
	</div>

	<!-- Drawer Lateral -->
	<div class="drawer-side h-screen z-20">
		<label for="my-drawer" class="drawer-overlay"></label>
		<ul class="menu p-4 w-80 bg-base-200 text-base-content h-screen flex flex-col overflow-y-auto">
			<!-- Topo do menu com logo -->
			<li class="menu-title text-center mb-4">
				<span><Logo /></span>
			</li>

			<!-- Accordion para as seções -->
			{#each menuItems as section, index}
				<div
					tabindex={index}
					class="collapse collapse-arrow font-medium text-lg rounded-box mb-2 bg-base-300"
				>
					<input type="checkbox" class="peer" />
					<div class="collapse-title text-sm uppercase tracking-wide text-lg text-secondary">
						{section.title}
					</div>
					<div class="collapse-content">
						{#each section.items as item}
							<li>
								<a
									href={item.link}
									class="flex items-center space-x-3 hover:bg-base-300 rounded-lg transition duration-200 ease-in-out text-lg"
									target={item.external ? '_blank' : '_self'}
								>
									<svelte:component this={item.icon} class="h-6 w-6 text-accent" />
									<span>{item.name}</span>
								</a>
							</li>
						{/each}
					</div>
				</div>
			{/each}
			<ThemeChanger />
			<!-- ThemeChanger e Logout no fim da página -->
			<div class="mt-auto mb-4">
				<div class="flex justify-between items-center">
					<button
						class="btn btn-outline btn-error flex items-center justify-center w-full h-full relative group"
						on:click={handleLogout}
					>
						<span
							class="username transition-opacity duration-300 ease-in-out group-hover:opacity-0"
						>
							{username}
						</span>
						<span
							class="logout flex items-center justify-center gap-3 absolute opacity-0 transition-opacity duration-400 ease-in-out group-hover:opacity-100"
						>
							Logout
							<Logout />
						</span>
					</button>
				</div>
			</div>
		</ul>
	</div>
</div>

<style>
	.collapse-title {
		cursor: pointer;
	}
	/* Faz com que o drawer seja rolável */
	.drawer-side {
		overflow-y: auto;
	}
</style>
