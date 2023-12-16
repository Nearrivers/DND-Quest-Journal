package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	router := chi.NewRouter()

	godotenv.Load()

	portString := os.Getenv("PORT")

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Serveur démarré sur le port %v", portString)
	log.Fatal(srv.ListenAndServe())
}
