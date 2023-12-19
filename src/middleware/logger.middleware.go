package middleware

import (
	"log"
	"net/http"
	"strings"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slicedUrl := strings.Split(r.URL.Path, "/")

		if len(slicedUrl[1]) > 0 {
			if slicedUrl[1][len(slicedUrl[1])-1:] != "s" {
				log.Println("Tu as oublié le 's' à la fin de la route")
			}
		}

		log.Println(r.Method + ": " + r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
