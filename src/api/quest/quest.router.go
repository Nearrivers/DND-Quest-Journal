package quest

import "github.com/go-chi/chi"

func ConfigureQuestRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/campaign/{id}", GetCampaignQuests)
	router.Post("/crud/create/campaign/{id}", CreateQuest)

	return router
}
