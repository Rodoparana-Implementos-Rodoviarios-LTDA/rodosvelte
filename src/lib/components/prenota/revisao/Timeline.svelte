<script lang="ts">
	import {
	  IconCirclePlusFilled,
	  IconCircleCheckFilled,
	  IconSettingsFilled,
	  IconExclamationCircleFilled,
	  IconClockFilled
	} from '@tabler/icons-svelte';
	import { getData, saveData } from '$lib/services/idb';
	import type { Revisao } from '$lib/types/revisao';
  
	// Recebe 'rec' e 'status' como props
	export let rec: string;
	export let status: string;
  
	let eventos: Revisao[] = []; // Lista de eventos
	let criacao: Revisao | null = null; // Dados de criação (campo == 'XX')
  
	// Função para carregar eventos da API diretamente no componente
	async function fetchRevisaoData() {
	  try {
		// Verifica se os dados já estão armazenados no IndexedDB para o 'rec' correto
		const cachedData = await getData(rec);
  
		if (cachedData) {
		  // Se já existir, carrega os dados do IndexedDB
		  eventos = cachedData;
		  console.log(`Dados encontrados no IndexedDB para o rec ${rec}:`, eventos);
		  processEventos(eventos);
		} else {
		  // Se não, faz a consulta e salva no IndexedDB
		  const response = await fetch('http://protheus-app:8400/rest/reidoapsdu/getHistorico', {
			method: 'GET',
			headers: {
			  RECSF1: rec, // Passa o 'rec' diretamente como header
			  'Content-Type': 'application/json'
			}
		  });
  
		  if (!response.ok) {
			throw new Error('Erro ao buscar histórico');
		  }
  
		  const data: Revisao[] = await response.json();
		  eventos = data;
  
		  // Salva os dados no IndexedDB
		  await saveData(rec, eventos);
  
		  console.log(`Dados obtidos da API para o rec ${rec}:`, eventos);
		  processEventos(eventos);
		}
	  } catch (error) {
		console.error('Erro ao buscar eventos:', error);
	  }
	}
  
	// Processa os eventos e separa criação e revisões
	function processEventos(data: Revisao[]) {
	  eventos = data;
  
	  // Procura o evento de criação (campo == 'XX')
	  criacao = eventos.find((event) => event.campo === 'XX') || null;
	}
  
	// Carrega os eventos ao montar o componente
	$: if (rec) {
	  fetchRevisaoData();
	}
  
	// Função para verificar se o evento é de revisão
	function isRevisao(event: Revisao) {
	  return event.campo === 'F1_ZOBSREV';
	}
  
	// Função para verificar se o status é "Classificado"
	function isClassificado() {
	  return status === 'Classificado';
	}
  </script>
  
  <div class="h-full">
	<h2 class="text-3xl text-start font-bold pb-10">Histórico de Eventos</h2>
  
	<ul class="timeline timeline-vertical">
	  <!-- Primeiro Step - Evento de Criação -->
	  {#if criacao}
		<li>
		  <div class="timeline-start timeline-box text-end">
			<span>{criacao.data} </span><br />
			<span>{criacao.hora}</span>
		  </div>
		  <div class="timeline-middle">
			<IconCirclePlusFilled class="text-info h-8 w-8" />
		  </div>
		  <div class="timeline-end timeline-box">
			<span class="font-bold">Criado por:</span> {criacao.usuario}
		  </div>
		  <hr class="bg-info" />
		</li>
	  {/if}
  
	  <!-- Itera sobre todos os eventos, exceto o primeiro -->
	  {#each eventos.slice(1) as event, index}
		<li>
		  <hr class={isRevisao(event) ? 'bg-error' : 'bg-neutral'} />
		  <div class="timeline-start">
			<span>{event.data} </span><br />
			<span>{event.hora}</span>
		  </div>
		  <div class="timeline-middle">
			{#if isRevisao(event)}
			  <IconExclamationCircleFilled class="text-error h-8 w-8" />
			{:else}
			  <IconSettingsFilled class="h-8 w-8" />
			{/if}
		  </div>
		  <div class="timeline-end timeline-box">
			{#if isRevisao(event)}
			  Revisado por {event.usuario}: {event.atual}
			{:else}
			  Alterado por {event.usuario}: {event.atual}
			{/if}
		  </div>
		  <hr class={isRevisao(event) ? 'bg-error' : 'bg-neutral'} />
		</li>
	  {/each}
  
	  <!-- Último Step - Classificação -->
	  <li>
		<hr class={isClassificado() ? 'bg-info' : 'bg-neutral'} />
		<div class="timeline-start"></div>
		<div class="timeline-middle">
		  {#if isClassificado()}
			<IconCircleCheckFilled class="text-info h-8 w-8" />
		  {:else}
			<IconClockFilled class="h-8 w-8 text-neutral" />
		  {/if}
		</div>
		<div class="timeline-end timeline-box">
		  {isClassificado() ? 'Classificado' : status}
		</div>
	  </li>
	</ul>
  </div>
  