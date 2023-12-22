package quest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
	db "github.com/Nearrivers/DND-quest-tracker/sql"
	questTemplate "github.com/Nearrivers/DND-quest-tracker/src/templates/quest"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

type createQuestDto struct {
	Name                 string
	Description          string
	Npc                  string
	Number               int32
	CompletedDescription string
}

func GetCampaignQuests(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	quests, err := db.GetAllCampaignQuests(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des quêtes : "+err.Error(), http.StatusInternalServerError)
		return
	}

	campaign, err := db.GetOneCampaign(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération d'une campagne :"+err.Error(), http.StatusInternalServerError)
		return
	}

	allQuests := questTemplate.CampaignQuests(campaign, quests)
	allQuests.Render(r.Context(), w)
}

func CreateQuest(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "erreur lors de la lecture du formulaire : "+err.Error(), http.StatusBadRequest)
		return
	}

	newQuest := createQuestDto{}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&newQuest, r.PostForm)
	if err != nil {
		http.Error(w, "erreur lors du décodage : "+err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	result, err := db.CreateQuest(r.Context(), database.CreateQuestParams{
		CreatedAt:            time.Now().UTC(),
		UpdatedAt:            time.Now().UTC(),
		Name:                 newQuest.Name,
		Npc:                  newQuest.Npc,
		Description:          newQuest.Description,
		Number:               newQuest.Number,
		CompletedDescription: newQuest.CompletedDescription,
		CampaignID:           int32(id),
	})

	if err != nil {
		http.Error(w, "Création de la quête impossible :"+err.Error(), http.StatusInternalServerError)
		return
	}

	affectedRows, _ := result.RowsAffected()
	fmt.Println(affectedRows)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "")
}
