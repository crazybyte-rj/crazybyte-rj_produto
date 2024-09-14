package main

import (
	"fmt"
	"net/http"
)

// Handler para a rota principal "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação Teste - Versão 1.0")
}

// Handler para a rota "/ajuda"
func ajudaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ajuda da Aplicação Teste - Versão 1.0")
}

func main() {
	// Define as rotas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ajuda", ajudaHandler)

	// Inicia o servidor web na porta 8080
	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
