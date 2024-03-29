package main

import (
	"AITUNews/pkg/forms"
	"AITUNews/pkg/models"
	"html/template"
	"path/filepath"
)

type templateData struct {
	CSRFToken       string
	Flash           string
	Form            *forms.Form
	News            *models.News
	IsAuthenticated bool
	IsTeacher       bool
	IsStudent       bool
	IsAdmin         bool
	NewsData        []*models.News
	FoodsData       []*models.Foods
	CurrentYear     int
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, err
}
