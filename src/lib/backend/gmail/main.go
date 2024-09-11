package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	gmail "google.golang.org/api/gmail/v1"
	"github.com/joho/godotenv"
)

func init() {
	// Carrega o arquivo .env.local com as variáveis de ambiente
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env.local: %v", err)
	}
}

// Substitua com suas credenciais carregadas do .env.local
var (
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURL  = os.Getenv("REDIRECT_URL")
)

// Configuração OAuth2 para Google
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  redirectURL,
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
	Endpoint:     google.Endpoint,
}

var oauthStateString = "pseudo-random-string"

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)
	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Página principal
func handleHome(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html><body><a href="/login">Google Login</a></body></html>`
	fmt.Fprint(w, htmlIndex)
}

// Handle do login do Google
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Gera a URL de autorização com access_type=offline para garantir que o refresh token seja gerado
	url := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle do callback OAuth2
func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Verifica o estado para garantir que a requisição é legítima
	if r.FormValue("state") != oauthStateString {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Obtém o código de autorização da URL
	code := r.FormValue("code")

	// Troca o código pelo token de acesso
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Falha ao trocar o código por token: %v", err)
	}

	// Exibe o token de acesso e o refresh token, se presente
	fmt.Fprintf(w, "Token de Acesso: %s\n", token.AccessToken)
	if token.RefreshToken != "" {
		fmt.Fprintf(w, "Refresh Token: %s\n", token.RefreshToken)
	} else {
		fmt.Fprintf(w, "Refresh Token não disponível, pode já ter sido emitido antes.")
	}

	// Aqui chamamos a função para enviar um e-mail usando o token de acesso
	sendEmail(token.AccessToken)
}

// Função para enviar e-mail usando a Gmail API
func sendEmail(accessToken string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	client := oauth2.NewClient(ctx, ts)

	// Criação do serviço Gmail
	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Erro ao criar serviço Gmail: %v", err)
	}

	// Criar a mensagem do e-mail
	var message gmail.Message
	emailTo := "suporte@rodoparana.com.br"             // Substitua pelo e-mail do destinatário
	emailFrom := "guilherme.correia@rodoparana.com.br" // O e-mail do remetente (autenticado)
	subject := "Teste de Envio de E-mail via API"      // Assunto do e-mail
	body := "Este é um teste de envio de e-mail usando a Gmail API com Go. Dale seus lindos"

	// Formatar a mensagem no formato correto (RFC 2822)
	rawMessage := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", emailFrom, emailTo, subject, body)
	encodedMessage := base64.URLEncoding.EncodeToString([]byte(rawMessage))

	// Definir a mensagem no campo 'Raw'
	message.Raw = encodedMessage

	// Enviar a mensagem
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Fatalf("Erro ao enviar e-mail: %v", err)
	} else {
		fmt.Println("E-mail enviado com sucesso!")
	}
}
