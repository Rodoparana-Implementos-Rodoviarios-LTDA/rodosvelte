<script lang="ts">
  import { onMount } from 'svelte';
  import { IconBrightnessHalf } from '@tabler/icons-svelte';
  import {lightThemes, darkThemes} from './themes'
  let selectedTheme: string = 'light'; // Tema padrão

  // Função para alterar o tema e salvar no localStorage
  function changeTheme(theme: string): void {
    selectedTheme = theme;
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('selectedTheme', theme); // Armazena o tema no localStorage
  }

  // Carrega o tema salvo no localStorage ao montar o componente
  onMount(() => {
    const savedTheme = localStorage.getItem('selectedTheme');
    if (savedTheme) {
      selectedTheme = savedTheme;
      document.documentElement.setAttribute('data-theme', savedTheme);
    }
  });
</script>

<!-- Dropdown para troca de temas -->
<div class="dropdown dropdown-end fixed top-5 right-5 z-0">
  <!-- Botão que abre o dropdown -->
  <div tabindex="0" role="button" class="btn m-1">
    <IconBrightnessHalf/> 
  </div>

  <!-- Conteúdo do dropdown para temas claros e escuros -->
  <ul tabindex="0" class="dropdown-content bg-base-300 flex justify-between rounded-box z-10 h-80 overflow-y-auto w-96 p-2 shadow-2xl">
    <div>
    <li class="font-bold text-sm">Light Themes</li>
    {#each lightThemes as theme}
      <li class="pb-2">
        <button class={`btn text-sm flex justify-between w-full ${theme.rounded} ${theme.bg}`} on:click={() => changeTheme(theme.name)}>
          <span class={`${theme.title} font-medium`}>{theme.name}</span>
          <div class="flex space-x-1">
            {#each theme.colors as color}
              <div class={`rounded-full h-5 w-2 ${color}`}></div>
            {/each}
          </div>
        </button>
      </li>
    {/each}
  </div>
  <div>
    <li class="font-bold text-sm">Dark Themes</li>
    {#each darkThemes as theme}
      <li class="pb-2">
        <button class={`btn text-sm flex justify-between w-full ${theme.rounded} ${theme.bg}`} on:click={() => changeTheme(theme.name)}>
          <span class={`font-medium ${theme.title}`}>{theme.name}</span>
          <div class="flex space-x-1">
            {#each theme.colors as color}
              <div class={`rounded-full h-5 w-2 ${color}`}></div>
            {/each}
          </div>
        </button>
      </li>
    {/each}
  </div>
  </ul>
</div>
