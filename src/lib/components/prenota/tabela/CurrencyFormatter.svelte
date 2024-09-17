<script lang="ts">
    export let value: number;
  
    // Converte o valor numérico para formato de moeda brasileira
    const formattedValue = !isNaN(value)
      ? value.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })
      : 'R$ 0,00';
  </script>
  
  <span>{formattedValue}</span>
  