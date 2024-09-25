<script lang="ts">
	import { HoverCard } from 'radix-svelte';
  
	// Propriedades recebidas
	export let username: string | undefined;
  
	// Função para pegar as iniciais a partir do username
	const getInitials = (username?: string) => {
	  if (!username) return '??'; // Valor padrão se username for undefined
	  const nameParts = username.split(' ').filter(part => isNaN(Number(part))); // Remove números
	  const firstInitial = nameParts[0]?.charAt(0).toUpperCase() || '';  // Pega a primeira letra do primeiro nome
	  const lastInitial = nameParts.length > 1 ? nameParts[1].charAt(0).toUpperCase() : ''; // Pega a primeira letra do sobrenome
	  return `${firstInitial}${lastInitial}`;
	};
  
	// Função para formatar o nome completo
	const getFullName = (username?: string) => {
	  if (!username) return 'Nome Desconhecido'; // Valor padrão se username for undefined
	  return username.split(' ').map(capitalize).join(' '); // Capitaliza cada parte do nome
	};
  
	// Função para capitalizar a primeira letra de cada palavra
	const capitalize = (name: string) => {
	  return name.charAt(0).toUpperCase() + name.slice(1).toLowerCase();
	};
  </script>
  
  <!-- Componente Radix-Svelte HoverCard -->
  <HoverCard.Root openDelay={30} closeDelay={30}>
	<HoverCard.Trigger asChild>
	  <!-- Avatar com as iniciais -->
	  <div class="flex items-center justify-center">
		<div
		  class="flex items-center justify-center bg-primary-content text-primary font-medium rounded-full h-9 w-9 2xl:h-12 2xl:w-12 text-xs 2xl:text-lg"
		>
		  {getInitials(username)}
		</div>
	  </div>
	</HoverCard.Trigger>
  
	<!-- Conteúdo do HoverCard exibindo o nome completo -->
	<HoverCard.Portal>
	  <HoverCard.Content sideOffset={5} align={"start"} class="p-4 bg-primary-content text-primary shadow rounded">
		<HoverCard.Arrow width={15} height={10} class="fill-primary-content" />
		<p class="text-base font-medium">{getFullName(username)}</p>
	  </HoverCard.Content>
	</HoverCard.Portal>
  </HoverCard.Root>
  