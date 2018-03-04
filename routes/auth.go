package routes

import (
	"fmt"
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
	id, err := models.InsertUser(user, db)
	if err != nil {
		log.Print(err)
	}
	fmt.Print(id)

	rnd.Redirect("/")
}
