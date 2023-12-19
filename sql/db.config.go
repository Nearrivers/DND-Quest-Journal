package db

import (
	"database/sql"
	"errors"
	"os"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
)

var db *database.Queries

func ConnectToDb() error {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return errors.New("identifiants de connexion à la base de données manquants")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		return errors.New("connexion à la base de données impossible")
	}

	db = database.New(conn)
	return nil
}

func GetDbConnection() *database.Queries {
	return db
}
