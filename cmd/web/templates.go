package main

import (
	"html/template"
	"path/filepath"

	"github.com/tanerijun/snip-snap/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Extract file name (will act as cache key)
		name := filepath.Base(page)

		// Slice containing base template, partials, and current page
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		// Parse files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
