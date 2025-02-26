package form

import (
	"errors"
	"net/http"

	"github.com/go-playground/form/v4"
)

func Decode(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	formDecoder := form.NewDecoder()
	err = formDecoder.Decode(&dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}
