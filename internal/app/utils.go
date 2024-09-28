package app

import (
	"errors"
	"github.com/go-playground/form"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
	"github.com/justinas/nosurf"
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

func (app *Application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), authKey)
}

func (app *Application) newTemplateData(r *http.Request) *template.TemplateData {
	return &template.TemplateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), flashKey),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}
