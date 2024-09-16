<script lang="ts">
	import { HoverCard } from 'radix-svelte'; // Se ainda estiver usando o Radix-Svelte HoverCard
	import {
	  IconCheck,
	  IconMessageCircle,
	  IconClock,
	  IconChevronsUp,
	  IconChevronsDown,
	  IconMinus,
	  IconFlag,
	} from '@tabler/icons-svelte';
  
	export let type: 'Status' | 'Prioridade'; // Tipos possíveis
	export let value: string; // O valor que determina o ícone
  
	// Definindo valores padrão para o ícone e a cor
	let IconComponent = IconFlag;
	let colorClass = 'text-gray-600';
  
	// Função que determina o ícone e a cor com base no tipo e valor
	$: {
	  if (type === 'Status') {
		switch (value) {
		  case 'Classificado':
			IconComponent = IconCheck;
			colorClass = 'text-success';
			break;
		  case 'Revisar':
			IconComponent = IconMessageCircle;
			colorClass = 'text-error';
			break;
		  case 'Pendente':
			IconComponent = IconClock;
			colorClass = 'text-info';
			break;
		  default:
			IconComponent = IconFlag;
			colorClass = 'text-gray-600';
		}
	  } else if (type === 'Prioridade') {
		switch (value?.toLowerCase().trim()) {
		  case 'alta':
			IconComponent = IconChevronsUp;
			colorClass = 'text-error';
			break;
		  case 'media':
			IconComponent = IconMinus;
			colorClass = 'text-info';
			break;
		  case 'baixa':
			IconComponent = IconChevronsDown;
			colorClass = 'text-success';
			break;
		  default:
			IconComponent = IconFlag;
			colorClass = 'text-gray-600';
		}
	  } else {
		IconComponent = IconFlag;
		colorClass = 'text-gray-600';
	  }
	}
  </script>
  
  <!-- Adicionando o HoverCard ao redor do ícone -->
  <HoverCard.Root openDelay={30} closeDelay={30}>
	<HoverCard.Trigger asChild>
	  <!-- O ícone é o Trigger do Hover -->
	  <div class="w-full flex items-center justify-center cursor-pointer">
		<IconComponent class={`size-5 2xl:size-7 ${colorClass}`} />
	  </div>
	</HoverCard.Trigger>
  
	<!-- Conteúdo do HoverCard exibindo o valor do status ou prioridade -->
	<HoverCard.Portal>
	  <HoverCard.Content sideOffset={5} align={"start"} class="p-4 bg-primary-content text-primary shadow rounded">
		<HoverCard.Arrow width={15} height={10} class="fill-primary-content" />
		<p class="text-base font-medium">{value}</p>
	  </HoverCard.Content>
	</HoverCard.Portal>
  </HoverCard.Root>
  
  