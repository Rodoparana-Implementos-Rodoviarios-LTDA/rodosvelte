<script lang="ts">
	import Cabecalho from '$lib/components/ui/xml/Cabecalho.svelte';
	import ProdutosXMLTable from '$lib/components/ui/xml/ProdutosXMLTable.svelte';
</script>

<div class="fixed top-8 left-1/2 transform -translate-x-1/2 text-center">
	<h1 class="text-3xl font-bold">Importação de XML</h1>
</div>
<main class="h-[85vh] w-[98vw] flex flex-col justify-end items-baseline gap-10">
	<header class="h-fit w-full">
		<Cabecalho />
	</header>
	<body class="w-full h-full"><ProdutosXMLTable /></body>
</main>
