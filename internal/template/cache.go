package template

import (
	"html/template"
	"path/filepath"
)

type TemplateCache map[string]*template.Template

func NewTemplateCache() (*TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/pages/base.tmpl",
			"./ui/html/pages/nav.tmpl",
			page,
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &cache, nil
}
