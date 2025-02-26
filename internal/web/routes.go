package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/url", createURL)
	return r
}
