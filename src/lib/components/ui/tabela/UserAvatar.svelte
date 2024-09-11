<script lang="ts">
  import { HoverCard } from 'radix-svelte';
  
  // Propriedades recebidas
  export let username: string;

  // Função para pegar as iniciais a partir do username
  const getInitials = (username?: string) => {
    if (!username) return "??"; // Valor padrão se username for undefined
    const [firstName, lastName] = username.split(".");
    const firstInitial = firstName?.charAt(0).toUpperCase() || "";
    const lastInitial = lastName?.charAt(0).toUpperCase() || "";
    return `${firstInitial}${lastInitial}`;
  };

  // Função para formatar o nome completo
  const getFullName = (username?: string) => {
    if (!username) return "Nome Desconhecido"; // Valor padrão se username for undefined
    const [firstName, lastName] = username.split(".");
    const formattedFirstName = firstName ? capitalize(firstName) : "";
    const formattedLastName = lastName ? capitalize(lastName) : "";
    return `${formattedFirstName} ${formattedLastName}`.trim();
  };

  // Função para capitalizar a primeira letra de cada nome
  const capitalize = (name: string) => {
    return name.charAt(0).toUpperCase() + name.slice(1).toLowerCase();
  };
</script>

<!-- Componente Radix-Svelte HoverCard -->
<HoverCard.Root>
  <HoverCard.Trigger asChild>
    <!-- Avatar com as iniciais -->
    <div class="flex items-center justify-start">
      <div class="flex items-center justify-center bg-base-300 rounded-full h-9 w-9 2xl:h-10 2xl:w-10 text-xs 2xl:text-base">
        {getInitials(username)}
      </div>
    </div>
  </HoverCard.Trigger>

  <!-- Conteúdo do HoverCard exibindo o nome completo -->
  <HoverCard.Portal>
    <HoverCard.Content class="p-2 bg-base-300 shadow rounded">
      <p class="text-sm font-medium">{getFullName(username)}</p>
    </HoverCard.Content>
  </HoverCard.Portal>
</HoverCard.Root>
