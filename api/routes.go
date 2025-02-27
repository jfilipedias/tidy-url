package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *API) newRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/{hash}", getURL(api.urlUseCase))
	r.Post("/url", createURL(api.urlUseCase))

	return r
}
