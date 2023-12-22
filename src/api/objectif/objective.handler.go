package objective

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Nearrivers/DND-quest-tracker/internal/database"
	db "github.com/Nearrivers/DND-quest-tracker/sql"
	objectiveTemplate "github.com/Nearrivers/DND-quest-tracker/src/templates/objective"
	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

type createObjectiveDto struct {
	Name        string
	Description string
	Number      int32
}

func GetAllQuestObjectives(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "%v n'est pas reconnu", http.StatusBadRequest)
		return
	}

	db := db.GetDbConnection()

	questName := r.URL.Query().Get("name")

	objectives, err := db.GetAllQuestObjectives(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "lecture des objectifs impossible :"+err.Error(), http.StatusNotFound)
		return
	}

	questObjectives := objectiveTemplate.ObjectiveList(questName, objectives)
	questObjectives.Render(r.Context(), w)
}

func CreateObjective(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "erreur lors de la lecture du formulaire : "+err.Error(), http.StatusBadRequest)
		return
	}

	newObjective := createObjectiveDto{}

	decoder := schema.NewDecoder()
	err = decoder.Decode(&newObjective, r.PostForm)
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

	result, err := db.CreateObjective(r.Context(), database.CreateObjectiveParams{
		UpdatedAt:   time.Now().UTC(),
		CreatedAt:   time.Now().UTC(),
		Name:        newObjective.Name,
		Description: newObjective.Description,
		Number:      newObjective.Number,
		QuestID:     int32(id),
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Création de la campaigne %s impossible : %s", newObjective.Name, err.Error()), http.StatusBadRequest)
		return
	}

	lastInsertId, _ := result.LastInsertId()

	objective, err := db.GetOneObjective(r.Context(), int32(lastInsertId))
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du nouvel objectif :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	newLine := objectiveTemplate.OneObjective(objective)
	newLine.Render(r.Context(), w)
}
