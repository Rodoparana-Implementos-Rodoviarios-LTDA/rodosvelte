<script lang="ts">
    // Definindo os tipos possíveis de "Tipo"
    export let tipo: string; // O valor do "Tipo" será passado como prop
    
    // Classe de cor padrão para o badge
    let badgeColor = "badge-neutral"; // Cor padrão caso o tipo não esteja mapeado
  
    // Atribuindo cores específicas com base no valor de "tipo"
    $: {
      switch (tipo?.toLowerCase().trim()) {
        case "despesa/imobilizado":
          badgeColor = "badge-primary";
          break;
        case "revenda":
          badgeColor = "badge-secondary";
          break;
        case "materia prima":
          badgeColor = "badge-accent";
          break;
        case "collection":
          badgeColor = "badge-info";
          break;
        default:
          badgeColor = "badge-neutral";
      }
    }
  </script>
  
  <!-- Estrutura do badge usando DaisyUI -->
  <span class={`badge badge-outline ${badgeColor}`}>
    {tipo ? tipo.charAt(0).toUpperCase() + tipo.slice(1).toLowerCase() : "N/A"}
  </span>
  
  <style>
    /* Ajustes de tamanho se necessário */
    .badge {
      padding: 0.25rem 0.5rem;
      font-size: 0.875rem; /* Tamanho pequeno do texto */
    }
  </style>
  