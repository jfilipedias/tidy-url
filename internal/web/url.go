package web

import (
	"fmt"
	"net/http"

	"github.com/jfilipedias/tidy-url/internal/form"
)

type createURLFormData struct {
	OriginalURL string `form:"originalUrl"`
	form.Validator
}

func createURL(w http.ResponseWriter, r *http.Request) {
	var formData createURLFormData
	err := form.Decode(r, &formData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	formData.CheckField(form.NotBlank(formData.OriginalURL), "originalUrl", "The field cannot be blank.")
	formData.CheckField(form.URL(formData.OriginalURL), "originalUrl", "Type a valid URL.")

	if !formData.Valid() {
		fmt.Print("invalid form")
	}

	fmt.Printf("%v", formData)
}
