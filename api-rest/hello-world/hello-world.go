package main

import (
	"fmt"
	"net/http"
)

func saludar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
	fmt.Fprintln(w, "Hola Mundo 2")
}

func main() {
	fmt.Println("hola aplicacion en go")

	http.HandleFunc("POST /character", saludar)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
