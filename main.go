package main

import (
	"pershotravndk.com/routes"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
)

func main() {

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
		IndentXML:  true,                       // Output human readable XML
	}))

	staticOptionsAssets := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptionsAssets))

	m.Get("/", routes.IndexHandler)

	m.Run()
}
