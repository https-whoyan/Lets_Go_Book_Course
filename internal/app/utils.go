package app

import (
	"errors"
	"github.com/go-playground/form"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
	"net/http"
	"time"
)

func (app *Application) decodePostForm(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	err := app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}

func (app *Application) newTemplateData(r *http.Request) *template.TemplateData {
	return &template.TemplateData{
		CurrentYear: time.Now().Year(),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
	}
}
