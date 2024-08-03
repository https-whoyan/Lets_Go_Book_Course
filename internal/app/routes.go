package app

import (
	"fmt"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/endpoints"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", fileServer))

	router.HandlerFunc(http.MethodGet, "/", app.home)

	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreatePage)

	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePageSendForm)
	router.HandlerFunc(http.MethodPost, "/snippet/create_api", app.snippetCreateByAPI)

	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)

	ch := alice.New(app.recoverPanic, app.logHandler, endpoints.SecureHeaders)

	return ch.Then(router)
}

// Errs
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *Application) serverError(w http.ResponseWriter, err error) {
	app.errorLogger.Println(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Template
func (app *Application) render(w http.ResponseWriter, status int, page string, data *template.TemplateData) {
	ts, ok := (*app.templates)[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}
	w.WriteHeader(status)
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
