package api

import (
	"net/http"
)

func (api *API) newRoutes() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("GET /{hash}", getURL(api.urlUseCase))
	r.HandleFunc("POST /url", createURL(api.urlUseCase))

	return r
}
