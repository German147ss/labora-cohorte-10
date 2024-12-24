package main

import (
	"api-rest-postgresql/internal/config"
	"api-rest-postgresql/internal/domain/user"
	"api-rest-postgresql/internal/infraestructure/http/handler"
	"api-rest-postgresql/internal/infraestructure/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	/*
			// Cargar configuración desde .env
		    cfg, err := config.LoadConfig()
		    if err != nil {
		        log.Fatalf("Error loading config: %v", err)
		    }
	*/

	portAsString := os.Getenv("DB_PORT")
	port, _ := strconv.Atoi(portAsString)

	newDatabase := config.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	// Inicializar conexión a la base de datos
	db, err := config.NewPostgresConnection(newDatabase)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Inicializar dependencias - repository
	//userMemoryRepo := repository.NewUserRepository()
	userPostgresRepo := repository.NewPostgresUserRepository(db)

	fmt.Println(userPostgresRepo)
	userService := user.NewService(userPostgresRepo)
	userHandler := handler.NewUserHandler(userService)

	// Configurar rutas
	http.HandleFunc("GET /users", userHandler.GetUsers)
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("GET /users/{id}", userHandler.GetUserByID)
	http.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	http.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Iniciar servidor
	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
