package main

import "github.com/tanerijun/snip-snap/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
