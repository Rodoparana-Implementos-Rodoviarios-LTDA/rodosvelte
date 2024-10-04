<script lang="ts">
    export let seqNumero: string | null = null;
    export let saldoConferencia: number;
    export let documento: string;
    export let produto: string;
    export let responsavel: string;
    export let disabled = false; // Propriedade para desabilitar o botão

    let modalVisible = false;

    function openModal() {
        if (!disabled) modalVisible = true;
    }

    function closeModal() {
        modalVisible = false;
    }

    async function confirmConferencia() {
        if (!seqNumero) {
            alert('Número sequencial não encontrado. Conferência não pode ser realizada.');
            return;
        }

        const body = {
            numseq: seqNumero,
            documento,
            produto,
            quantidade: saldoConferencia,
            responsavel
        };

        console.log('Enviando POST para ConferênciaSaida:', body);

        try {
            const response = await fetch('http://protheus-vm:9010/rest/MovPortaria/ConferenciaSaida', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(body)
            });

            if (!response.ok) {
                const errorText = await response.text();
                console.error('Erro ao confirmar conferência:', errorText);
                alert('Erro ao confirmar conferência: ' + errorText);
                return;
            }

            alert('Conferência realizada com sucesso!');
            closeModal();
        } catch (error) {
            console.error('Erro ao enviar conferência:', error);
            alert('Erro ao confirmar conferência.');
        }
    }
</script>

<!-- Botão que abre o modal, com suporte a "disabled" -->
<button class="btn btn-primary w-full" on:click={openModal} disabled={disabled}>
    Confirmar Conferência
</button>

<!-- Modal de confirmação -->
{#if modalVisible}
    <dialog class="modal modal-open">
        <div class="modal-box">
            <h3 class="font-bold text-primary text-lg">Tem certeza?</h3>
            <p class="py-4 text-neutral-content">Você está prestes a confirmar a conferência. Deseja continuar?</p>
            <div class="modal-action">
                <button class="btn btn-primary-content" on:click={closeModal}>
                    Cancelar
                </button>
                <button class="btn btn-primary" on:click={confirmConferencia}>
                    Sim, Confirmar
                </button>
            </div>
        </div>
    </dialog>
{/if}
