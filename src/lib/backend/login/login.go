package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url" // Aqui está o pacote necessário para usar url.Values
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

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar o corpo da requisição JSON em creds
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// URL do servidor OAuth
	oauthURL := "http://172.16.99.174:8400/rest/api/oauth2/v1/token"

	// Montar os parâmetros no formato x-www-form-urlencoded
	data := url.Values{} // Utilizando url.Values para montar os parâmetros
	data.Set("grant_type", "password")
	data.Set("username", creds.Username)
	data.Set("password", creds.Password)

	// Preparar a requisição para a API de autenticação
	client := &http.Client{}
	req, err := http.NewRequest("POST", oauthURL, strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Definir os cabeçalhos apropriados
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Executar a requisição
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to authenticate", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Ler o corpo da resposta
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	// Verificar se a resposta é de erro (status code >= 400)
	if res.StatusCode >= 400 {
		http.Error(w, string(body), res.StatusCode)
		return
	}

	// Retornar a resposta da API de autenticação para o frontend
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
