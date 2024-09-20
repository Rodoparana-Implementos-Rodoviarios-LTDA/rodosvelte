<script lang="ts">
  import { IconChevronsUp, IconChevronsDown, IconSelector } from '@tabler/icons-svelte';

  export let title: string;
  export let isSorted: 'asc' | 'desc' | false | undefined = false;
  export let onClick: () => void;

  let IconComponent = IconSelector;

  // Altera o ícone conforme a ordenação
  $: {
    if (isSorted === 'asc') {
      IconComponent = IconChevronsUp;
    } else if (isSorted === 'desc') {
      IconComponent = IconChevronsDown;
    } else {
      IconComponent = IconSelector;
    }
  }
</script>

<!-- Botão com centralização -->
<button
  on:click={onClick}
  class="h-full flex items-center justify-center text-center space-x-2 hover:text-gray-700 focus:outline-none"
>
  <!-- Centraliza o título e o ícone -->
  <span class="text-sm">{title}</span>
  <IconComponent class="h-4 w-4" />
</button>
