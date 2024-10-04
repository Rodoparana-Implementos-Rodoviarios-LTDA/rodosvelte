// backend/pneus/borracharia/listanfs/listanfs.go
package listanfs

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
)

// Estrutura para os filtros
type Filtros struct {
    Filial    string `json:"filial"`
    Documento string `json:"documento"`
    Serie     string `json:"serie"`
    Cliente   string `json:"cliente"`
    Loja      string `json:"loja"`
}

// Handler para ListaItensNF
func GetListaNFS(w http.ResponseWriter, r *http.Request) {
    // Extrair filtros dos headers
    filtros := Filtros{
        Filial:    r.Header.Get("X-Filial"),
        Documento: r.Header.Get("X-Documento"),
        Serie:     r.Header.Get("X-Serie"),
        Cliente:   r.Header.Get("X-Cliente"),
        Loja:      r.Header.Get("X-Loja"),
    }

    // Validar se todos os filtros necessários foram fornecidos
    if filtros.Filial == "" || filtros.Documento == "" || filtros.Serie == "" || filtros.Cliente == "" || filtros.Loja == "" {
        http.Error(w, "Filtros incompletos nos headers", http.StatusBadRequest)
        return
    }

    // Construir o corpo JSON
    bodyBytes, err := json.Marshal(filtros)
    if err != nil {
        http.Error(w, "Erro ao construir o corpo da requisição", http.StatusInternalServerError)
        return
    }

    // Criar a requisição GET com o corpo
    req, err := http.NewRequest("GET", "http://protheus-vm:9010/rest/MovPortaria/ListaItensNF", bytes.NewBuffer(bodyBytes))
    if err != nil {
        http.Error(w, "Erro ao criar a requisição", http.StatusInternalServerError)
        return
    }

    // Definir o header Content-Type como application/json
    req.Header.Set("Content-Type", "application/json")

    // Realizar a requisição
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, "Erro ao realizar a requisição para a API original", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Copiar os headers da resposta original
    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }

    // Definir o status code da resposta
    w.WriteHeader(resp.StatusCode)

    // Copiar o corpo da resposta original para o cliente
    io.Copy(w, resp.Body)
}
