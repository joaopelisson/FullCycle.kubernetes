package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

/**
* Example of a simple Go web server that responds with "Hello, World!" to any HTTP request.
* Objective: Kubernetes deployment example.
**/

func main() {
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	fmt.Println("Iniciando servidor na porta 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, World! My name is %s and I am %s years old.", name, age)	

}


func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	
	fmt.Fprintf(w, "Family members: %s", string(data))	
}
