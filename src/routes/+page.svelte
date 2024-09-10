<script>
	import { goto } from '$app/navigation';
	import Logo from '$lib/components/theme/logo.svelte';

	let username = '';
	let password = '';
	let error = '';

	// Função de login
	async function handleLogin() {
		error = ''; // Reseta a mensagem de erro

		try {
			const bodyData = {
				username: username,
				password: password
			};

			const res = await fetch('http://localhost:8080/api/login', {
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

			// Salva o token e o username no sessionStorage
			sessionStorage.setItem('token', data.access_token);
			sessionStorage.setItem('username', username);

			// Redireciona o usuário para a rota protegida
			goto('/prenotas');
		} catch (err) {
			// Verificação do tipo de erro para evitar erro de tipo 'unknown'
			if (err instanceof Error) {
				error = err.message || 'Ocorreu um erro desconhecido';
			} else {
				error = 'Ocorreu um erro desconhecido';
			}
		}
	}

	// Função de logout para remover o token e redirecionar
	function handleLogout() {
		sessionStorage.removeItem('token');
		sessionStorage.removeItem('username');
		goto('/login'); // Redireciona para a página de login
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
	<div class="relative z-10 bg-base-200 p-8 rounded-box shadow-lg w-96 backdrop-blur-lg bg-opacity-40">
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
