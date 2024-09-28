package template

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"github.com/https_whoyan/Lets_Go_Book_Course/ui"
)

type TemplateCache map[string]*template.Template

func NewTemplateCache() (*TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		patterns := []string{
			"html/pages/base.tmpl",
			"html/pages/nav.tmpl",
			page,
		}

		ts, internalErr := template.New(name).ParseFS(ui.Files, patterns...)
		if internalErr != nil {
			return nil, err
		}
		cache[name] = ts
	}

	return &cache, nil
}
