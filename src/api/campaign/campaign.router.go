package campaign

import "github.com/go-chi/chi"

func ConfigureCampaignRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/crud/read", getAllCampaignsMenuCollapsed)
	router.Get("/crud/read/{id}", getOneCampaign)
	router.Post("/crud/create", createCampaign)
	router.Post("/crud/create/first", createFirstCampaign)
	router.Put("/crud/update/{id}", updateCampaign)
	router.Delete("/crud/delete/{id}", deleteCampaign)

	router.Get("/template/menu/expanded", getAllCampaignsMenuExpanded)
	router.Get("/template/create/first", getCreateFirstCampaignTemplate)
	router.Get("/template/edit/form/{id}", getEditCampaignTemplate)
	return router
}
