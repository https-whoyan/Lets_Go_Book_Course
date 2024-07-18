package config

import (
	"log"
	"net/http"
	"os"

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
	mux.HandleFunc("/home", h.Home)
	mux.HandleFunc("/snippet/view", h.SnippetView)
	mux.HandleFunc("/snippet/create", h.SnippetCreate)

	snippetModel, err := models.NewSnippetModel(flagCfg.DbAddr)
	if err != nil {
		return nil, err
	}

	return &Application{
		netPort:     flagCfg.NetAddr,
		infoLogger:  infoLogger,
		errorLogger: errLogger,
		mux:         mux,
		Handler:     h,
		snippets:    snippetModel,
	}, nil
}

func (app *Application) Run() {
	srv := http.Server{
		Addr:    app.netPort,
		Handler: app.mux,
	}

	app.infoLogger.Printf("Running on port %s\n", app.netPort)
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLogger.Fatal(err)
	}
}
