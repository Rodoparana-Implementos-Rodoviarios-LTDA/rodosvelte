<script lang="ts">
	import '../globals.css';
	import Nav from '$lib/components/theme/nav.svelte';
	import ThemeChanger from '$lib/components/theme/changeTheme.svelte';
	import { page } from '$app/stores'; // Para acessar a URL atual
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';

	let token;

	// Função para verificar e redirecionar para login se o token não existir
	onMount(() => {
		token = sessionStorage.getItem('token'); // Obtém o token do sessionStorage

		if (!token && $page.url.pathname !== '/' && !$page.url.pathname.startsWith('/intranet')) {
			// Se o token não existir e a rota não for login ou /intranet, redireciona para a página de login
			goto('/');
		}
	});
</script>

<html lang="pt-br" class="h-screen w-screen">
	<header class="z-10">
		<!-- Exibe a navegação se a rota não for "/" (login) nem "/intranet" -->
		{#if $page.url.pathname !== '/' && !$page.url.pathname.startsWith('/intranet')}
			<Nav />
		{/if}
		<ThemeChanger />
	</header>

	<body class="min-h-screen max-h-screen w-screen flex justify-center items-center">
		<slot />
	</body>
</html>
