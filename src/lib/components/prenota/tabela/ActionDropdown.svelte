<!-- ActionDropdown.svelte -->
<script>
	import { IconEye, IconHistory } from '@tabler/icons-svelte'; // Certifique-se de que os ícones corretos sejam importados
</script>

<div class="flex">
	<div class="drawer">
		<input id="Revisao" type="checkbox" class="drawer-toggle" />
		<div class="drawer-content">
			<!-- Page content here -->
			<label for="Revisao" class="btn btn-outline drawer-button z-0" aria-label="Abrir revisão">
				<IconEye class="inline-block h-5 w-5" />
			</label>
		</div>
		<div class="drawer-side z-[100]">
			<label for="Revisao" aria-label="Fechar revisão" class="drawer-overlay"></label>
			<ul class="menu bg-base-200 text-base-content min-h-full w-80 p-4">
				<!-- Sidebar content here -->
				<li><a href="#">Revisão Item 1</a></li>
				<li><a href="#">Revisão Item 2</a></li>
			</ul>
		</div>
	</div>

	<div class="drawer">
		<input id="historico" type="checkbox" class="drawer-toggle" />
		<div class="drawer-content ">
			<!-- Page content here -->
			<label for="historico" class="btn btn-outline drawer-button z-0" aria-label="Abrir histórico">
				<IconHistory class="inline-block h-5 w-5" />
			</label>
		</div>
		<div class="drawer-side z-[100]">
			<label for="historico" aria-label="Fechar histórico" class="drawer-overlay"></label>
			<ul class="menu bg-base-200 text-base-content min-h-full w-80 p-4">
				<!-- Sidebar content here -->
				<li><a href="#">Histórico Item 1</a></li>
				<li><a href="#">Histórico Item 2</a></li>
			</ul>
		</div>
	</div>
</div>
