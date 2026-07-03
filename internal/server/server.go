package server

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"
	"log"

	"github.com/pressly/goose/v3"
)

type Config struct {
	FS fs.FS
	DB *sql.DB
}

func Run(cfg *Config) error {
	// Before runing the application call runMigrations
	err := runMigrations(cfg)
	if err != nil {
		return err
	}

	return nil
}

func runMigrations(cfg *Config) error {
	if err := goose.SetDialect("sqlite3"); err != nil {
		err = fmt.Errorf("failed to set goose dialect: %w", err)
		return err
	}
	goose.SetBaseFS(cfg.FS)
	log.Println("Running database migrations...")
	if err := goose.UpContext(context.Background(), cfg.DB, "migrations"); err != nil {
		err = fmt.Errorf("migration failed: %w", err)
		return err
	}
	log.Println("Migrations completed successfully!")
	return nil
}
