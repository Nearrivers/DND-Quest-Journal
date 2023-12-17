package campaign

import "github.com/go-chi/chi"

func ConfigureCampaignRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", GetAllCampaigns)
	router.Get("/{id}", GetOneCampaign)
	router.Post("/", CreateCampaign)
	router.Put("/{id}", UpdateCampaign)
	router.Delete("/{id}", DeleteCampaign)

	return router
}
