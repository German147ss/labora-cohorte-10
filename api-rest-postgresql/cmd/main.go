package main

import (
	"api-rest-postgresql/internal/config"
	"api-rest-postgresql/internal/domain/user"
	"api-rest-postgresql/internal/infraestructure/http/handler"
	"api-rest-postgresql/internal/infraestructure/repository"
	"log"
	"net/http"
)

func main() {
	/*
			// Cargar configuración desde .env
		    cfg, err := config.LoadConfig()
		    if err != nil {
		        log.Fatalf("Error loading config: %v", err)
		    }
	*/

	newDatabase := config.DatabaseConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "germanmendieta",
		Password: "admin",
		DBName:   "myapp",
	}
	// Inicializar conexión a la base de datos
	db, err := config.NewPostgresConnection(newDatabase)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Inicializar dependencias
	userRepo := repository.NewUserRepository()
	userService := user.NewService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Configurar rutas
	http.HandleFunc("/users", userHandler.HandleUsers)
	http.HandleFunc("/users/", userHandler.HandleUserByID)

	// Iniciar servidor
	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
