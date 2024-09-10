/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,svelte,js,ts}'],
	theme: {
		extend: {}
	},
	plugins: [require('daisyui')],
	daisyui: {
		themes: ['winter', 'autumn', 'valentine', 'sunset', 'forest', 'dracula'] // Temas escolhidos
	}
};
