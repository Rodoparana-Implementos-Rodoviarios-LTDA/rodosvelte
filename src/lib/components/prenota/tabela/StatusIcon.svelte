<script lang="ts">
	import { HoverCard } from 'radix-svelte';

  // Importando os ícones do Tabler Icons
  import { Check, MessageCircle, Clock, ChevronsUp, ChevronsDown, Minus, Flag, Box, Fold, Shirt, BuildingWarehouse, BuildingStore, Mountain } from 'tabler-icons-svelte';

  export let type: "Status" | "Prioridade" | "Tipo";  // Tipos possíveis
  export let value: string;  // O valor que determina o ícone

  // Definindo valores padrão para o ícone e a cor
  let IconComponent = Flag;
  let colorClass = "text-gray-600";

  // Função que determina o ícone e a cor com base no tipo e valor
  $: {
    if (type === "Status") {
      switch (value) {
        case "Classificado":
          IconComponent = Check;
          colorClass = "text-primary";
          break;
        case "Revisar":
          IconComponent = MessageCircle;
          colorClass = "text-secondary";
          break;
        case "Pendente":
          IconComponent = Clock;
          colorClass = "text-accent";
          break;
        default:
          IconComponent = Flag;
          colorClass = "text-gray-600";
      }
    } else if (type === "Prioridade") {
      switch (value?.toLowerCase().trim()) {
        case "alta":
          IconComponent = ChevronsUp;
          colorClass = "text-primary";
          break;
        case "media":
          IconComponent = Minus;
          colorClass = "text-secondary";
          break;
        case "baixa":
          IconComponent = ChevronsDown;
          colorClass = "text-accent";
          break;
        default:
          IconComponent = Flag;
          colorClass = "text-gray-600";
      }
    } else if (type === "Tipo") {
      switch (value) {
        case "Despesa/Imobilizado":
          IconComponent = BuildingWarehouse;
          colorClass = "text-primary";
          break;
        case "Revenda":
          IconComponent = BuildingStore;
          colorClass = "text-secondary";
          break;
        case "Materia Prima":
          IconComponent = Mountain;
          colorClass = "text-accent";
          break;
        case "Collection":
          IconComponent = Shirt;
          colorClass = "text-teal-500";
          break;
        default:
          IconComponent = Flag;
          colorClass = "text-gray-600";
      }
    } else {
      IconComponent = Flag;
      colorClass = "text-gray-600";
    }
  }
</script>

<!-- Radix HoverCard para exibir o ícone e descrição -->
<HoverCard.Root>
	<HoverCard.Trigger>
		<div class="w-full flex items-center justify-start">
			<IconComponent class={`size-5 2xl:size-7 ${colorClass}`} />
		</div>
	</HoverCard.Trigger>

	<HoverCard.Portal>
		<HoverCard.Content class="p-2 bg-base-300 shadow-md rounded-md">
			<p>{value ? value.charAt(0).toUpperCase() + value.slice(1) : 'N/A'}</p>
		</HoverCard.Content>
	</HoverCard.Portal>
</HoverCard.Root>

<style>
	.size-5 {
		width: 1.25rem;
		height: 1.25rem;
	}
	.size-7 {
		width: 1.75rem;
		height: 1.75rem;
	}
</style>
