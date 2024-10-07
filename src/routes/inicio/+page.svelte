<script lang="ts">
	import { IconTableDown, IconTruck, IconSignature, IconCloud } from '@tabler/icons-svelte';
	import BentoGrid from '$lib/components/ui/BentoGrid/BentoGrid.svelte';
	import BentoGridItem from '$lib/components/ui/BentoGrid/BentoGridItem.svelte';
	import WavyBackground from '$lib/components/ui/WavyBackground/WavyBackground.svelte';
	const items = [
		{
			title: 'Lançamento de Pré Notas',
			description: 'Verifique ou lance pré documentos de entrada.',
			lottie: 'https://cdn.lordicon.com/fjvfsqea.json',
			href: '/lancamento-notas',
			className: 'md:col-span-2',
			icon: IconTableDown
		},
		{
			title: 'Conferência de Saída de Pneus.',
			description: 'Precisa entregar pneus, lance aqui!',
			href: '/lancamento-notas',
			lottie: 'https://cdn.lordicon.com/zzjjvkam.json',
			className: 'md:col-span-1',
			icon: IconTruck
		},
		{
			title: 'Gerador de Assinaturas',
			description: 'Acabou de entrar na empresa e precisa de uma assinatura? Crie já a sua.',
			href: '/lancamento-notas',
			lottie: 'https://cdn.lordicon.com/fikcyfpp.json',
			className: 'md:col-span-1',
			icon: IconSignature
		},
		{
			title: 'Intranet',
			description: 'Quer conhecer mais sobre nossos processos? Acesse a nossa nova documentação.',
			href: '/lancamento-notas',
			lottie: 'https://cdn.lordicon.com/fqbvgezn.json',
			className: 'md:col-span-2',
			icon: IconCloud
		}
	];
</script>

<WavyBackground className="mx-auto flex flex-col gap-24">
	<div>
		<p class="inter-var text-center text-2xl font-bold text-white md:text-4xl lg:text-7xl">
			Bem-vindo ao RodoApp!
		</p>
		<p class="inter-var mt-4 text-center text-base font-normal text-white md:text-lg">
			Acesse nossos aplicativos e descubra ferramentas práticas para o seu dia a dia.
		</p>
	</div>
	<BentoGrid className="max-w-6xl mx-auto md:auto-rows-[20rem]">
		{#each items as item, i (i)}
			<BentoGridItem
				title={item.title}
				description={item.description}
				href={item.href}
				lottie={item.lottie}
				className={item.className}
			>
				<div
					slot="header"
					class="flex h-full min-h-[6rem] w-full flex-1 rounded-xl border border-transparent bg-neutral-100 bg-dot-black/[0.2] [mask-image:radial-gradient(ellipse_at_center,white,transparent)] dark:border-white/[0.2] dark:bg-black dark:bg-dot-white/[0.2]"
				></div>
				<svelte:component this={item.icon} slot="icon" class="h-4 w-4 text-neutral-500" />
			</BentoGridItem>
		{/each}
	</BentoGrid>
</WavyBackground>
