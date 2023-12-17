package campaign

import "github.com/go-chi/chi"

func ConfigureCampaignRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/", CreateCampaign)
	return router
}
