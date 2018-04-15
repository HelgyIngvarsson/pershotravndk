package main

import (
	"fmt"
	"log"
	"os"

	"pershotravndk.com/models"
	"pershotravndk.com/routes"
	"pershotravndk.com/utils"

	"github.com/codegangsta/martini-contrib/cors"
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
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
		IndentXML:  true,    // Output human readable XML
	}))

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:3000", "https://infinite-coast-51488.herokuapp.com"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	m.Get("/api/getArticles", routes.GetArticles)
	m.Get("/api/getAnonses", routes.GetAnonses)
	m.Options("/**", routes.RespOptions)
	m.Post("/api/login", routes.LogIn)
	m.Get("/api/getAlbums", routes.GetAlbums)
	m.Get("/api/getAdmins", routes.GetAdmins)

	// m.Get("/", routes.IndexHandler)
	// m.Get("/signUp", routes.SignUp)
	// m.Get("/signIn", routes.SignIn)
	// m.Get("/signOut", routes.SignOut)
	// m.Get("/guest", routes.GuestCabinet)
	// m.Get("/admin", routes.AdminCabinet)
	// m.Get("/cabinet", routes.Cabinet)
	// m.Get("/confirmation", routes.Confirm)
	// m.Get("/gallery", routes.Gallery)
	// m.Post("/registration", routes.Registration)
	// m.Post("/feedback", routes.LeaveFeedback)
	// m.Post("/post-anonse", routes.PostAnonse)
	// m.Post("/post-article", routes.PostArticle)
	// m.Post("/auth", routes.Authorization)
	// m.Post("/add-comment", routes.AddComment)
	// m.Get("/confirm-email/:token", routes.ConfirmProfile)
	// m.Get("/article/:id", routes.GetArticle)
	port, err := determineListenAddress()
	if err != nil {
		m.Run() //run on default port
	} else {
		m.RunOnAddr(port) //run on determine port
	}
	// m.RunOnAddr(":8000")
}

//function checking environment variable PORT
func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
