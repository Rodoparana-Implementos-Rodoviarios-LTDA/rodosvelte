<script lang="ts">
  export let label: string;   // O texto que será exibido no badge
  export let category: 'tipo' | 'status' | 'prioridade';  // A categoria que determina as cores do badge

  let badgeClass = 'badge';  // Classe base para o badge

  // Define as cores do badge com base na categoria e valor
  $: if (category === 'tipo') {
    switch (label) {
      case 'Revenda':
        badgeClass += ' badge-primary';
        break;
      case 'Matéria Prima':
        badgeClass += ' badge-secondary';
        break;
      case 'Collection':
        badgeClass += ' badge-accent';
        break;
      case 'Despesa/Imobilizado':
        badgeClass += ' badge-info';
        break;
      default:
        badgeClass += ' badge-neutral';
    }
  } else if (category === 'status') {
    switch (label) {
      case 'Classificado':
        badgeClass += ' badge-success';
        break;
      case 'Pendente':
        badgeClass += ' badge-warning';
        break;
      case 'Revisar':
        badgeClass += ' badge-error';
        break;
      default:
        badgeClass += ' badge-neutral';
    }
  } else if (category === 'prioridade') {
    switch (label) {
      case 'Baixa':
        badgeClass += ' badge-success';
        break;
      case 'Média':
        badgeClass += ' badge-warning';
        break;
      case 'Alta':
        badgeClass += ' badge-error';
        break;
      default:
        badgeClass += ' badge-neutral';
    }
  }
</script>

<span class={badgeClass}>{label}</span>

<style>
  .badge {
    padding: 0.25rem 0.5rem;
    border-radius: 0.25rem;
    font-size: 0.875rem;
    font-weight: 500;
    display: inline-block;
  }
</style>
