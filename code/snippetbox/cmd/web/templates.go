package main

import "github.com/des-ant/2022-lets-go/code/snippetbox/internal/models"

// Include a Snippets field in the templateData struct.
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
