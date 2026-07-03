package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/trokhos/trokhos"
	"github.com/trokhos/trokhos/internal/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := createDB()
	if err != nil {
		log.Fatalf("Error creating a db: %v", err)
	}

	// Create the configuration for Trokhos application
	cfg := new(server.Config)
	cfg.DB = db
	cfg.FS = trokhos.FS

	err = server.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func createDB() (*sql.DB, error) {
	dbString := os.Getenv("TROKHOS_DB_URL")
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
