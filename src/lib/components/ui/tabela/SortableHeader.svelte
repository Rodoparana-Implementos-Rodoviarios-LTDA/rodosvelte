<script lang="ts">
  import { ChevronsUp, ChevronsDown, Selector } from 'tabler-icons-svelte';

  export let title: string;
  export let isSorted: 'asc' | 'desc' | false | undefined = false;
  export let onClick: () => void;

  let IconComponent = Selector;

  $: {
    if (isSorted === 'asc') {
      IconComponent = ChevronsUp;
    } else if (isSorted === 'desc') {
      IconComponent = ChevronsDown;
    } else {
      IconComponent = Selector;
    }
  }
</script>

<button
  on:click={onClick}
  class="flex items-center justify-center text-center space-x-2 hover:text-gray-700 focus:outline-none"
>
  <span  class="text-sm ">{title}</span>
  <IconComponent class="h-4 w-4" />
</button>
