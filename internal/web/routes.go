package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/jfilipedias/tidy-url/ui"
	"github.com/jfilipedias/tidy-url/ui/page"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/", templ.Handler(page.Home()).ServeHTTP)
	r.Get("/static/*", http.FileServerFS(ui.Files).ServeHTTP)
	return r
}
