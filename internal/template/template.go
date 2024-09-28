package template

import (
	"net/http"
	"time"

	"github.com/https_whoyan/Lets_Go_Book_Course/internal/models"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/usecases/validator"
)

type TemplateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Flash       string
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
	validator.Validator `form:"-"`
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
}

type UserSignupForm struct {
	validator.Validator `form:"-"`
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
}

type UserLoginForm struct {
	validator.Validator `form:"-"`
	Email               string `form:"email"`
	Password            string `form:"password"`
}
