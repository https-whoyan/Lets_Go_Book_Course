package template

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
	"net/http"
	"time"
)

type TemplateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
}

func NewTemplateData(_ *http.Request) *TemplateData {
	return &TemplateData{
		CurrentYear: time.Now().Year(),
		Snippet:     nil,
		Snippets:    nil,
		Form: SnippetCreateForm{
			Expires: 365,
		},
	}
}

type SnippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}
