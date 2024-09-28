package app

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/pkg/postgres"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form"
	"github.com/https_whoyan/Lets_Go_Book_Course/cmd/flag"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/repositories/snippets"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/repositories/users"
	myTemplate "github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
)

type Application struct {
	netPort     string
	infoLogger  *log.Logger
	errorLogger *log.Logger
	handler     *http.Handler
	formDecoder *form.Decoder
	templates   *myTemplate.TemplateCache

	snippets *snippets.SnippetModel
	users    *users.UsersModel

	sessionManager *scs.SessionManager
}

func NewApplication() (*Application, error) {
	flagCfg, err := flag.NewFlagConfig()
	if err != nil {
		return nil, err
	}
	infoLogger := log.New(os.Stdout, "INGO\t", log.LstdFlags|log.Lshortfile)
	errLogger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := postgres.Open(flagCfg.DbAddr)
	if err != nil {
		return nil, err
	}

	snippetModel, err := snippets.NewSnippetModel(db)
	if err != nil {
		return nil, err
	}

	usersModels, err := users.NewUsersModel(db)
	if err != nil {
		return nil, err
	}

	templateCache, err := myTemplate.NewTemplateCache()
	if err != nil {
		return nil, err
	}

	apl := &Application{
		netPort:     flagCfg.NetAddr,
		formDecoder: form.NewDecoder(),

		infoLogger:  infoLogger,
		errorLogger: errLogger,
		templates:   templateCache,

		snippets: snippetModel,
		users:    usersModels,
	}
	apl.standSessionManager(flagCfg)
	routes := apl.routes()
	apl.handler = &routes

	return apl, nil
}

func (app *Application) Run() {
	srv := http.Server{
		Addr:    app.netPort,
		Handler: *app.handler,

		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
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
