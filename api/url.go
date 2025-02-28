package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jfilipedias/tidy-url/constant"
	"github.com/jfilipedias/tidy-url/url"
)

func createURL(uc url.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			OriginalURL string `json:"originalUrl"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		w.WriteHeader(201)
	}
}

func getURL(uc url.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		u, err := uc.Get(r.Context(), hash)
		if err != nil {
			if errors.Is(err, constant.ErrEntityNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}
