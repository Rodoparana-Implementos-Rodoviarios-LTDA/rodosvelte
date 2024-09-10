package login

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

// Credentials contém as credenciais do usuário
type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// AuthHandler é o handler que processa o login e faz a requisição à API de autenticação
func AuthHandler(w http.ResponseWriter, r *http.Request) {
    // Habilitar CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Se for uma requisição OPTIONS, responde e sai
    if r.Method == "OPTIONS" {
        w.WriteHeader(http.StatusOK)
        return
    }

    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // URL do seu servidor OAuth
    url := "http://172.16.99.174:8400/rest/api/oauth2/v1/token"
    method := "POST"

    client := &http.Client{}
    payload := strings.NewReader(fmt.Sprintf(`{
        "grant_type": "password",
        "username": "%s",
        "password": "%s"
    }`, creds.Username, creds.Password))

    req, err := http.NewRequest(method, url, payload)
    if err != nil {
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }

    req.Header.Add("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        http.Error(w, "Failed to read response", http.StatusInternalServerError)
        return
    }

    w.Write(body) // Retorna a resposta da API de autenticação para o frontend
}
