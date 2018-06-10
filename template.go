package main

import (
	"bytes"
	"html/template"

	"github.com/gobuffalo/packr"
	"github.com/rs/zerolog/log"
)

type cardTemplates struct {
	image string
	video string
	music string
}

var templateData cardTemplates

func loadTemplates() {
	log.Debug().Msg("Loading html templates.")
	templates := packr.NewBox("./templates")

	html := templates.String("image.html")
	templateData.image = html
}

type card struct {
	Src    string
	Name   string
	Width  int
	Height int
}

func (c *card) loadHTML() string {
	log.Debug().
		Str("card", c.Name).
		Msg("Parsing html for card.")

	var tpl bytes.Buffer

	tmpl := template.New("card")
	tmpl, _ = tmpl.Parse(templateData.image)
	tmpl.Execute(&tpl, c)

	return tpl.String()
}
