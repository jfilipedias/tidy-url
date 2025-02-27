package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jfilipedias/tidy-url/url"
)

func NewRoutes(urlUseCase url.UseCase) http.Handler {
	r := chi.NewRouter()

	r.Get("/{hash}", getURL(urlUseCase))
	r.Post("/url", createURL(urlUseCase))

	return r
}
