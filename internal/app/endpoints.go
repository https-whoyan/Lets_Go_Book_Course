package app

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	myErrors "github.com/https_whoyan/Lets_Go_Book_Course/internal/errors"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/template"
	"github.com/https_whoyan/Lets_Go_Book_Course/internal/usecases/validator"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) snippetCreatePage(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = template.SnippetCreateForm{Expires: 365}
	app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *Application) snippetCreatePageSendForm(w http.ResponseWriter, r *http.Request) {
	form := template.SnippetCreateForm{}
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	form.CheckField(validator.NonBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Content, 100), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365),
		"expires", "This field cannot be blank")
	if !form.Valid() {
		data := template.NewTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}
	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.sessionManager.Put(r.Context(), flashKey, "Snippet successfully created")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *Application) snippetCreateByAPI(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	title := params.ByName("title")
	content := params.ByName("content")
	expires, err := strconv.Atoi(params.ByName("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	if title == "" || content == "" || expires == 0 {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, myErrors.ErrNoRecords) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, http.StatusOK, "view.tmpl", data)
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

	app.render(w, http.StatusOK, "home.tmpl", &template.TemplateData{
		Snippets: snippets,
	})
}

// Auth

func (app *Application) userSignupGet(w http.ResponseWriter, r *http.Request) {
	//...
	data := app.newTemplateData(r)
	data.Form = template.UserSignupForm{}
	app.render(w, http.StatusOK, "signup.tmpl", data)
}

func (app *Application) userSignupPost(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) userLoginGet(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) userLoginPost(w http.ResponseWriter, r *http.Request) {

}

func (app *Application) userLogoutPost(w http.ResponseWriter, r *http.Request) {

}
