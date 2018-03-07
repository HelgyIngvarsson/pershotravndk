package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"pershotravndk.com/models"
)

func LeaveFeedback(rnd render.Render, r *http.Request, db *sql.DB, session sessions.Session) {

	feedback := new(models.Feedback)
	feedback.Message = r.FormValue("feedback")
	feedback.UserID = session.Get("userID").(string)

	err := models.InsertFeedback(feedback, db)
	if err != nil {
		log.Print(err)
	}
	rnd.Redirect("/guest")
}

func PostAnonse(rnd render.Render, r *http.Request, db *sql.DB, session sessions.Session) {

	anonse := new(models.Anonse)
	anonse.AnonseDate = r.FormValue("anonse_date")
	anonse.Body = r.FormValue("anonse_body")
	anonse.UserID = session.Get("userID").(string)

	err := models.InsertAnonse(anonse, db)
	if err != nil {
		log.Print(err)
	}
	rnd.Redirect("/admin")
}
