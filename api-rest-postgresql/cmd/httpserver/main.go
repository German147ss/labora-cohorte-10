package main

import (
	"api-rest-postgresql/internal/config"
	"api-rest-postgresql/internal/domain/user"
	"api-rest-postgresql/internal/infraestructure/http/handler"
	"api-rest-postgresql/internal/infraestructure/repository"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetRepository() user.Repository {
	if os.Getenv("REPOSITORY_TYPE") == "postgres" {
		portAsString := os.Getenv("DB_PORT")
		port, _ := strconv.Atoi(portAsString)

		newDatabase := config.DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		}
		db, err := config.NewPostgresConnection(newDatabase)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
		return repository.NewPostgresUserRepository(db)
	}

	if os.Getenv("REPOSITORY_TYPE") == "memory" {
		userMemoryRepo := repository.NewUserRepository()
		return userMemoryRepo
	}

	return &repository.FakeUserRepository{}
}

func main() {
	userService := user.NewService(GetRepository())
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
