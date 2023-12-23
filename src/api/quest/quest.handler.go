package quest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
	db "github.com/Nearrivers/DND-quest-tracker/sql"
	questTemplate "github.com/Nearrivers/DND-quest-tracker/src/templates/quest"
	"github.com/go-chi/chi"
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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	lastQuestNumber, err := db.GetLastQuest(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	newQuestNumber := 0

	if lastQuestNumber == nil {
		newQuestNumber = 1
	} else {
		n, ok := lastQuestNumber.(int)
		if ok {
			newQuestNumber = n + 1
		} else {
			newQuestNumber = 1
		}
	}

	result, err := db.CreateQuest(r.Context(), database.CreateQuestParams{
		CreatedAt:            time.Now().UTC(),
		UpdatedAt:            time.Now().UTC(),
		Name:                 "Nouvelle quête",
		Npc:                  "Donneur de quête",
		Description:          "Description de la quête en cours",
		CompletedDescription: "Description de la quête lorsqu'elle est terminée",
		Number:               int32(newQuestNumber),
		CampaignID:           int32(id),
	})

	if err != nil {
		http.Error(w, "Création de la quête impossible :"+err.Error(), http.StatusInternalServerError)
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	quest, err := db.GetOneQuest(r.Context(), int32(lastInsertId))
	if err != nil {
		http.Error(w, "Récupération de la quête créée impossible :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	questTemplate.OneQuest(quest).Render(r.Context(), w)
}
