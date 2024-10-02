<script lang=ts>
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';
	import TimberBlack from '$lib/assets/TimberBlack.svelte';
	import Riddara from '$lib/assets/Riddara.svelte';

	let nome = '';
	let cargo = '';
	let email = '';
	let skype = '';
	let telefone = '';
	let celular = '';
	let logo = '';
	let filial = '';
	let empresasSelecionadas = [];
	let endereco1 = '';
	let endereco2 = '';
	let empresasCount = 0; // Contador de empresas selecionadas

	// Extrair parâmetros da URL
	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		nome = params.get('nome') || '';
		cargo = params.get('cargo') || '';
		email = params.get('email') || '';
		skype = params.get('skype') || '';
		telefone = params.get('telefone') || '';
		celular = params.get('celular') || '';
		logo = params.get('logo') || '';
		filial = params.get('filial') || '';

		const empresas = params.get('empresas') || '';
		empresasSelecionadas = empresas ? empresas.split(',') : [];
		empresasCount = empresasSelecionadas.length; // Conta o número de empresas
		console.log(empresasCount);
		getFilialEndereco();
	});

	// Função para copiar o conteúdo e mostrar o toast
	async function copyToClipboard() {
		try {
			const signatureElement = document.getElementById('output');
			const range = document.createRange();
			range.selectNode(signatureElement);
			const selection = window.getSelection();
			selection.removeAllRanges();
			selection.addRange(range);

			// Executa o comando de cópia
			document.execCommand('copy');
			selection.removeAllRanges(); // Remove a seleção após a cópia

			// Exibe o toast
			toast.success('Assinatura copiada com sucesso!', {
				icon: '👌🏼'
			});
		} catch (err) {
			toast.error('Erro ao copiar o texto: ', err);
		}
	}

	// Função para obter o endereço da filial e atribuir os valores diretamente
	function getFilialEndereco() {
		switch (filial.toLowerCase()) {
			case 'curitiba/pr':
				endereco1 = 'Av. Juscelino K. de Oliveira 3545';
				endereco2 = 'Cidade Industrial de Curitiba | Curitiba (PR)';
				break;
			case 'lages/sc':
				endereco1 = 'Av. Dr. João Pedro Arruda, 1437';
				endereco2 = 'Área Industrial | Lages (SC)';
				break;
			case 'guaíba/rs':
				endereco1 = 'Rua da Balança, 96';
				endereco2 = 'Ramada | Guaíba (RS)';
				break;
			case 'pelotas/rs':
				endereco1 = 'BR 116, SN';
				endereco2 = 'Distrito Rodoviário | Capão do Leão (RS)';
				break;
			case 'três lagoas/rs':
				endereco1 = 'BR 158, SN';
				endereco2 = 'Jardim Paranapunga | Três Lagoas (MS)';
				break;
			case 'curvelo/mg':
				endereco1 = 'Av. Bias Fortes, 2015';
				endereco2 = 'Tibira | Curvelo (MG)';
				break;
			case 'palhoça/sc':
				endereco1 = 'Rua José Cosme Pamplona';
				endereco2 = 'Bela Vista | Palhoça (SC)';
				break;
			case 'marialva/pr':
				endereco1 = 'BR-376, KM 188';
				endereco2 = 'Sarandi | Marialva - PR';
				break;
			case 'chapecó/sc':
				endereco1 = 'Ac Plinio Arlindo de Nes, 1915D';
				endereco2 = 'Belvedere | Chapecó (SC)';
				break;
			default:
				endereco1 = '';
				endereco2 = '';
		}
	}
</script>

<main class="h-screen w-screen flex flex-col justify-center items-center">
	<Toaster />
	<div style="text-align: center; margin-bottom: 20px;">
		<!-- Botão para copiar assinatura -->
		<button type="button" on:click={copyToClipboard} class="btn btn-primary">
			Copiar Assinatura
		</button>
	</div>

	<!-- Layout de assinatura -->
	<div id="output" class="signature-container" style="width:60vw; padding: 20px;">
		<table border="0" cellpadding="10" cellspacing="0" style="width: 100%; font-family: Arial, sans-serif;">
			<tr>
				<!-- Coluna esquerda com logo principal -->
				<td valign="top" style="border-right: 1px solid black; text-align: center; padding-right: 20px;">
					<img src="http://140.238.186.242/signaturegen/imagens/timber.png" alt="Logo Principal" width="160" height="auto" />
					
					<!-- Tabela para logos das empresas -->
					<table border="0" cellpadding="5" cellspacing="0" style="width: 100%; margin-top: 20px;">
						<tbody>
							{#each empresasSelecionadas as empresa, i}
								{#if i % 2 === 0}
									<tr>
								{/if}
								<td style="text-align: center;">
									{#if empresa === 'SANY'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/sany.png" width="80" alt="SANY Logo" />
									{/if}
									{#if empresa === 'Ponsse'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/ponsse.png" width="80" alt="Ponsse Logo" />
									{/if}
									{#if empresa === 'Fuchs'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/fuchs.png" width="80" alt="Fuchs Logo" />
									{/if}
									{#if empresa === 'XAG'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/xag.png" width="80" alt="XAG Logo" />
									{/if}
									{#if empresa === 'EcoFlow'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/ecoflow.png" width="80" alt="EcoFlow Logo" />
									{/if}
									{#if empresa === 'PIX4D'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/pix4d.png" width="80" alt="PIX4D Logo" />
									{/if}
									{#if empresa === 'Riddara'}
										<img src="http://140.238.186.242/signaturegen/imagens/empresas/riddara.png" width="80" alt="Riddara Logo" />
									{/if}
								</td>
							{/each}
						</tbody>
					</table>
				</td>

				<!-- Coluna direita com nome e cargo -->
				<td
					style="padding: 0 15px 0 15px; line-height:14pt; text-align: start; vertical-align: middle;"
				>
					<p style="font-weight: bold; font-style: italic; font-size: 12pt; margin: 3px;">{nome}</p>
					<p
						style="font-weight: bold; font-size: 9pt; font-style: italic; margin: 3px; width:250px;"
					>
						{cargo}
					</p>
					<p style="font-size: 7pt;">
						<b>Telefone:</b> +55 {telefone}
						{#if celular !== ''}
							| {celular}
						{/if}
						<br />
						<b>E-mail:</b>
						{email}<br />
						<b>Skype:</b>
						{skype}<br />
						<b>Endereço:</b>
						{endereco1}<br />{endereco2}<br />
					</p>
					<!-- Ícones sociais -->
					<div style="display:inline-flex; padding-top:15px; justify-content: center;">
						<a href="http://www.grupotimber.com.br">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/timber.png"
								width="30"
								height="30"
								style="margin-right: 5px; "
							/>
						</a>
						<a href="https://www.facebook.com/timberforestequipamentos/">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/facebook.png"
								width="30"
								height="30"
								style="margin-right: 5px; "
							/>
						</a>
						<a
							href="https://www.linkedin.com/company/timber-forest-equipamentos/?viewAsMember=true"
						>
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/linkedIn.png"
								width="30"
								height="30"
								style="margin-right: 5px; "
							/>
						</a>
						<a href="https://www.instagram.com/timber_br/?hl=pt-br">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/instagram.png"
								width="30"
								height="30"
								style="margin-right: 5px; "
							/>
						</a>
						<a href="https://www.youtube.com/channel/UCflgytZMsL78RUFhECk8IAQ">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/youtube.png"
								width="30"
								height="30"
								style="margin-right: 5px; "
							/>
						</a>
					</div>
				</td>
			</tr>
		</table>
	</div>
</main>

<style>
	.signature-container {
		display: flex;
		flex-direction: column;
		align-items: start;
		justify-content: start;
		background: #fff;
		max-width: fit-content;
		border-radius: 8px;
		color: black;
	}

	p {
		white-space: nowrap;
	}

	.toast {
		position: fixed;
		bottom: 20px;
		left: 50%;
		transform: translateX(-50%);
		z-index: 50;
	}

	.toast .alert {
		padding: 10px 20px;
	}
</style>
