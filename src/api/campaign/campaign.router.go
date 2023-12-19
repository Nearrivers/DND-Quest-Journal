package campaign

import "github.com/go-chi/chi"

func ConfigureCampaignRoutes() *chi.Mux {
	router := chi.NewRouter()

	// CRUD routes
	router.Get("/crud/read", getAllCampaigns)
	router.Get("/crud/read/{id}", getOneCampaign)
	router.Post("/crud/create", createCampaign)
	router.Put("/crud/update/{id}", updateCampaign)
	router.Delete("/crud/delete/{id}", deleteCampaign)

	// Templates only routes (don't use database, just render templates for HTMX)
	router.Get("/template/edit/form/{id}", getEditCampaignTemplate)
	return router
}
