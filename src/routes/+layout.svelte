<script lang="ts">
	import 'tailwindcss/tailwind.css';
	import Nav from '$lib/components/theme/nav.svelte';
	import ThemeChanger from '$lib/components/theme/changeTheme.svelte';
	import { page } from '$app/stores'; // Para acessar a URL atual
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import

	onMount(() => {
		const token = sessionStorage.getItem('token'); // Obtém o token do sessionStorage

		if (!token && $page.url.pathname !== '/') {
			// Se o token não existir e não for a página de login, redireciona para a página de login (raiz /)
			goto('/');
		}
	});
</script>

<html lang="pt-br" class="h-screen w-screen max-w-screen max-h-screen">
	<header class=" z-10">
		<!-- Exibe a navegação se a rota não for "/" (login) -->
		<div class="">
			{#if $page.url.pathname !== '/'}
				<Nav />
				<ThemeChanger />
			{/if}

		</div>
	</header>

	<body class="min-h-screen max-h-screen w-screen flex justify-center items-center">
		<slot />
	</body>
</html>
