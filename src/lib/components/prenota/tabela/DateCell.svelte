<script lang="ts">
  import { HoverCard } from 'radix-svelte';
  import { differenceInDays, parse } from 'date-fns';
	import { onMount } from 'svelte';

  // Propriedades
  export let rawDateInclusion: string;
  export let rawDateEmission: string;
  export let rawDateDue: string;

  // Controlando as classes e os cálculos
  let differenceInDaysValueInclusion = 0;
  let differenceInDaysValueDue = 0;
  let inclusionTextClass = 'text-lime-600';
  let dueTextClass = 'text-lime-600';

  // Configurações do HoverCard
  let contentSide = "bottom";
  let contentSideOffset = 10;
  let contentAlign = "center";
  let contentArrowPadding = 5;
  let contentAvoidCollisions = true;

  // Cálculo da diferença de dias e ajuste das classes
  onMount(() => {
    try {
      const parsedInclusionDate = parse(rawDateInclusion, "dd-MM-yyyy", new Date());
      const parsedEmissionDate = parse(rawDateEmission, "dd-MM-yyyy", new Date());
      const parsedDueDate = parse(rawDateDue, "dd-MM-yyyy", new Date());
      const currentDate = new Date();

      differenceInDaysValueInclusion = differenceInDays(parsedInclusionDate, parsedEmissionDate);
      differenceInDaysValueDue = differenceInDays(parsedDueDate, currentDate);

      // Ajuste das classes de acordo com a diferença de dias para inclusão
      if (differenceInDaysValueInclusion > 10) {
        inclusionTextClass = 'text-rose-600';
      } else if (differenceInDaysValueInclusion > 5) {
        inclusionTextClass = 'text-amber-600';
      }

      // Ajuste das classes de acordo com a diferença de dias para vencimento
      if (differenceInDaysValueDue < 0) {
        dueTextClass = 'text-rose-600';
      } else if (differenceInDaysValueDue < 2) {
        dueTextClass = 'text-amber-600';
      } else if (differenceInDaysValueDue > 3) {
        dueTextClass = 'text-lime-600';
      }
    } catch (e) {
      console.error("Erro ao parsear datas:", e);
    }
  });
</script>

<!-- Radix-Svelte HoverCard -->
<div class="contents">
  <HoverCard.Root>
    <HoverCard.Trigger class="cursor-pointer">
      <span class={dueTextClass}>
        {rawDateDue}
      </span>
    </HoverCard.Trigger>
    
    <HoverCard.Portal>
      <HoverCard.Content
        side={contentSide}
        sideOffset={contentSideOffset}
        align={contentAlign}
        arrowPadding={contentArrowPadding}
        avoidCollisions={contentAvoidCollisions}
        class="w-64 p-2 bg-base-300 shadow rounded"
      >
        <HoverCard.Arrow width={10} height={5} class="fill-white" />
        <div class="flex flex-col gap-2 ">
          <p class="flex justify-between">Inclusão: <span class={inclusionTextClass}>{rawDateInclusion}</span></p>
          <p class="flex justify-between">Emissão: <span class="text-gray-600">{rawDateEmission}</span></p>
          <p class="flex justify-between">Vencimento: <span class={dueTextClass}>{rawDateDue}</span></p>
        </div>
      </HoverCard.Content>
    </HoverCard.Portal>
  </HoverCard.Root>
</div>

<style>
  .contents {
    display: inline-block;
  }
</style>
