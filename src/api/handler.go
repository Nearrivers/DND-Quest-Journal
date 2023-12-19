package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/schema"
)

func HandleForm(entity interface{}, r *http.Request) error {
	defer r.Body.Close()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return errors.New("erreur lors de la lecture du formulaire : " + err.Error())
	}

	decoder := schema.NewDecoder()
	err = decoder.Decode(entity, r.PostForm)
	if err != nil {
		return errors.New("erreur lors du décodage : " + err.Error())
	}

	return nil
}
