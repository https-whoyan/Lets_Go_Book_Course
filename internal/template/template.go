package template

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/validator"
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
	validator.Validator `form: "-"`
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	FieldErrors         map[string]string
}
