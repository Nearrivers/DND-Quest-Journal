package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/Nearrivers/DND-quest-tracker/sql"
	"github.com/Nearrivers/DND-quest-tracker/src/api/campaign"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func setupFilesRoutes() *chi.Mux {
	filesRouter := chi.NewRouter()

	// Routes pour servir les fichiers (script HTMX et styles Tailwind)
	styleFilesServer := http.FileServer(http.Dir("./styles"))
	scriptFilesServer := http.FileServer(http.Dir("./scripts"))

	filesRouter.Handle("/styles/*", http.StripPrefix("/styles", styleFilesServer))
	filesRouter.Handle("/scripts/*", http.StripPrefix("/scripts", scriptFilesServer))

	return filesRouter
}

func main() {
	godotenv.Load()

	db.ConnectToDb()

	router := chi.NewRouter()
	filesRouter := setupFilesRoutes()
	campaignRouter := campaign.ConfigureCampaignRoutes()

	router.Use(cors.Handler(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Mount("/", filesRouter)
	router.Mount("/campaigns", campaignRouter)

	// Affichage de la page index
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		homePage := index()
		homePage.Render(r.Context(), w)
	})

	portString := os.Getenv("PORT")

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Serveur démarré sur le port %v", portString)
	log.Fatal(srv.ListenAndServe())
}
