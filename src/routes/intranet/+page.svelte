<script lang="ts">
    import NavIntranet from '$lib/components/theme/navIntranet.svelte'; // Importar o componente de navegação
    import { onMount } from 'svelte'; // Para obter a data atual
  
    let username = 'Usuário'; // Nome padrão do usuário
    let currentDate = ''; // Variável para armazenar a data atual
  
    // Função que vai ser executada quando o componente for montado
    onMount(() => {
      // Definir a data atual
      const today = new Date();
      const options = { year: 'numeric', month: 'long', day: 'numeric' };
      currentDate = today.toLocaleDateString('pt-BR', options);
    });
  </script>
  
  <style>
    .welcome-box {
      padding: 20px;
      background-color: #f3f4f6; /* Cor de fundo suave */
      border-radius: 8px;
      text-align: center;
      margin: 20px 0;
    }
  
    .welcome-box h1 {
      font-size: 2rem;
      color: #111827; /* Texto em cinza escuro */
    }
  
    .welcome-box .username {
      font-weight: bold;
      color: #1d4ed8; /* Azul (similar ao tema primário) */
    }
  
    .welcome-box .date {
      font-size: 1rem;
      color: #9ca3af; /* Cinza claro (estilo placeholder) */
    }
  </style>
  
  <!-- Usar o componente NavIntranet para a barra de navegação -->
  <NavIntranet>
    <!-- Caixa de boas-vindas -->
    <div class="welcome-box">
      <h1>Bem-vindo(a), <span class="username">{username}</span>!</h1>
      <p class="date">{currentDate}</p> <!-- Data com estilo de placeholder -->
    </div>
  </NavIntranet>
  