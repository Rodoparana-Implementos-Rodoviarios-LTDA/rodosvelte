<script lang="ts">
    import { HoverCard } from 'radix-svelte';
  
    export let text: string;  // O texto a ser exibido
  
    // Verifica se o texto é muito longo para exibição na célula
    const isLongText = text.length > 20;  // Ajuste esse valor conforme necessário
  </script>
  
  <!-- Condicionalmente exibe o HoverCard se o texto for longo -->
  {#if isLongText}
    <HoverCard.Root>
      <HoverCard.Trigger asChild>
        <span class="truncate w-full inline-block cursor-pointer">
          {text.slice(0, 20)}...  <!-- Exibe apenas uma parte do texto -->
        </span>
      </HoverCard.Trigger>
  
      <HoverCard.Portal>
        <HoverCard.Content class="p-2 bg-base-300 shadow-md rounded-md max-w-xs">
          <p>{text}</p>  <!-- Exibe o texto completo dentro do HoverCard -->
        </HoverCard.Content>
      </HoverCard.Portal>
    </HoverCard.Root>
  {:else}
    <span>{text}</span>  <!-- Se o texto não for longo, exibe-o normalmente -->
  {/if}
  
  <style>
    .truncate {
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
    }
  </style>
  