<script lang="ts">
  // Importando os ícones do Tabler Icons
  import {
    Check,
    MessageCircle,
    Clock,
    ChevronsUp,
    ChevronsDown,
    Minus,
    Flag,
  } from 'tabler-icons-svelte';

  export let type: 'Status' | 'Prioridade'; // Tipos possíveis
  export let value: string; // O valor que determina o ícone

  // Definindo valores padrão para o ícone e a cor
  let IconComponent = Flag;
  let colorClass = 'text-gray-600';

  // Função que determina o ícone e a cor com base no tipo e valor
  $: {
    if (type === 'Status') {
      switch (value) {
        case 'Classificado':
          IconComponent = Check;
          colorClass = 'text-success';
          break;
        case 'Revisar':
          IconComponent = MessageCircle;
          colorClass = 'text-error';
          break;
        case 'Pendente':
          IconComponent = Clock;
          colorClass = 'text-info';
          break;
        default:
          IconComponent = Flag;
          colorClass = 'text-gray-600';
      }
    } else if (type === 'Prioridade') {
      switch (value?.toLowerCase().trim()) {
        case 'alta':
          IconComponent = ChevronsUp;
          colorClass = 'text-error';
          break;
        case 'media':
          IconComponent = Minus;
          colorClass = 'text-info';
          break;
        case 'baixa':
          IconComponent = ChevronsDown;
          colorClass = 'text-success';
          break;
        default:
          IconComponent = Flag;
          colorClass = 'text-gray-600';
      }
    } else {
      IconComponent = Flag;
      colorClass = 'text-gray-600';
    }
  }
</script>

<div class="w-full flex items-center justify-center">
  <IconComponent class={`size-5 2xl:size-7 ${colorClass}`} />
</div>

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
