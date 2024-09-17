<script lang="ts">
	import { onMount } from 'svelte';
	let solicitacao = ''; // Texto da solicitação
	let emailInput = ''; // Input temporário para e-mails
	let emails: { email: string, color: string }[] = []; // Lista de e-mails com cor
	export let rec: number; // Rec recebido por props
	let finalizacao = false; // Indica se a revisão foi concluída
	let username = 'Usuário'; // Nome de usuário buscado do sessionStorage
	let isLoading = false; // Controle de carregamento
	let errorMessage = ''; // Mensagem de erro

	// Função para adicionar um e-mail ao pressionar Enter ou Tab
	function adicionarEmail() {
		const trimmedEmail = emailInput.trim();
		// Verifica se o e-mail é válido e permitido pelos domínios
		if (trimmedEmail !== '' && !emails.some(e => e.email === trimmedEmail)) {
			const color = definirCorEmail(trimmedEmail); // Define a cor do e-mail baseado no domínio
			emails = [...emails, { email: trimmedEmail, color }]; // Adiciona o e-mail à lista
		}
		// Limpa o campo de entrada
		emailInput = '';
	}

	// Função para validar o formato do e-mail e permitir apenas domínios autorizados
	function validarEmail(email: string) {
		const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
		const isValidFormat = emailRegex.test(email);
		const isAllowedDomain = email.endsWith('@rodoparana.com.br') || email.endsWith('@grupotimber.com.br');

		return isValidFormat && isAllowedDomain;
	}

	// Função para definir a cor do e-mail com base no domínio
	function definirCorEmail(email: string) {
		if (email.endsWith('@rodoparana.com.br')) {
			return 'badge-info';
		} else if (email.endsWith('@grupotimber.com.br')) {
			return 'badge-warning';
		} else {
			return 'badge-error'; // E-mails de domínios externos
		}
	}

	// Função para buscar o username no sessionStorage
	onMount(() => {
		const storedUsername = sessionStorage.getItem('username');
		if (storedUsername) {
			username = storedUsername;
		}
	});

	// Função para enviar a solicitação de revisão
	async function enviarSolicitacao(finalizacao = false) {
		// Verifica se há algum e-mail inválido
		const hasInvalidEmails = emails.some(email => email.color === 'badge-error');

		if (hasInvalidEmails) {
			errorMessage = 'Um ou mais e-mails são de domínios externos não permitidos!';
		} else {
			// Lógica de envio da solicitação
			const body = {
				RECSF1: rec,
				REVISAR: solicitacao,
				USER: username,
				EMAILS: emails.map(email => email.email).join(';'), // Concatena e-mails com ;
				FINALIZADO: finalizacao
			};

			try {
				isLoading = true;
				const response = await fetch('http://172.16.99.174:8400/rest/PreNota/RevisaoPreNota', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify(body)
				});

				if (!response.ok) {
					throw new Error('Erro ao enviar a solicitação.');
				}

				console.log('Solicitação enviada com sucesso');
				errorMessage = ''; // Limpa a mensagem de erro
			} catch (error) {
				console.error(error);
				errorMessage = 'Erro ao enviar a solicitação.';
			} finally {
				isLoading = false;
			}
		}
	}
</script>
  
  <div class="w-full flex justify-between">
	<h2 class="text-3xl text-center font-bold pb-10">Solicitar Revisão</h2>
  </div>
  
  <!-- Formulário para solicitar revisão -->
  <div class="flex flex-col w-full flex-1 gap-3">
	<textarea
	  placeholder="Digite sua solicitação"
	  bind:value={solicitacao}
	  class="textarea textarea-bordered flex-1"
	/>
	<div class="flex gap-2">
	  <input
		placeholder="Adicionar e-mails de acompanhamento"
		bind:value={emailInput}
		class="input input-bordered w-full"
		on:keydown={(e) => (e.key === 'Enter' || e.key === 'Tab') && adicionarEmail()}
	  />
	</div>
  
	<!-- Exibe os e-mails adicionados como badges -->
	<div class="flex flex-wrap gap-2">
	  {#each emails as { email, color }}
		<div class={`badge badge-outline ${color}`}>{email}</div>
	  {/each}
	</div>
  
	<!-- Exibe uma mensagem de erro se houver e-mails inválidos -->
	{#if errorMessage}
	  <div class="text-red-500 font-semibold">{errorMessage}</div>
	{/if}
  
	<div class="flex justify-between mt-4">
	  <button class="btn btn-outline btn-info w-fit" on:click={enviarSolicitacao}>
		Solicitar revisão
	  </button>
	  <button class="btn btn-outline btn-success w-fit">Concluir Revisão</button>
	</div>
  </div>
  