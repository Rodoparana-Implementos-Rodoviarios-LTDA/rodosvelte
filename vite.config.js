import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
  plugins: [sveltekit()],
  server: {
    proxy: {
      // Requisições para o servidor Go rodando em rodoapp:8080
      '/api': {
        target: 'http://rodoapp:8080', // Backend Go principal
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '') // Remove o "/api" da rota antes de enviar ao backend Go
      },

    }
  },
  test: {
    include: ['src/**/*.{test,spec}.{js,ts}']
  }
});
