<script lang="ts">
    import HoverCard from '$lib/components/ui/HoverCard/HoverCard.svelte';
    import HoverCardTrigger from '$lib/components/ui/HoverCard/HoverCardTrigger.svelte';
    import HoverCardContent from '$lib/components/ui/HoverCard/HoverCardContent.svelte';
  
    export let fornecedorInfo: string;
  
    // Função para dividir a string em Código, Filial e Nome
    function parseFornecedorInfo(info: string) {
      const regex = /^(\d+)\s+(\d+)\s+-\s+(.+)$/;
      const match = info.match(regex);
      if (match) {
        const [_, cod, filial, nome] = match;
        return { cod, filial, nome };
      }
      return { cod: "", filial: "", nome: info };
    }
  
    const { cod, filial, nome } = parseFornecedorInfo(fornecedorInfo);
  
    // Gerar um ID único para o HoverCardContent
    let hoverCardId = 'hovercard-' + Math.random().toString(36).substr(2, 9);
  </script>
  
  <HoverCard>
    <HoverCardTrigger slot="trigger">
      <span class="cursor-pointer underline text-primary">{nome}</span>
    </HoverCardTrigger>
  
    <!-- Passando o id gerado para HoverCardContent -->
    <HoverCardContent slot="content" position="top" id={hoverCardId}>
      <div class="card w-64 bg-base-100 shadow-xl p-4 text-sm">
        <p><strong>Cód:</strong> {cod}</p>
        <p><strong>Filial:</strong> {filial}</p>
        <p><strong>Nome:</strong> {nome}</p>
      </div>
    </HoverCardContent>
  </HoverCard>
  