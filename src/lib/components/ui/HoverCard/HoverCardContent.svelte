<script lang="ts">
    export let position: 'top' | 'bottom' | 'left' | 'right' = 'top';
    export let id: string;
  </script>
  
  <div
    class="absolute z-50 bg-base-100 shadow-lg border border-base-300 p-4 rounded-lg text-sm text-base-content transition-all duration-200 mb-2"
    class:bottom={position === 'top'}
    class:top={position === 'bottom'}
    class:left={position === 'right'}
    class:right={position === 'left'}
    tabindex="-1"
    aria-labelledby={id}
    role="dialog"
  >
    <slot></slot>
  </div>
  