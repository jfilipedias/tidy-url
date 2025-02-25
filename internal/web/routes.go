package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jfilipedias/tidy-url/ui/pages"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(pages.Home()).ServeHTTP)
	return r
}
