<script lang="ts">
	import { goto } from '$app/navigation';
	import Logo from '$lib/components/theme/logo.svelte';
	import { saveLoginData, getLoginData, clearLoginData } from '$lib/utils/idb';

	let username: string = '';
	let password: string = '';
	let error: string = '';

	// Função para carregar os dados de login do IndexedDB
	async function loadLoginData() {
		// Certifica-se de que estamos no cliente antes de tentar acessar o IndexedDB
		if (typeof window !== 'undefined') {
			const data = await getLoginData();
			if (data) {
				username = data.username;
			}
		}
	}

	// Chama a função ao carregar a página
	loadLoginData();

	// Função de login
	async function handleLogin() {
		error = '';

		try {
			const bodyData = {
				username: username,
				password: password
			};

			const res = await fetch('http://rodoapp:8080/api/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(bodyData)
			});

			const data = await res.json();

			if (!res.ok) {
				throw new Error(data.message || 'Falha no login');
			}

			// Salva os dados no IndexedDB
			await saveLoginData(username, data.token);

			// Redireciona o usuário para a rota protegida
			goto('/prenotas');
		} catch (err) {
			if (err instanceof Error) {
				error = err.message || 'Ocorreu um erro desconhecido';
			} else {
				error = 'Ocorreu um erro desconhecido';
			}
		}
	}

</script>

<!-- Página de Login com Vídeo de Fundo -->
<div class="">
	<!-- Vídeo de Fundo -->
	<video autoplay muted loop playsinline class="absolute inset-0 w-screen h-screen object-cover">
		<source src="/background.mp4" type="video/mp4" />
		Seu navegador não suporta vídeos em HTML5.
	</video>

	<!-- Overlay para criar um efeito de blur e escurecimento sobre o vídeo -->
	<div class="absolute inset-0 bg-black opacity-20"></div>

	<!-- Card de Login -->
	<div
		class="relative z-10 bg-base-200 p-8 rounded-box shadow-lg w-96 backdrop-blur-lg bg-opacity-40"
	>
		<Logo />
		<!-- Exibe mensagem de erro -->
		{#if error}
			<div class="alert alert-error">
				<span>{error}</span>
			</div>
		{/if}

		<!-- Formulário de Login -->
		<form on:submit|preventDefault={handleLogin} class="space-y-4">
			<div class="form-control">
				<label class="label" for="username">
					<span class="label-text text-white">Usuário</span>
				</label>
				<input
					id="username"
					type="text"
					class="input input-bordered w-full"
					bind:value={username}
					required
				/>
			</div>

			<div class="form-control">
				<label class="label" for="password">
					<span class="label-text text-white">Senha</span>
				</label>
				<input
					id="password"
					type="password"
					class="input input-bordered w-full"
					bind:value={password}
					required
				/>
			</div>

			<button type="submit" class="btn btn-primary w-full"> Login </button>
		</form>
	</div>
</div>
