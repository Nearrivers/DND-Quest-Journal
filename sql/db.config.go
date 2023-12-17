package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
)

var db *database.Queries

func ConnectToDb() {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Identifiants de connexion à la base de données manquants")
	}

	conn, err := sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal("Connexion à la base de données impossible")
	}

	db = database.New(conn)
}

func GetDbConnection() *database.Queries {
	return db
}
