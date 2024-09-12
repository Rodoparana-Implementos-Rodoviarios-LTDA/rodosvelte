<script lang="ts">
	import { IconCheck, IconDots } from '@tabler/icons-svelte';
	import type { Revisao } from '$lib/types/revisao';

	// Recebe os eventos como prop
	export let eventos: Revisao[] = [];

	// Variáveis para capturar dados da solicitação de revisão
	let solicitacao = ''; // Texto da solicitação
	let email = ''; // E-mail(s) de acompanhamento

	// Função para lidar com o envio da solicitação
	function enviarSolicitacao() {
		// Aqui você pode fazer a lógica para enviar a solicitação
		console.log('Solicitação:', solicitacao);
		console.log('E-mails de acompanhamento:', email);
	}
</script>

<div class="h-full flex flex-col justify-between items-center">
	<div class="w-full flex justify-between">
		<h2 class="text-3xl text-center font-bold pb-10">Solicitar Revisão</h2>
		<button class="btn btn-outline">Anexos</button>
	</div>

	<!-- Formulário para solicitar revisão -->
	<div class="flex flex-col w-full flex-1 gap-3">
		<textarea
			placeholder="Digite sua solicitação"
			bind:value={solicitacao}
			class="textarea textarea-bordered flex-1"
		/>
		<input
			placeholder="E-mails de acompanhamento"
			bind:value={email}
			class="input input-bordered w-full"
		/>
		<button class="btn btn-primary w-fit" on:click={enviarSolicitacao}>
			Solicitar revisão
		</button>
	</div>

	<div class="divider text-xl">Histórico de Eventos</div>

	<!-- Exibição do histórico de eventos -->
	<ul class="timeline timeline-snap-icon max-md:timeline-compact timeline-vertical max-h-[80vh] overflow-y-auto">
		{#each eventos as event, index}
			<li>
				{#if index > 0}
					<hr class="bg-primary" />
				{/if}
				<div class="timeline-start text-end">
					<div class="text-xl capitalize">{event.status}</div>
					<time class="text-xs font-mono">{event.data} - {event.hora}</time>
				</div>

				<div class="timeline-middle">
					{#if event.status === 'classificar'}
						<div class="rounded-full p-1 m-0.5 bg-primary/20">
							<IconDots stroke={3} class="h-3 w-3 text-black" />
						</div>
					{:else}
						<div class="rounded-full p-1 m-0.5 bg-primary">
							<IconCheck stroke={3} class="h-3 w-3 text-black" />
						</div>
					{/if}
				</div>

				<div class="timeline-end timeline-box w-full">
					<div class="flex justify-between">
						<div class="text-start">
							<div class="text-sm font-black">Usuário</div>
							<div class="text-sm font-light">{event.usuario}</div>
						</div>
						<div class="text-end">
							<div class="text-sm font-black">Campo</div>
							<div class="text-sm font-light">{event.campo}</div>
						</div>
					</div>

					{#if event.anterior}
						<div class="flex justify-between">
							<div class="text-start">
								<div class="text-sm font-black">Anterior</div>
								<div class="text-sm font-light">{event.anterior}</div>
							</div>
							<div class="text-end">
								<div class="text-sm font-black">Valor Atual</div>
								<div class="text-sm font-light max-w-36">{event.atual}</div>
							</div>
						</div>
					{/if}
				</div>

				<hr class="bg-primary" />
			</li>
		{/each}
	</ul>
</div>
