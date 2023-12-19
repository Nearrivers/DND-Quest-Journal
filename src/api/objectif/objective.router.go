package objective

import "github.com/go-chi/chi"

func ConfigureObjectiveRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/quest/{id}", GetAllQuestObjectives)
	router.Post("/quest/{id}", CreateObjective)

	return router
}
