package storage

import (
	"database/sql"
	"fmt"
	"github.com/TeamTells/portal_documents_service/config"
	"log"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(config config.Config) *Storage {

	connectString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)

	db, err := sql.Open(config.DBDriver, connectString)
	if err != nil {
		log.Fatalf("open database file error: %w", err)
		return nil
	}

	if errPing := db.Ping(); errPing != nil {
		log.Fatalf("ping error: %w", errPing)
		return nil
	}

	return &Storage{
		db: db,
	}
}
