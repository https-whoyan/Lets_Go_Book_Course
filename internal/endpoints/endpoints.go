package endpoints

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewHandler(infoLogger *log.Logger, errorLogger *log.Logger) *Handler {
	return &Handler{
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl",
		"./ui/html/pages/home.tmpl",
		"./ui/html/pages/nav.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		h.errorLogger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		h.errorLogger.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, err = fmt.Fprintf(w, "Your id: %d", id)
	if err != nil {
		h.errorLogger.Print(err.Error())
	}
}

func (h *Handler) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := w.Write([]byte("Creating new snippet..."))
	if err != nil {
		h.errorLogger.Print(err.Error())
	}
}
