<script lang="ts">
	import { createEventDispatcher } from 'svelte';
  
	// Estados para os filtros
	export let selectedFilial = '';
	export let selectedNF = '';
	export let selectedCliente = '';
	export let selectedVendedor = '';
	export let selectedProduto = '';
	export let selectedSaldo = '';
	export let selectedEmissao = '';
  
	const dispatch = createEventDispatcher();
  
	// Função para aplicar os filtros
	function applyFilters() {
	  dispatch('applyFilters', {
		selectedFilial,
		selectedNF,
		selectedCliente,
		selectedVendedor,
		selectedProduto,
		selectedSaldo,
		selectedEmissao
	  });
	}
  
	// Função para limpar os filtros
	function resetFilters() {
	  selectedFilial = '';
	  selectedNF = '';
	  selectedCliente = '';
	  selectedVendedor = '';
	  selectedProduto = '';
	  selectedSaldo = '';
	  selectedEmissao = '';
	  applyFilters(); // Aplica os filtros após limpar
	}
  </script>
  
  <!-- Menu de Filtros dentro do Drawer -->
  <div class="menu p-4 w-80 bg-base-100 text-base-content h-screen flex flex-col justify-between" role="form" aria-labelledby="filters-heading">
	<div>
	  <h2 id="filters-heading" class="text-lg font-bold mb-4">Filtros Borracharia</h2>
  
	  <!-- Filtro de Filial -->
	  <div class="form-control mb-4">
		<label class="label" for="filial-input"><span class="label-text">Filial</span></label>
		<input
		  id="filial-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedFilial}
		  placeholder="Filial"
		  aria-label="Filtrar por Filial"
		/>
	  </div>
  
	  <!-- Filtro de NF -->
	  <div class="form-control mb-4">
		<label class="label" for="nf-input"><span class="label-text">NF</span></label>
		<input
		  id="nf-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedNF}
		  placeholder="Nota Fiscal"
		  aria-label="Filtrar por Nota Fiscal"
		/>
	  </div>
  
	  <!-- Filtro de Cliente -->
	  <div class="form-control mb-4">
		<label class="label" for="cliente-input"><span class="label-text">Cliente</span></label>
		<input
		  id="cliente-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedCliente}
		  placeholder="Cliente"
		  aria-label="Filtrar por Cliente"
		/>
	  </div>
  
	  <!-- Filtro de Vendedor -->
	  <div class="form-control mb-4">
		<label class="label" for="vendedor-input"><span class="label-text">Vendedor</span></label>
		<input
		  id="vendedor-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedVendedor}
		  placeholder="Vendedor"
		  aria-label="Filtrar por Vendedor"
		/>
	  </div>
  
	  <!-- Filtro de Produto -->
	  <div class="form-control mb-4">
		<label class="label" for="produto-input"><span class="label-text">Produto</span></label>
		<input
		  id="produto-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedProduto}
		  placeholder="Produto"
		  aria-label="Filtrar por Produto"
		/>
	  </div>
  
	  <!-- Filtro de Saldo -->
	  <div class="form-control mb-4">
		<label class="label" for="saldo-input"><span class="label-text">Saldo</span></label>
		<input
		  id="saldo-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedSaldo}
		  placeholder="Saldo"
		  aria-label="Filtrar por Saldo"
		/>
	  </div>
  
	  <!-- Filtro de Data de Emissão -->
	  <div class="form-control mb-4">
		<label class="label" for="emissao-input"><span class="label-text">Data de Emissão</span></label>
		<input
		  id="emissao-input"
		  class="input input-sm input-bordered"
		  type="text"
		  bind:value={selectedEmissao}
		  placeholder="Data de Emissão"
		  aria-label="Filtrar por Data de Emissão"
		/>
	  </div>
	</div>
  
	<!-- Botões de Aplicar e Limpar Filtros -->
	<div class="flex justify-between">
	  <button class="btn btn-sm btn-primary" on:click={applyFilters} aria-label="Aplicar Filtros">Aplicar Filtros</button>
	  <button class="btn btn-sm btn-outline" on:click={resetFilters} aria-label="Limpar Filtros">Limpar Filtros</button>
	</div>
  </div>
  