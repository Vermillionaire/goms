package main

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/gobuffalo/packr"
)

type cardTemplates struct {
	image string
	video string
	music string
}

var templateData cardTemplates

func loadTemplates() {
	templates := packr.NewBox("./templates")

	html := templates.String("image.html")
	templateData.image = html

	fmt.Println("Loaded html: \n" + html)

	var test card
	test.Src = "test1"
	test.Name = "test2"
	fmt.Println(test.loadHTML())
}

type card struct {
	Src  string
	Name string
}

func (c *card) loadHTML() string {
	var tpl bytes.Buffer

	tmpl := template.New("card")
	tmpl, _ = tmpl.Parse(templateData.image)
	tmpl.Execute(&tpl, c)

	return tpl.String()
}
