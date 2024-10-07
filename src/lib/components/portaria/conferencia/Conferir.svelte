<script lang="ts">
    import { onMount } from 'svelte';
    import { IconClipboard, IconTagsFilled, IconUser } from '@tabler/icons-svelte';
    import ButtonConferir from './buttonConferir.svelte'; // Importa o componente ButtonConferir

    export let documento: string;
    export let produto: string;
    export let saldoConferencia: number;
    export let responsavel: string;
    export let filial: string;
    export let cliente: string;

    let saldoMaximo: number | null = null;
    let seqNumero: string | null = null; // Variável para armazenar o Seq
    let isLoading = true;
    let errorMessageFetch = '';
    let isCheckboxChecked = false; // Estado do checkbox

    // Função para buscar saldo máximo utilizando a API
    async function fetchSaldoMaximo() {
        console.log('Iniciando fetchSaldoMaximo');

        const url = `http://rodoapp:8080/api/portaria`;

        try {
            const response = await fetch(url, {
                method: 'GET',
                headers: {
                    'X-Filter-Filial': filial.trim()
                }
            });

            if (!response.ok) {
                const errorText = await response.text();
                throw new Error(`Erro ao buscar saldo máximo: ${response.status} ${errorText}`);
            }

            const data = await response.json();
            console.log('Dados recebidos da API:', data);

            const itemCorrespondente = data.find((item: any) =>
                item.NF?.trim() === documento.trim() && item.Produto?.trim() === produto.trim());

            if (itemCorrespondente) {
                saldoMaximo = Number(itemCorrespondente.Saldo);
                seqNumero = itemCorrespondente.Seq?.trim() || null; // Armazenar o Seq
                console.log('Seq encontrado:', seqNumero);
            } else {
                saldoMaximo = null;
            }
        } catch (error: any) {
            errorMessageFetch = `Erro ao carregar saldo máximo: ${error.message}`;
            saldoMaximo = null;
        } finally {
            isLoading = false;
        }
    }

    onMount(() => {
        fetchSaldoMaximo();
    });
</script>

<!-- Conteúdo principal -->
<div class="w-full h-full bg-base-100 text-base-content p-6 rounded-lg shadow-lg">
    <h2 class="text-2xl font-bold mb-6 text-primary">Conferência de Produto</h2>

    <!-- Informações do Documento -->
    <p class="mb-4">
        <span class="inline-flex items-center">
            <IconClipboard class="mr-2 text-primary" />
            <strong class="text-primary text-lg">Documento:</strong>
        </span>
        <br />
        <span class="text-base-content text-lg">{documento}</span>
    </p>

    <!-- Informações do Produto -->
    <p class="mb-4">
        <span class="inline-flex items-center">
            <IconTagsFilled class="mr-2 text-primary" />
            <strong class="text-primary text-lg">Produto:</strong>
        </span>
        <br />
        <span class="text-base-content text-lg">{produto}</span>
    </p>

    <!-- Retirado por -->
    <p class="mb-4">
        <span class="inline-flex items-center">
            <IconUser class="mr-2 text-primary" />
            <strong class="text-primary text-lg">Retirado por:</strong>
        </span>
        <br />
        <span class="text-base-content text-lg">{responsavel}</span>
    </p>

    <!-- Comparação entre saldo original e saldo conferido -->
    <div class="mb-6 p-4 bg-base-300 border-l-4 border-primary rounded-lg shadow-md">
        <div class="flex justify-between items-center mb-4">
            <div>
                <span class="text-lg font-semibold text-primary">Quantidade na Nota Fiscal</span>
                <p class="text-sm text-warning">Quantidade registrada</p>
            </div>

            <!-- Mostrando o saldo máximo -->
            <span class="text-2xl font-bold text-base-content">
                {#if saldoMaximo !== null}
                    {saldoMaximo}
                {:else}
                    {isLoading ? 'Buscando saldo...' : 'Saldo não encontrado'}
                {/if}
            </span>
        </div>

        <hr class="my-2" />
        <div class="flex justify-between items-center">
            <div>
                <span class="text-lg font-semibold text-primary">Quantidade Apanhada</span>
                <p class="text-sm text-warning">Quantidade real</p>
            </div>
            <span class="text-2xl font-bold text-base-content">{saldoConferencia}</span>
        </div>
    </div>

    <!-- Checkbox de confirmação -->
    <div class="form-control mb-8">
        <label class="cursor-pointer flex items-center space-x-4">
            <input type="checkbox" class="checkbox checkbox-primary" bind:checked={isCheckboxChecked} />
            <span class="label-text text-primary text-lg">Confirmar conferência</span>
        </label>
    </div>

    <!-- Botão de Conferência -->
    <ButtonConferir
        {seqNumero}
        {documento}
        {produto}
        {saldoConferencia}
        {responsavel}
        disabled={!isCheckboxChecked}
    />
</div>
