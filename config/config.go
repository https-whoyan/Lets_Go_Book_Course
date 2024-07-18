package config

import (
	"log"
	"net/http"
	"os"

	"github.com/https_whoyan/Lets_Go_Book_Course/internal/endpoints"
)

type Application struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	mux         *http.ServeMux
	Handler     *endpoints.Handler
}

func NewApplication() *Application {
	infoLogger := log.New(os.Stdout, "INGO\t", log.LstdFlags|log.Lshortfile)
	errLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	h := endpoints.NewHandler(infoLogger, errLogger)
	mux.HandleFunc("/home", h.Home)
	mux.HandleFunc("/snippet/view", h.SnippetView)
	mux.HandleFunc("/snippet/create", h.SnippetCreate)

	return &Application{
		infoLogger:  infoLogger,
		errorLogger: errLogger,
		mux:         mux,
		Handler:     h,
	}
}

func (app *Application) Run(port string) {
	srv := http.Server{
		Addr:    port,
		Handler: app.mux,
	}

	app.infoLogger.Printf("Running on port %s\n", port)
	err := srv.ListenAndServe()
	if err != nil {
		app.errorLogger.Fatal(err)
	}
}
