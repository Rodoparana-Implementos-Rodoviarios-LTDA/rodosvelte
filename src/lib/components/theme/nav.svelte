<script lang="ts">
	import {
		IconInfoCircle,
		IconPhone,
		IconLogout,
		IconDashboard,
		IconList,
		IconPlus,
		IconUpload,
		IconShieldCheck,
		IconCategory,
		IconBulldozer,
		IconTruck
	} from '@tabler/icons-svelte';
	import { goto } from '$app/navigation';
	import { deleteData } from '$lib/services/idb';
	import { IconTruckReturn, IconClockCheck } from '@tabler/icons-svelte';
	interface MenuItem {
		name: string;
		icon: typeof IconDashboard | typeof IconPhone;
		link: string;
		external?: boolean;
	}

	interface MenuSection {
		title: string;
		link?: string;
		items: MenuItem[];
	}

	const menuItems: MenuSection[] = [
		{
			title: 'Lançamento de Notas',
			items: [
				{ name: 'Lista de Pre Notas', icon: IconList, link: '/lancamento-notas/' },
				{ name: 'Incluir Manualmente', icon: IconPlus, link: '/lancamento-notas/incluir' },
				{ name: 'Incluir XML', icon: IconUpload, link: '/lancamento-notas/xml' }
			]
		},
		{
			title: 'Controle de Itens',
			items: [
				{ name: 'Borracharia', icon: IconDashboard, link: '/controle-itens' },
				{ name: 'Conferência', icon: IconClockCheck, link: '/controle-itens/conferencia' },
				{ name: 'Conferidos', icon: IconShieldCheck, link: '/controle-itens/conferidos' }
			]
		},
		{
			title: 'Assinatura de Email',
			items: [
				{
					name: 'Timber',
					icon: IconBulldozer,
					link: '/assinatura/timber'
				},
				{
					name: 'Rodoparaná',
					icon: IconTruck,
					link: '/assinatura/rodo'
				}
			]
		},
		{
			title: 'Suporte e Intranet',
			items: [
				{
					name: 'Suporte',
					icon: IconPhone,
					link: 'https://hesk.rodoparana.com.br',
					external: true
				},
				{
					name: 'Intranet',
					icon: IconInfoCircle,
					link: 'https://sites.google.com/site/baserodoparana/home?authuser=1',
					external: true
				}
			]
		}
	];

	let username = 'Usuário';
	let isLoading = true;

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

	async function handleLogout() {
		try {
			await deleteData('username');
			await deleteData('authToken');
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
		<label for="my-drawer" class="btn btn-square btn-ghost fixed top-5 left-5 z-10">
			<IconCategory class="h-6 w-6" />
		</label>
		<slot />
	</div>

	<!-- Drawer Lateral -->
	<div class="drawer-side h-screen z-20">
		<label for="my-drawer" class="drawer-overlay"></label>
		<ul class="menu p-4 w-80 bg-base-200 text-base-content h-screen flex flex-col overflow-y-auto">
			<div class=" flex items-center justify-center pb-10">
				<IconTruckReturn class="w-12 h-12 text-primary" />
				<span class="text-3xl font-bold text-primary">RodoApp</span>
			</div>
			<!-- Accordion para as seções -->
			{#each menuItems as section, index}
				<div class="collapse bg-base-300 mb-2">
					<!-- Usamos radio buttons para garantir que apenas uma seção esteja aberta -->
					<input type="radio" name="menu-accordion" id="section-{index}" class="peer" />
					<label for="section-{index}" class="collapse-title text-lg font-medium"
						>{section.title}</label
					>
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
									<svelte:component this={item.icon} class="h-6 w-6 text-accent" />
									<span>{item.name}</span>
								</a>
							</li>
						{/each}
					</div>
				</div>
			{/each}

			<!-- ThemeChanger e Logout no fim da página -->
			<div class="mt-auto mb-4">
				<button
					class="btn btn-outline btn-error flex items-center justify-center w-full h-full relative group"
					on:click={handleLogout}
				>
					<span class="username transition-opacity duration-300 ease-in-out group-hover:opacity-0">
						{username}
					</span>
					<span
						class="logout flex items-center justify-center gap-3 absolute opacity-0 transition-opacity duration-400 ease-in-out group-hover:opacity-100"
					>
						Logout
						<IconLogout />
					</span>
				</button>
			</div>
		</ul>
	</div>
</div>

<style>
	.collapse-title {
		cursor: pointer;
	}
	.drawer-side {
		overflow-y: auto;
	}
</style>
