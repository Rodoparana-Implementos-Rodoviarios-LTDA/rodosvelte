import { defineConfig } from 'vite';
import { sveltekit } from '@sveltejs/kit/vite';

export default defineConfig({
  plugins: [sveltekit()],
  resolve: {
    alias: {
      $lib: '/src/lib',  // Certifique-se de ajustar para seus diretórios corretos
      $components: '/src/components'
    }
  }
});
