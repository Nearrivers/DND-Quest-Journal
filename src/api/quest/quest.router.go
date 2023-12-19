package quest

import "github.com/go-chi/chi"

func ConfigureQuestRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/campaign/{id}", GetCampaignQuests)
	router.Post("/campaign/{id}", CreateQuest)

	return router
}
