package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jfilipedias/tidy-url/view"
	"github.com/jfilipedias/tidy-url/view/page"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(page.Home()).ServeHTTP)
	r.Get("/static/*", http.FileServerFS(view.Files).ServeHTTP)

	r.Post("/url", createURL)
	return r
}
