package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "germanmendieta"
	password = "admin"
	dbname   = "myapp"
)

func main() {
	// Construir la cadena de conexión
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conectar a la base de datos
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Crear tabla si no existe
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE
	)`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Insertar un nuevo usuario
	_, err = db.Exec(`INSERT INTO users (name, email) VALUES ($1, $2)`, "John Doe", "john.doe@example.com")
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
	}

	// Consultar usuarios
	rows, err := db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		log.Fatalf("Error querying users: %v", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error with rows: %v", err)
	}
}
