<script lang="ts">
  import flatpickr from 'flatpickr';
  import 'flatpickr/dist/flatpickr.min.css';
  import { onMount } from 'svelte';

  let dateInput: Node;
  let selectedDates = [];

  onMount(() => {
    flatpickr(dateInput, {
      mode: 'range',
      dateFormat: 'Y-m-d',
      onChange: (dates) => {
        selectedDates = dates;
        console.log('Datas selecionadas:', dates);
      }
    });
  });
</script>

<!-- Estilize o campo de input diretamente se necessário -->
<input bind:this={dateInput} class="input input-sm input-bordered" />

<style>
  .input {
    background-color: #fff;
    border: 1px solid #ccc;
    border-radius: 5px;
    padding: 8px;
    color: #333;
    font-size: 14px;
  }
</style>
