package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
* Example of a simple Go web server that responds with "Hello, World!" to any HTTP request.
* Objective: Kubernetes deployment example.
**/

func main() {
	http.HandleFunc("/", Hello)
	fmt.Println("Iniciando servidor na porta 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

