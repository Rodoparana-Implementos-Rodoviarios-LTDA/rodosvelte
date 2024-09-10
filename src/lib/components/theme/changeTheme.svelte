<script lang="ts">
  import { BrightnessHalf } from 'tabler-icons-svelte'; // Ícone para o tema

  let lightThemes = [
    { name: 'winter', colors: ['bg-[#5e81ac]', 'bg-[#81a1c1]', 'bg-[#88c0d0]', 'bg-[#4c566a]'], bg: 'bg-[#ffffff]', rounded: 'rounded-xl', title: 'text-[#394e6a]' },
    { name: 'autumn', colors: ['bg-[#8c0327]', 'bg-[#d85251]', 'bg-[#d59b6a]', 'bg-[#1a1a1a]'], bg: 'bg-[#f1f1f1]', rounded: 'rounded-md', title: 'text-[#141414]' },
    { name: 'valentine', colors: ['bg-[#5e81ac]', 'bg-[#81a1c1]', 'bg-[#88c0d0]', 'bg-[#4c566a]'], bg: 'bg-[#fae7f4]', rounded: 'rounded-full', title: 'text-[#632c3b]' }
  ];

  let darkThemes = [
    { name: 'sunset', colors: ['bg-[#ff865b]', 'bg-[#fd6f9c]', 'bg-[#b387fa]', 'bg-[#826a5c]'], bg: 'bg-[#121c22]', rounded: 'rounded-lg', title: 'text-[#9fb9d0]' },
    { name: 'dracula', colors: ['bg-[#ff79c6]', 'bg-[#bd93f9]', 'bg-[#ffb86c]', 'bg-[#414558]'], bg: 'bg-[#282a36]', rounded: 'rounded-xl', title: 'text-[#f8f8f2]' },
    { name: 'forest', colors: ['bg-[#1eb854]', 'bg-[#1db88e]', 'bg-[#1db8ab]', 'bg-[#19362d]'], bg: 'bg-[#171212]', rounded: 'rounded-full', title: 'text-[#cbc9c9]' }
  ];

  let selectedTheme: string = 'winter'; // O tema padrão

  function changeTheme(theme: string): void {
    selectedTheme = theme;
    document.documentElement.setAttribute('data-theme', theme); // Aplica o tema no <html>
  }
</script>

<!-- Accordion para o ThemeChanger -->
<div class="collapse collapse-arrow  bg-base-300">
  <input type="checkbox" class="peer" />
  <div class="collapse-title font-medium uppercase tracking-wide text-sm text-secondary">
   Tema
  </div>

  <div class="collapse-content">
    <ul class="space-y-1">
      {#each lightThemes as theme}
        <li>
          <button
            class={`btn text-sm flex justify-between ${theme.rounded} ${theme.bg}`}
            on:click={() => changeTheme(theme.name)}
          >
            <span class={`${theme.title} font-medium`}>{theme.name}</span>
            <div class="flex space-x-1">
              {#each theme.colors as color}
                <div class={`rounded-full h-5 w-2 ${color}`}></div>
              {/each}
            </div>
          </button>
        </li>
      {/each}

      {#each darkThemes as theme}
        <li>
          <button
            class={`btn text-sm flex justify-between ${theme.rounded} ${theme.bg}`}
            on:click={() => changeTheme(theme.name)}
          >
            <span class={`font-medium ${theme.title}`}>{theme.name}</span>
            <div class="flex space-x-1">
              {#each theme.colors as color}
                <div class={`rounded-full h-5 w-2 ${color}`}></div>
              {/each}
            </div>
          </button>
        </li>
      {/each}
    </ul>
  </div>
</div>
