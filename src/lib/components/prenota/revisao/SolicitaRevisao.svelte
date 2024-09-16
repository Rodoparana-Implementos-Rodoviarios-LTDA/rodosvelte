<script>
	// Variáveis para capturar dados da solicitação de revisão
	let solicitacao = ''; // Texto da solicitação
	let email = ''; // E-mail(s) de acompanhamento

	// Função para lidar com o envio da solicitação
	function enviarSolicitacao() {
		// Aqui você pode fazer a lógica para enviar a solicitação
		console.log('Solicitação:', solicitacao);
		console.log('E-mails de acompanhamento:', email);
	}
</script>

<div class="w-full flex justify-between">
	<h2 class="text-3xl text-center font-bold pb-10">Solicitar Revisão</h2>
	<button class="btn btn-outline btn-info w-fit">Anexos</button>
</div>

<!-- Formulário para solicitar revisão -->
<div class="flex flex-col w-full flex-1 gap-3">
	<textarea
		placeholder="Digite sua solicitação"
		bind:value={solicitacao}
		class="textarea textarea-bordered flex-1"
	/>
	<input
		placeholder="E-mails de acompanhamento"
		bind:value={email}
		class="input input-bordered w-full"
	/>
	<button class="btn btn-outline btn-info w-fit" on:click={enviarSolicitacao}> Solicitar revisão </button>
</div>
