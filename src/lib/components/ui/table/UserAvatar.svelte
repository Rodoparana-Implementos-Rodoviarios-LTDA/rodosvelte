<script lang="ts">
    export let name: string; // Nome completo do usuário
    export let imageUrl: string | null = null; // URL da imagem do perfil (opcional)
  
    // Função para gerar as iniciais a partir do nome completo
    function getInitials(name: string): string {
      const nameParts = name.split(' ');
      if (nameParts.length > 1) {
        return nameParts[0][0] + nameParts[1][0]; // Primeira letra do nome e do sobrenome
      } else {
        return nameParts[0][0]; // Caso tenha apenas um nome
      }
    }
  
    // Iniciais geradas a partir do nome
    const initials = getInitials(name);
  </script>
  
  <style>
    .avatar {
      width: 40px;
      height: 40px;
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 50%;
      background-color: #ccc;
      color: white;
      font-weight: bold;
      font-size: 16px;
      overflow: hidden;
    }
    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      object-fit: cover;
    }
  </style>
  
  <div class="avatar" title={name}>
    <!-- Mostra a imagem do usuário, se disponível, caso contrário, mostra as iniciais -->
    {#if imageUrl}
      <img src={imageUrl} alt={name} />
    {:else}
      {initials}
    {/if}
  </div>
  