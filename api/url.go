package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createURL(w http.ResponseWriter, r *http.Request) {
	var input struct {
		OriginalURL string `json:"originalUrl"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	fmt.Printf("%v", input)
}
