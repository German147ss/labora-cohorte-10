package main

import (
	"fmt"
	"net/http"
)

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
	fmt.Fprintln(w, "Hola Mundo 2")
}

func despedirse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "chau")
	fmt.Fprintln(w, "Hola Mundo 2")
}

func main() {
	http.HandleFunc("/", saludar)
	http.HandleFunc("/chau", despedirse)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
