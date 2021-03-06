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

	var SigningKey = []byte("supersecret")

	m.Map(db)
	m.Map(SigningKey)

	m.Use(render.Renderer(render.Options{
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
		IndentXML:  true,    // Output human readable XML
	}))

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost:3000", "https://pershotravndk.herokuapp.com"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "x-access-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	m.Get("/api/getArticles", routes.GetArticles)
	m.Get("/api/getAnonses", routes.GetAnonses)
	m.Options("/**", routes.RespOptions)
	m.Post("/api/login", routes.LogIn)
	m.Get("/api/getAlbums", routes.GetAlbums)
	m.Get("/api/getAdmins", routes.GetAdmins)
	m.Get("/api/getArticle/:id", routes.GetArticle)
	m.Post("/api/registration", routes.Registration)
	m.Get("/api/current_profile", routes.VarifyToken, routes.GetProfile)
	m.Post("/api/update_profile", routes.VarifyToken, routes.UpdateProfile)
	m.Post("/api/send_feedback", routes.VarifyToken, routes.LeaveFeedback)
	m.Post("/api/update_image", routes.VarifyToken, routes.UpdateImage)
	m.Post("/api/send_comment", routes.VarifyToken, routes.SendComment)

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
