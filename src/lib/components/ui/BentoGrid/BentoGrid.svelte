<script lang="ts">
	import { cn } from '$lib/utils/index';

	export let className: string | undefined = undefined;
</script>

<div
	class={cn(
		'mx-auto grid grid-cols-1 gap-4 md:auto-rows-[25rem] md:grid-cols-3 ',
		className
	)}
>
	<slot />
</div>