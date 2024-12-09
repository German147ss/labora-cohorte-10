package main

import (
	"log"
	"net/http"
)

// Middleware de logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Método: %s, URL: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Llama al siguiente manejador
	})
}

// Manejador simple
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola, mundo!"))
}

func main() {
	mux := http.NewServeMux()

	// Registrar la ruta
	mux.HandleFunc("/", helloHandler)

	// Envolver el mux con el middleware
	loggedMux := loggingMiddleware(mux)

	// Iniciar el servidor
	log.Println("Servidor ejecutándose en http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
