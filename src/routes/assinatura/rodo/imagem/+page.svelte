<script>
	// @ts-nocheck
	import { onMount } from 'svelte';
	import toast, { Toaster } from 'svelte-french-toast';
	let nome = '';
	let cargo = '';
	let email = '';
	let telefone = '';
	let celular = '';
	let fin = false;

	// Extrair parâmetros da URL
	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		nome = params.get('nome') || '';
		cargo = params.get('cargo') || '';
		email = params.get('email') || '';
		telefone = params.get('telefone') || '';
		celular = params.get('celular') || '';
		fin = params.get('fin') === 'true';
	});
	// Variável para controlar a visibilidade do toast
	let showToast = false;

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
			toast.success('Texto Copiado com sucesso!', {
				icon: '👌🏼'
			});

			// Oculta o toast após 3 segundos
			setTimeout(() => {
				showToast = false;
			}, 3000);
		} catch (err) {
			toast.error('Erro ao copiar o texto: ', err);
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

	<!-- Layout sem classes externas -->
	<div id="output" class="signature-container" style="width:60vw; padding: 20px;">
		<table
			cellpadding="0"
			cellspacing="0"
			border="0"
			style="width: 100%; font-family: Arial, sans-serif;"
		>
			<tr>
				<!-- Coluna esquerda com logo -->
				<td style="padding: 5px 15px 0 28px; text-align:right; border-right: 1px solid black;">
					<a href="http://www.rodoparana.com.br">
						<img
							src="http://140.238.186.242/signaturegen/imagens/30anos.png"
							alt=""
							width="155px"
							height="auto"
							style="width:fit-content; height:fit-content"
						/>
						<br />
					</a>
				</td>
				<!-- Coluna direita com nome e cargo -->
				<td style="padding: 0 15px 0 15px; line-height:14pt; align-content:flex-start">
					<p style="font-weight: bold; font-style: italic; font-size: 12pt; margin: 3px;">
						{nome}
					</p>
					<p
						style="font-weight: bold; font-size: 9pt; font-style: italic; margin: 3px; width:250px;"
					>
						{cargo}
					</p>
				</td>
			</tr>
			<tr style="height: 100%;">
				<!-- Coluna esquerda com ícones sociais e website -->
				<td
					style="padding: 0 15px 0 28px; text-align: right; border-right: 1px solid black; line-height:12px; "
				>
					<div>
						<a href="https://www.facebook.com/rodoparana" style="display: inline-block; ">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/facebook.png"
								width="15"
								height="15"
								alt=""
								style="margin-right: 5px;"
							/>
						</a>
						<a href="https://www.instagram.com/rodoparanarandon/" style="display: inline-block;">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/instagram.png"
								width="15"
								height="15"
								alt=""
								style="margin-right: 5px;"
							/>
						</a>
						<a href="https://www.youtube.com/@rodoparanarandon6682" style="display: inline-block;">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/youtube.png"
								width="15"
								height="15"
								alt=""
							/>
						</a>
					</div>

					<div style="padding-top: 5px; line-height:10px;">
						<a
							href="http://www.rodoparana.com.br"
							style="font-size: 7pt; color: #009fe3; text-decoration: none;"
							>www.rodoparana.com.br</a
						>
					</div>
				</td>

				<!-- Coluna direita com telefone, celular e e-mail -->
				<td style="padding: 0px 0 0 15px;">
					<p style="font-size: 7pt; ">Fone: +55 {telefone}</p>
					{#if celular !== ''}
						<div style="max-height:10px; margin-bottom:10px;">
							<img
								src="http://140.238.186.242/signaturegen/imagens/social/whatsapp.png"
								width="12"
								height="12"
								alt=""
								style="max-height:fit-content; display: inline-block;"
							/>
							<p style="font-size: 7pt; display:inline-block">
								+55 {celular}
							</p>
						</div>
					{/if}
					<a
						href="mailto:{email}"
						style="font-size: 7pt; padding-top:15px;  color: #00acee;"
						>{email}</a
					>
				</td>
			</tr>
		</table>

		<!-- Financeiro -->

		<br /><br />{#if fin}
			<table
				style="font-family:arial; color:black; text-align: center; background:#ffffff; width:100%; border-top: 1px solid #757575;"
			>
				<tr>
					<td>
						<strong style="font-size:12pt; color: #009fe3;">COMUNICADO IMPORTANTE!</strong>
						<p style="font-size:10pt; ">
							Verificamos que estão sendo enviados boletos bancários falsos em nome da <br />
							Rodoparana Implementos Rodoviários Ltda. Informamos que todos os boletos da empresa
							<br />são encaminhados via <b style="color: #009fe3; font-size:11pt;">Correios</b>,
							entregues
							<b style="color: #009fe3; font-size:11pt;">pessoalmente</b> <br />ou encaminhados por
							e-mails com domínio <b style="color: #009fe3; font-size:11pt;">@rodoparana.com.br</b>
							quando solicitados pelo cliente.
							<small
								><br /><br />
								Em caso de dúvida, favor entrar em contato com o departamento <br />financeiro
								através do telefone
								<b style="color: #009fe3;font-size:10pt; ">(41) 3317-1414</b></small
							><br />
						</p>
					</td>
				</tr>
			</table>
		{/if}
	</div>

	<!-- Toast usando DaisyUI -->
	{#if showToast}
		<div class="toast toast-top toast-center">
			<div class="alert alert-info">
				<span>Assinatura copiada para a área de transferência!</span>
			</div>
		</div>
	{/if}
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
</style>
