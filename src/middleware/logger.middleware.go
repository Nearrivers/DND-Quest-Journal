package middleware

import (
	"log"
	"net/http"
)

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method + ": " + r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
