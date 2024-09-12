<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let open = false; // Controla a visibilidade do hover card
  const dispatch = createEventDispatcher();
  let cardId = 'hovercard-' + Math.random().toString(36).substr(2, 9); // Gerar um ID único

  function show() {
    open = true;
    dispatch('open');
  }

  function hide() {
    open = false;
    dispatch('close');
  }
</script>

<div
  class="relative inline-block"
  on:mouseenter={show}
  on:mouseleave={hide}
  on:focus={show}
  on:blur={hide}
  tabindex="0"
  aria-expanded={open}
  aria-controls={cardId}
  on:keydown={hide}
  role="button"
>
  <slot name="trigger"></slot>

  {#if open}
    <slot name="content" id={cardId}></slot>
  {/if}
</div>
