package campaign

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
	db "github.com/Nearrivers/DND-quest-tracker/sql"
	campaignDtos "github.com/Nearrivers/DND-quest-tracker/src/api/campaign/dtos"
	campaignTemplate "github.com/Nearrivers/DND-quest-tracker/src/templates/campaign"
	questTemplate "github.com/Nearrivers/DND-quest-tracker/src/templates/quest"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

// CRUD Routes

func getAllCampaigns(w http.ResponseWriter, r *http.Request) {
	db := db.GetDbConnection()

	campaings, err := db.GetAllCampaigns(r.Context())
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des campagnes : "+err.Error(), http.StatusInternalServerError)
		return
	}

	if len(campaings) == 0 {
		http.Error(w, "Aucune campagne trouvée", http.StatusNotFound)
		return
	}

	allCampaings := campaignTemplate.AllCampaigns(campaings)
	allCampaings.Render(r.Context(), w)
}

func getOneCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	campaign, err := db.GetOneCampaign(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération d'une campagne :"+err.Error(), http.StatusInternalServerError)
		return
	}

	campaignQuests := campaignTemplate.CampaignQuests(campaign)
	campaignQuests.Render(r.Context(), w)
}

func createCampaign(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du formulaire : "+err.Error(), http.StatusBadRequest)
		return
	}

	newCampaign := campaignDtos.CreateCampaignDto{}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&newCampaign, r.PostForm)
	if err != nil {
		http.Error(w, "Erreur lors du décodage : "+err.Error(), http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	result, err := db.CreateCampaign(r.Context(), database.CreateCampaignParams{
		UpdatedAt: time.Now().UTC(),
		CreatedAt: time.Now().UTC(),
		Name:      newCampaign.Name,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Création de la campaigne %s impossible : %s", newCampaign.Name, err.Error()), http.StatusBadRequest)
		return
	}

	lastInsertId, _ := result.LastInsertId()

	campaign, err := db.GetOneCampaign(r.Context(), int32(lastInsertId))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la campagne créée :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	newLine := campaignTemplate.CreatedCampaign(campaign)
	newLine.Render(r.Context(), w)
}

func updateCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	db := db.GetDbConnection()

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du formulaire : "+err.Error(), http.StatusBadRequest)
		return
	}

	editedCampaign := campaignDtos.CreateCampaignDto{}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&editedCampaign, r.PostForm)
	if err != nil {
		http.Error(w, "Erreur lors du décodage : "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.UpdateCampaign(r.Context(), database.UpdateCampaignParams{
		Name:      editedCampaign.Name,
		UpdatedAt: time.Now().Local(),
		ID:        int32(id),
	})

	if err != nil {
		http.Error(w, "Modification de la campagne impossible : "+err.Error(), http.StatusInternalServerError)
		return
	}

	campaign, err := db.GetOneCampaign(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de la campagne modifiée :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	editedLine := questTemplate.CampaignTitle(campaign)
	editedLine.Render(r.Context(), w)
}

func deleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, fmt.Sprintf("%v n'est pas reconnu", id), http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	err = db.DeleteCampaign(r.Context(), int32(id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Suppression de la campagne impossible :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

// Templates only routes

func getEditCampaignTemplate(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	campaign, err := db.GetOneCampaign(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération d'une campagne :"+err.Error(), http.StatusInternalServerError)
		return
	}

	editCampaign := campaignTemplate.EditCampaign(campaign.ID, campaign.Name)
	editCampaign.Render(r.Context(), w)
}
