package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/login"
	"backend/pneus/borracharia"
	"backend/pneus/borracharia/listanfs"
	"backend/pneus/historico"
	"backend/pneus/portaria"
	"backend/prenotas/produtos" // Importa o package de produtos
	"backend/prenotas/tabela"
)

// Middleware para habilitar CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Configura os cabeçalhos CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Permite os cabeçalhos que são solicitados pelo cliente
		if reqHeaders := r.Header.Get("Access-Control-Request-Headers"); reqHeaders != "" {
			w.Header().Set("Access-Control-Allow-Headers", reqHeaders)
		}

		// Lida com requisições OPTIONS (preflight requests)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Passa a requisição para o próximo handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Inicia processos de busca periódica para outros serviços

	// Carrega os dados iniciais na subida do sistema
	borracharia.InitializeCache() // Carrega cache inicial de Borracharia
	portaria.InitializeCache()    // Carrega cache inicial de Portaria
	historico.InitializeCache()   // Carrega cache inicial de Histórico
	tabela.InitializeCache()      // Carrega cache inicial de Pré-Notas
	produtos.InitializeCache()    // Carrega cache inicial de Produtos

	// Multiplexer de rotas
	mux := http.NewServeMux()

	// Rotas de API
	mux.HandleFunc("/api/prenotas", tabela.GetPreNotas)                     // Pré-Notas
	mux.HandleFunc("/api/login", login.AuthHandler)                         // Login
	mux.HandleFunc("/api/borracharia", borracharia.GetBorracharia)          // Borracharia
	mux.HandleFunc("/api/portaria", portaria.GetMovimentosPortaria)         // Portaria
	mux.HandleFunc("/api/pneus/borracharia/listanfs", listanfs.GetListaNFS) // Lista de NFs Borracharia
	mux.HandleFunc("/api/pneus/historico", historico.GetHistorico)          // Histórico de Saídas
	mux.HandleFunc("/api/produtos", produtos.GetProdutos)                   // Produtos

	// Adiciona o middleware de CORS ao multiplexer
	handler := enableCors(mux)

	// Inicia o servidor HTTP
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
