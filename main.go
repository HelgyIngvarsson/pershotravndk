package main

import (
	"log"

	"pershotravndk.com/models"
	"pershotravndk.com/routes"
	"pershotravndk.com/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func main() {
	// db connection
	db, err := models.NewDB("postgres://jsopcnfzumgznz:20807490dae09e58991e7a56179e659d80a9169bafbba8b01bb996464fed4347@ec2-107-22-175-33.compute-1.amazonaws.com:5432/db56m8m3hnlru2")
	if err != nil {
		log.Panic(err)
	}

	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("secret01121996"))
	m.Use(sessions.Sessions("auth_session", store))

	utils.TokensMonitor()

	m.Map(db)

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
	m.Get("/signUp", routes.SignUp)
	m.Get("/signIn", routes.SignIn)
	m.Get("/signOut", routes.SignOut)
	m.Get("/guest", routes.GuestCabinet)
	m.Get("/admin", routes.AdminCabinet)
	m.Get("/cabinet", routes.Cabinet)
	m.Post("/registration", routes.Registration)
	m.Post("/feedback", routes.LeaveFeedback)
	m.Post("/auth", routes.Authorization)
	m.Get("/confirm-email/:token", routes.ConfirmProfile)
	m.Run()
}
