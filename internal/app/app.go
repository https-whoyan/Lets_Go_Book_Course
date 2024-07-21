package app

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/https_whoyan/Lets_Go_Book_Course/cmd/flag"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/endpoints"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
)

type Application struct {
	netPort     string
	infoLogger  *log.Logger
	errorLogger *log.Logger
	mux         *http.ServeMux
	Handler     *endpoints.Handler
	snippets    *models.SnippetModel
}

func NewApplication() (*Application, error) {
	flagCfg, err := flag.NewFlagConfig()
	if err != nil {
		return nil, err
	}
	infoLogger := log.New(os.Stdout, "INGO\t", log.LstdFlags|log.Lshortfile)
	errLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	h := endpoints.NewHandler(infoLogger, errLogger)

	snippetModel, err := models.NewSnippetModel(flagCfg.DbAddr)
	if err != nil {
		return nil, err
	}

	apl := &Application{
		netPort:     flagCfg.NetAddr,
		infoLogger:  infoLogger,
		errorLogger: errLogger,
		mux:         mux,
		Handler:     h,
		snippets:    snippetModel,
	}

	mux.HandleFunc("/", apl.home)
	mux.HandleFunc("/snippet/create", apl.snippetCreate)
	mux.HandleFunc("/snippet/view", apl.snippetView)

	return apl, nil
}

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func (app *Application) Run() {
	srv := http.Server{
		Addr:    app.netPort,
		Handler: app.mux,
	}

	defer func() {
		var err error
		err = srv.Close()
		if err != nil {
			app.errorLogger.Println(err)
		}
		err = nil
	}()

	app.infoLogger.Printf("Running on port %s\n", app.netPort)
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLogger.Fatal(err)
	}
}

func (app *Application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := app.snippets.Insert("example", "example", 1)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecords) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLogger.Print(err.Error())
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", &templateData{Snippet: snippet})
	if err != nil {
		app.errorLogger.Print(err.Error())
		app.serverError(w, err)
		return
	}
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateData{Snippets: snippets}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}
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
