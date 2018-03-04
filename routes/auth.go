package routes

import (
	"log"

	"pershotravndk.com/models"

	"database/sql"
	"net/http"

	"github.com/martini-contrib/render"
	"golang.org/x/crypto/bcrypt"
)

func Registration(rnd render.Render, r *http.Request, db *sql.DB) {

	user := new(models.User)

	user.Username = r.FormValue("username")
	user.Hashpassword, _ = bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	// if err != nil {
	// 	log.Print(err)
	// }
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
	if err != nil {
		log.Print(err)
	}
	profile.UserID = userID

	err = models.InsertProfile(profile, db)
	if err != nil {
		log.Print(err)
	}
	rnd.Redirect("/")
}
