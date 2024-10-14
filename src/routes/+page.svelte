<script>
	import { goto } from '$app/navigation';
	import { IconTruckReturn } from '@tabler/icons-svelte';
	import { clearIndexedDB } from '$lib/services/idb'; // Importar a função para limpar o IndexedDB

	let username = '';
	let password = '';
	let error = '';

	// Função de login
	async function handleLogin() {
		error = '';

		// Limpar o IndexedDB antes de realizar o login
		await clearIndexedDB();

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

			// Salva o token e o username no sessionStorage
			sessionStorage.setItem('token', data.access_token);
			sessionStorage.setItem('username', username);

			// Redireciona o usuário para a rota protegida
			goto('/inicio');
		} catch (err) {
			// Verificação do tipo de erro para evitar erro de tipo 'unknown'
			if (err instanceof Error) {
				error = err.message || 'Ocorreu um erro desconhecido';
			} else {
				error = 'Ocorreu um erro desconhecido';
			}
		}
	}
</script>

<!-- Página de Login com Formulário e Vídeo -->
<div class="bg-blue-900/40 flex justify-center items-center w-screen h-screen">
	<div class="card lg:card-side h-[80vh] w-[80dvw] bg-base-100 shadow-xl grid grid-cols-2">
		<!-- Vídeo no lado esquerdo -->
		<div class="relative h-full bg-base-100 col-span-1 rounded-xl">
			<video autoplay muted loop class="object-cover h-full w-full rounded">
				<source src="/background.mp4" type="video/mp4" />
				Seu navegador não suporta vídeos em HTML5.
			</video>
			<!-- Overlay para criar um efeito de escurecimento -->
			<div class="absolute inset-0 bg-black opacity-40"></div>
		</div>

		<div class="flex flex-col justify-evenly items-center">
			<div class="flex items-start justify-center">
				<div class=" flex items-center justify-center">
					<IconTruckReturn class="w-16 h-16 text-primary" />
					<span class="text-5xl font-bold text-primary">RodoApp</span>
				</div>
			</div>

			<!-- Formulário de Login no lado direito -->
			<div class="col-span-1 flex flex-col justify-center items-center space-y-10">
				<div class="space-y-3">
					<h1 class="text-6xl font-medium">Bem vindo de Volta!</h1>
					<p class="text-2xl">Entre com suas credenciais para acessar o sistema.</p>
				</div>
				<form on:submit|preventDefault={handleLogin} class="space-y-4 w-full max-w-lg">
					<div class="form-control w-full">
						<label class="label" for="username">
							<span class="label-text">Usuário do Protheus</span>
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
							<span class="label-text">Senha</span>
						</label>
						<input
							id="password"
							type="password"
							class="input input-bordered w-full"
							bind:value={password}
							required
						/>
						{#if error}
							<div class="alert alert-error mt-2">
								<span>{error}</span>
							</div>
						{/if}
					</div>

					<button type="submit" class="btn btn-primary w-full"> Login </button>
				</form>
			</div>

			<div class="w-full flex justify-evenly items-baseline">
				<button class="btn btn-outline btn-primary"
					><a href="http://hesk.rodoparana.com.br/">SUPORTE</a></button
				>
				<button class="btn btn-outline btn-primary"
					><a href="/intranet">INTRANET</a></button
				>
		</div>
	</div>
</div>
