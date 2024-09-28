package app

import (
	"fmt"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/middleware"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
)

func (app *Application) routes() http.Handler {
	router := httprouter.New()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", fileServer))

	withSessions := alice.New(app.sessionManager.LoadAndSave)

	configureRouter := func(method string, route string, handler http.HandlerFunc) {
		router.Handler(method, route, withSessions.ThenFunc(handler))
	}

	configureRouter(http.MethodGet, "/", app.home)
	configureRouter(http.MethodGet, "/snippet/view/:id", app.snippetView)

	// Auth

	configureRouter(http.MethodGet, "/auth/signup", app.userSignupGet)
	configureRouter(http.MethodPost, "/auth/signup", app.userSignupPost)

	configureRouter(http.MethodGet, "/auth/login", app.userLoginGet)
	configureRouter(http.MethodPost, "/auth/login", app.userLoginPost)

	protectedWithAuth := withSessions.Append(app.requireAuthentication)

	ch := alice.New(app.recoverPanic, app.logHandler, middleware.SecureHeaders)

	configureRouter = func(method string, route string, handler http.HandlerFunc) {
		router.Handler(method, route, protectedWithAuth.ThenFunc(handler))
	}

	configureRouter(http.MethodGet, "/snippet/create", app.snippetCreatePage)

	configureRouter(http.MethodPost, "/snippet/create", app.snippetCreatePageSendForm)
	configureRouter(http.MethodPost, "/snippet/create_api", app.snippetCreateByAPI)

	configureRouter(http.MethodPost, "/auth/logout", app.userLogoutPost)
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
