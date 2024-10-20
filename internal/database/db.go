package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		connectionString = "postgresql://neondb_owner:UJ5swpdQR9Hz@ep-proud-breeze-a50z2015.us-east-2.aws.neon.tech/neondb?sslmode=require"
	}
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
