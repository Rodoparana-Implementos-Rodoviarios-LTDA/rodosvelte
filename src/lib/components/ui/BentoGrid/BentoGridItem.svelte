<script lang="ts">
	import { cn } from '$lib/utils/index';
	import { IconArrowBigRightLines } from '@tabler/icons-svelte';

	export let className: string | undefined = undefined;
	export let lottie: string | HTMLElement | undefined = undefined;
	export let title: string | HTMLElement | undefined = undefined;
	export let description: string | HTMLElement | undefined = undefined;
	export let href: string | HTMLElement | undefined = undefined;
	export let icon: HTMLElement | undefined = undefined;
</script>

<div
	class={cn(
		'group/bento row-span-1 flex flex-col justify-between space-y-4 rounded-xl border border-transparent bg-white p-4 shadow-input transition duration-200 hover:shadow-xl dark:border-white/[0.2] dark:bg-black dark:shadow-none',
		className
	)}
>
	<div class="items-center justify-center h-64">
		<lord-icon
			src={lottie}
			trigger="hover"
			stroke="light"
			state="hover-pinch"
			colors="primary:#3080e8,secondary:#b4b4b4"
			style="width:100%;height:100%"
		>
		</lord-icon>
	</div>

	<div class="transition duration-200 group-hover/bento:translate-x-2">
		<slot name="icon" />
		<div class="flex justify-between items-center">
			<div>
				<div class="mb-2 mt-2 font-sans font-bold text-neutral-600 dark:text-neutral-200">
					{title}
				</div>
				<div class="font-sans text-xs font-normal text-neutral-600 dark:text-neutral-300">
					{description}
				</div>
			</div>
			<div>
				<button class="btn"><a {href}><IconArrowBigRightLines /></a></button>
			</div>
		</div>
	</div>
</div>
