package main

import (
	"fmt"
	"net/http"

	"backend/login"    // Importa o package de login
	"backend/prenotas" // Importa o package prenotas
)

// Middleware para habilitar CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permitir acesso de qualquer origem
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Permitir métodos HTTP específicos
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Permitir cabeçalhos específicos
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Responder imediatamente para requisições OPTIONS (preflight requests)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passa a requisição para o próximo handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Inicia o processo de busca periódica das pre_notas
	prenotas.StartFetchingPreNotas()

	// Multiplexer de rotas
	mux := http.NewServeMux()

	// Registrar a rota de pre_notas
	mux.HandleFunc("/api/prenotas", prenotas.GetPreNotas)

	// Registrar a rota de login
	mux.HandleFunc("/api/login", login.AuthHandler)

	// Adiciona o middleware de CORS ao multiplexer de rotas
	handler := enableCors(mux)

	// Iniciar o servidor HTTP na porta 8080 com o middleware de CORS
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", handler)
}
