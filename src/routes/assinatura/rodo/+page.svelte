<script>
	import Rodo from '$lib/assets/Rodo.svelte';

	let nome = '';
	let cargo = '';
	let email = '';
	let telefone = '';
	let celular = '';
	let fin = false;

	const formatTelefone = (value, mask) => {
		const digits = value.replace(/\D/g, '');
		let formatted = '';
		let maskIndex = 0;

		for (let char of mask) {
			if (char === 'X') {
				if (digits[maskIndex]) {
					formatted += digits[maskIndex];
					maskIndex++;
				} else {
					break;
				}
			} else {
				formatted += char;
			}
		}
		return formatted;
	};

	function handleTelefoneInput(event) {
		telefone = formatTelefone(event.target.value, '(XX) XXXX.XXXX');
	}

	function handleCelularInput(event) {
		celular = formatTelefone(event.target.value, '(XX) XXXXX.XXXX');
	}

	function handleSubmit(event) {
		event.preventDefault();

		if (!nome || !cargo || !email || !telefone) {
			alert('Por favor, preencha todos os campos obrigatórios.');
			return;
		}

		// Redireciona para a página de assinatura com os parâmetros
		window.location.href = `/assinatura/rodo/imagem?nome=${nome}&cargo=${cargo}&email=${email}&telefone=${telefone}&celular=${celular}&fin=${fin}`;
	}
</script>

<body class="flex flex-col justify-center items-center">
	<div class="flex flex-col items-center justify-evenly">
		<Rodo width="250" height="auto" color="text-neutral-content"/>
		<h1 class="text-3xl font-bold">Modelo de Assinatura da Rodoparaná</h1>
		<p>Preencha os dados para adquirir sua nova assinatura.</p>
		<!-- Formulário -->
		<form
			id="formDados"
			on:submit|preventDefault={handleSubmit}
			class="container mx-auto px-4 py-8"
		>
			<div class="max-w-xl mx-auto space-y-4">
				<!-- Nome -->
				<input
					type="text"
					id="nome"
					bind:value={nome}
					placeholder="Nome completo"
					class="input input-primary input-bordered w-full mt-1"
					required
				/>

				<!-- Cargo -->
				<input
					type="text"
					id="cargo"
					bind:value={cargo}
					placeholder="Seu cargo"
					class="input input-bordered input-primary w-full mt-1"
					required
				/>

				<!-- E-mail -->
				<input
					type="email"
					id="email"
					bind:value={email}
					placeholder="seuemail@dominio.com"
					class="input input-primary input-bordered w-full mt-1"
					required
				/>

				<!-- Telefone -->
				<input
					type="text"
					id="telefone"
					bind:value={telefone}
					on:input={handleTelefoneInput}
					placeholder="(xx) XXXX.XXXX"
					maxlength="14"
					class="input input-primary input-bordered w-full mt-1"
					required
				/>

				<!-- Celular -->
				<input
					type="text"
					id="celular"
					bind:value={celular}
					on:input={handleCelularInput}
					placeholder="(xx) XXXXX.XXXX"
					maxlength="15"
					class="input input-primary input-bordered w-full mt-1"
				/>

				<!-- Financeiro -->
				<div class="flex items-center">
					<input
						type="checkbox"
						id="fin"
						bind:checked={fin}
						class="checkbox checkbox-primary mr-2"
					/>
					<label for="fin" class="text-sm font-medium text-content">Financeiro</label>
				</div>

				<!-- Botão de Gerar -->
				<div class="flex justify-center space-x-4">
					<button type="submit" class="btn btn-primary"> Gerar Assinatura </button>
				</div>
			</div>
		</form>
	</div>
</body>
