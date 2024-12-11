package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgresConnection(dbConfig DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConfig.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}
