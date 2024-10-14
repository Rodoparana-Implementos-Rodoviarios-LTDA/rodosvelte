<script lang="ts">
	// Props para o componente: podemos passar o tamanho da logo e do texto
	export let size: string = 'w-20 h-20'; // Tamanho padrão (20x20) para a logo
	export let textSize: string = 'text-5xl'; // Tamanho padrão para o texto
	export let spinSpeed: string = 'animate-spin-slow'; // Animação padrão
</script>

<div class="flex items-center justify-center gap-3">
	<!-- Logo com tamanho controlado por props -->
	<img src="/rodoapp.svg" alt="logo" class={`${size} ${spinSpeed}`} />

	<!-- Texto com tamanho controlado por props -->
	<span class={`${textSize} font-bold text-blue-500`}>RodoApp</span>
</div>
