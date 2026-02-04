package main

import "net/http"

/**
* Example of a simple Go web server that responds with "Hello, World!" to any HTTP request.
* Objective: Kubernetes deployment example.
**/

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

