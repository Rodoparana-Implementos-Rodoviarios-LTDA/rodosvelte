<script>

	import Rodo from "$lib/assets/Rodo.svelte";

</script>
<svelte:head>
	<title>Assinaturas de Email Rodoparana</title>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>

<body class="flex flex-col min-h-screen w-screen justify-center items-center p-20">
	<div class="flex flex-col gap-10">
		<h1 class="text-5xl text-center">Criador de assinaturas de e-mail</h1>
		<p class="max-w-6xl text-center text-2xl">
			Esta ferramenta permite criar assinaturas de e-mail alinhadas com os padrões da empresa.
			<br />
			Modelos disponíveis: Rodoparana e Timberforest, incluindo logos adicionais das marcas Sany, Fuchs
			e Timberfleet.
		</p>

		<h2 class="text-2xl text-center">Selecione o modelo desejado:</h2>
	</div>
	<hr class="" />
	<!-- Conteúdo Principal -->

	<div class="flex justify-evenly w-full p-20">
		<a href="/assinatura/rodo">
			<Rodo width="250" height="auto" color="text-neutral-content"/>
		</a>
		<a href="/assinatura/timber">
			<img
				src="http://140.238.186.242/signaturegen/imagens/timber_bkp.png"
				alt="Timberforest"
				class="w-64 object-contain"
			/>
		</a>
	</div>
	<div
		class="flex flex-col md:flex-row justify-center items-center space-y-4 md:space-y-0 md:space-x-4"
	>
		<a href="http://hesk.rodoparana.com.br/" class="btn btn-primary"> Preciso de ajuda </a>
		<a
			href="https://sites.google.com/site/baserodoparana/ti/email/assinatura-de-email"
			class="btn btn-primary"
		>
			Como incluir a assinatura no Gmail/ Outlook?
		</a>
	</div>
</body>
