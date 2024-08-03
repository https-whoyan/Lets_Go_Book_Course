package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/https_whoyan/Lets_Go_Book_Course/cmd/flag"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
	myTemplate "github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
)

type Application struct {
	netPort     string
	infoLogger  *log.Logger
	errorLogger *log.Logger
	handler     *http.Handler
	templates   *myTemplate.TemplateCache
	snippets    *models.SnippetModel
}

func NewApplication() (*Application, error) {
	flagCfg, err := flag.NewFlagConfig()
	if err != nil {
		return nil, err
	}
	infoLogger := log.New(os.Stdout, "INGO\t", log.LstdFlags|log.Lshortfile)
	errLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	snippetModel, err := models.NewSnippetModel(flagCfg.DbAddr)
	if err != nil {
		return nil, err
	}

	templateCache, err := myTemplate.NewTemplateCache()
	if err != nil {
		return nil, err
	}

	apl := &Application{
		netPort:     flagCfg.NetAddr,
		infoLogger:  infoLogger,
		errorLogger: errLogger,
		snippets:    snippetModel,
		templates:   templateCache,
	}
	routes := apl.routes()
	apl.handler = &routes

	return apl, nil
}

func (app *Application) Run() {
	srv := http.Server{
		Addr:    app.netPort,
		Handler: *app.handler,
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

func (app *Application) logHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLogger.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *Application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				w.WriteHeader(http.StatusInternalServerError)
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
