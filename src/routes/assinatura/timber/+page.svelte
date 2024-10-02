<script>
	import Rodo from '$lib/assets/Rodo.svelte';

	let nome = '';
	let cargo = '';
	let email = '';
	let skype = '';
	let telefone = '';
	let celular = '';
	let fin = false; // Verifique se precisa dessa variável aqui
	let logoSelecionada = 'Timber'; // Valor padrão para a logo
	let filialSelecionada = ''; // Valor para a filial selecionada
	let empresasSelecionadas = []; // Valor para as empresas selecionadas

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
	function handleEmpresaChange(event) {
		const selectedOptions = Array.from(event.target.selectedOptions);
		empresasSelecionadas = selectedOptions.map((option) => option.value);
		console.log(empresasSelecionadas); // Verifique o array de empresas selecionadas
	}
	function handleSubmit(event) {
		event.preventDefault();

		if (!nome || !cargo || !email || !telefone) {
			alert('Por favor, preencha todos os campos obrigatórios.');
			return;
		}

		// Redireciona para a página de assinatura com os parâmetros
		const empresas = empresasSelecionadas.join(','); // Concatena as empresas selecionadas
		window.location.href = `/assinatura/timber/imagem?nome=${nome}&cargo=${cargo}&email=${email}&skype=${skype}&telefone=${telefone}&celular=${celular}&logo=${logoSelecionada}&filial=${filialSelecionada}&empresas=${empresas}`;
	}
</script>

<body class="flex flex-col justify-center items-center">
	<div class="flex flex-col items-center gap-10">
		<img
			src="http://140.238.186.242/signaturegen/imagens/timber_bkp.png"
			alt="Timberforest"
			class="w-64 object-contain pb-5"
		/>
		<div class="text-center">
			<h1 class="text-3xl font-bold">Modelo de Assinatura do Grupo Timber</h1>
			<p>Preencha os dados para adquirir sua nova assinatura.</p>
		</div>
		<!-- Formulário -->
		<form id="formDados" on:submit|preventDefault={handleSubmit} class="flex flex-col gap-10">
			<div class="flex gap-10">
				<div class="max-w-xl space-y-4">
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
					<!-- Skype -->
					<input
						type="text"
						id="skype"
						bind:value={skype}
						placeholder="nomesobrenome.empresa"
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
				</div>
				<div class="card bg-base-200 px-10 py-5 form-control">
					<fieldset id="logo" class="gap-2 flex flex-col mb-5">
						<label for="logo" class="font-bold">Selecione a Logo:</label>
						<div class="flex justify-start items-center gap-2">
							<input
								type="radio"
								name="logo"
								class="radio"
								checked
								value="Timber"
								bind:group={logoSelecionada}
							/>
							<label for="">Timber</label>
							<input
								type="radio"
								name="logo"
								class="radio"
								value="Timber Consórcio"
								bind:group={logoSelecionada}
							/>
							<label for="">Timber Consórcio</label>
						</div>
					</fieldset>
					<label class="form-control w-full max-w-xs">
						<span class="label-text">Filial:</span>
						<select class="select select-bordered" bind:value={filialSelecionada}>
							<option disabled>Selecione a Filial:</option>
							<option>Curitiba/PR</option>
							<option>Guaíba/RS</option>
							<option>Lages/SC</option>
							<option>Pelotas/RS</option>
							<option>Curvelo/MG</option>
							<option>Três Lagoas/RS</option>
							<option>Curvelo/MG</option>
							<option>Palhoça/SC</option>
							<option>Marialva/PR</option>
							<option>Chapecó/SC</option>
						</select>
					</label>
					<label class="form-control w-full max-w-xs">
						<span class="label-text py-2">Empresa:</span>
						<select class="select select-bordered" multiple on:change={handleEmpresaChange}>
							<option>EcoFlow</option>
							<option>Riddara</option>
							<option>Fuchs</option>
							<option>Ponsse</option>
							<option>XAG</option>
							<option>PIX4D</option>
							<option>SANY</option>
						</select>
					</label>
				</div>
			</div>
			<!-- Botão de Gerar -->
			<div class="flex justify-center space-x-4">
				<button type="submit" class="btn btn-primary">Gerar Assinatura</button>
			</div>
		</form>
	</div>
</body>
