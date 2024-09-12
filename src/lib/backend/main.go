package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/login"
	"backend/pneus/borracharia"
	"backend/pneus/portaria"
	"backend/prenotas/tabela"
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
	// Inicia o processo de busca periódica das pre_notas, borracharia, portaria e revisao
	tabela.StartFetchingPreNotas()          // Iniciar busca periódica para pre_notas
	portaria.StartFetchingMovimentosPortaria() // Iniciar busca periódica para portaria
	borracharia.StartFetchingBorracharia()     // Iniciar busca periódica para borracharia

	// Multiplexer de rotas
	mux := http.NewServeMux()

	// Registrar as rotas de pre_notas
	mux.HandleFunc("/api/prenotas", tabela.GetPreNotas)

	// Registrar a rota de login
	mux.HandleFunc("/api/login", login.AuthHandler)

	// Registrar a rota de borracharia
	mux.HandleFunc("/api/borracharia", borracharia.GetBorracharia)

	// Registrar a rota de portaria
	mux.HandleFunc("/api/portaria", portaria.GetMovimentosPortaria)



	// Adiciona o middleware de CORS ao multiplexer de rotas
	handler := enableCors(mux)

	// Iniciar o servidor HTTP na porta 8080 com o middleware de CORS
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
	