package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"golang.org/x/crypto/bcrypt"
	"pershotravndk.com/models"
	"pershotravndk.com/utils"
)

func Registration(rnd render.Render, r *http.Request, db *sql.DB) {

	user := new(models.User)

	user.Username = r.FormValue("username")
	user.Hashpassword, _ = bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

	userID, err := models.InsertUser(user, db)
	if err != nil {
		log.Print(err)
	}

	profile := new(models.Profile)
	profile.Name = r.FormValue("name")
	profile.Sername = r.FormValue("sername")
	profile.Email = r.FormValue("email")
	profile.Description = r.FormValue("description")
	if r.FormValue("mailing") == "on" {
		profile.Mailing = true
	}
	profile.UserID = userID

	err = models.InsertProfile(profile, db)
	if err != nil {
		log.Print(err)
	}

	token := utils.GenerateToken()

	err = models.InsertToken(userID, token, db)
	if err != nil {
		log.Print(err)
	}

	err = utils.SendMessage(
		[]string{profile.Email},
		"To confirm your registration on pershotravndk.com follow this link https://pershotravndk.herokuapp.com/confirm-email/"+token)
	if err != nil {
		log.Print(err)
	}

	rnd.HTML(200, "regSuccess", nil)
}

func ConfirmProfile(rnd render.Render, params martini.Params, db *sql.DB) {

	token := params["token"]

	userID, err := models.GetUserIDFromToken(token, db)

	if err != nil {
		log.Print(err)
		rnd.Redirect("/", 404)
		return
	}

	access, err := models.GetAccess(userID, db)
	if err != nil {
		log.Print(err)
	}
	if access != "0" {
		rnd.Redirect("/")
		return
	}

	err = models.SetAccess(userID, db)
	if err != nil {
		log.Print(err)
	}
	rnd.Redirect("/")
}

func Authorization(rnd render.Render, r *http.Request, db *sql.DB) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.GetUserByUsername(username, db)
	if err != nil {
		log.Print(err)
		rnd.Redirect("/signIn")
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Hashpassword, []byte(password))
	if err != nil {
		rnd.Redirect("/signIn")
		return
	}
	switch user.Access {
	case 0:
		{
			rnd.Redirect("/signIn")
			return
		}
	case 1:
		{
			rnd.Redirect("/")
			return
		}
	}
}
