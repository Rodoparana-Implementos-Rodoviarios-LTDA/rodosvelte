<script lang="ts">
	interface ItemNF {
		D2_ITEM: string;
		D2_COD: string;
		B1_DESC: string;
		SALDO: number;
		quantity: number; // Campo para armazenar a quantidade inserida pelo usuário
	}

	// Recebendo as propriedades do componente pai
	export let documentoCompleto: string;
	export let clienteCompleto: string;
	export let filial: string;

	// Criando um array para armazenar os itens selecionados
	export let itensSelecionados: ItemNF[] = [];

	let errorMessage = '';

	// Função para separar documento e série
	function separarDocumentoESerie(docCompleto: string) {
		console.log('SepararDocumentoESerie - docCompleto:', docCompleto);
		if (!docCompleto) return { documento: '', serie: '' };

		const [documento, serie] = docCompleto.split(' - ').map((part) => part.trim());
		return { documento: documento || '', serie: serie || '' };
	}

	// Função para extrair código do cliente, loja e nome
	function separarClienteLoja(clienteCompleto: string) {
		console.log('SepararClienteLoja - clienteCompleto:', clienteCompleto);
		if (!clienteCompleto) return { cliente: '', loja: '', nomeCliente: '' };

		const partes = clienteCompleto.split(' - ').map((part) => part.trim());
		const cliente = partes[0] || '';

		let loja = '';
		let nomeClientePartes;

		if (partes[1] && /^\d+$/.test(partes[1])) {
			loja = partes[1];
			nomeClientePartes = partes.slice(2);
		} else {
			nomeClientePartes = partes.slice(1);
		}

		const nomeCliente = nomeClientePartes.join(' - ') || '';
		console.log('Cliente:', cliente, 'Loja:', loja, 'Nome do Cliente:', nomeCliente);
		return { cliente, loja, nomeCliente };
	}

	// Função para buscar os itens da API
	async function fetchItensNF() {
		const { documento, serie } = separarDocumentoESerie(documentoCompleto);
		const { cliente, loja } = separarClienteLoja(clienteCompleto);

		// Garantir que 'filial' seja de 4 caracteres numéricos
		const filialTrimmed = filial.trim();
		if (!/^\d{4}$/.test(filialTrimmed)) {
			errorMessage = `Filial inválida: Deve conter exatamente 4 dígitos numéricos. Valor atual: '${filialTrimmed}'`;
			console.error(errorMessage);
			return;
		}

		// Criação do objeto de parâmetros de filtragem
		const body = {
			filial: filialTrimmed,
			documento: documento.trim(),
			serie: serie.trim(),
			cliente: cliente.trim(),
			loja: loja.trim()
		};

		// Log dos valores antes de criar a query string
		console.log('Parâmetros antes da codificação:', body);

		// Criando os parâmetros de consulta da URL
		const queryParams = new URLSearchParams(paramsObj);
		const url = `http://protheus-vm:9010/rest/MovPortaria/ListaItensNF?${queryParams.toString()}`;

		// Log da URL completa
		console.log('URL de requisição:', url);

		try {
			const response = await fetch(url, {
				method: 'GET'
				// Removido o header 'Content-Type'
			});

			if (!response.ok) {
				const errorText = await response.text();
				console.error('Erro na resposta da API:', errorText);
				throw new Error(
					`Erro ao buscar itens: ${response.status} ${response.statusText} - ${errorText}`
				);
			}

			const data = await response.json();
			console.log('Dados recebidos da API:', data);

			// Inicializa a lista de itens com os dados retornados
			itensSelecionados = data.map((item: ItemNF) => ({ ...item, quantity: 0 })) || [];
		} catch (error) {
			console.error('Erro ao buscar itens da NF:', error);
			errorMessage = `Erro ao carregar itens da nota fiscal: ${error.message}`;
		}
	}
	// Chamando a função fetch ao montar o componente
	fetchItensNF();

	function atualizarItensSelecionados() {
		itensSelecionados = itensSelecionados.filter((item) => item.quantity > 0);
	}
</script>

<!-- Exibindo a lista de itens -->
<div>
	<h3 class="text-lg font-semibold text-neutral-content">Produtos da Nota Fiscal:</h3>
	{#if itensSelecionados.length > 0}
		<ul class="list-none">
			{#each itensSelecionados as item}
				<li class="mb-2">
					<div class="flex justify-between items-center">
						<div>
							<span>{item.D2_COD.trim()} - {item.B1_DESC.trim()} (Saldo: {item.SALDO})</span>
						</div>
						<div class="flex items-center space-x-2">
							<input
								type="number"
								min="0"
								max={item.SALDO}
								bind:value={item.quantity}
								on:input={atualizarItensSelecionados}
								class="input input-bordered w-24"
							/>
							<span>Qtd</span>
						</div>
					</div>
				</li>
			{/each}
		</ul>
	{:else}
		<p class="text-neutral-content">Nenhum item encontrado.</p>
	{/if}

	{#if errorMessage}
		<p class="text-error mt-4">{errorMessage}</p>
	{/if}
</div>
