package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"time"
)

/**
* Example of a simple Go web server that responds with "Hello, World!" to any HTTP request.
* Objective: Kubernetes deployment example.
**/

var startedAt = time.Now()



func main() {
	http.HandleFunc("/healthz", Healthz) 
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.HandleFunc("/secret", secret)
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

func secret(w http.ResponseWriter, r *http.Request) {
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	fmt.Fprintf(w, "Secret data - Username: %s, Password: %s", username, password)
}
	
func ConfigMap(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	
	fmt.Fprintf(w, "Family members: %s", string(data))	
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(startedAt)
	
	if uptime.Seconds() < 10 || uptime.Seconds() > 30 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Duration: %s - Unhealthy", uptime.String())
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK - Uptime: %s", uptime.String())
	}

}