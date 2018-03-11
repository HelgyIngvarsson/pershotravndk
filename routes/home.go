package routes

import (
	"database/sql"
	"log"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"pershotravndk.com/models"
)

func IndexHandler(rnd render.Render, db *sql.DB) {
	anonses, err := models.GetAllAnonses(db)
	if err != nil {
		log.Print(err)
		return
	}
	articles, err := models.GetAllArticles(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.HTML(200, "index", map[string]interface{}{
		"Anonses":  anonses,
		"Articles": articles,
	})
}

func Gallery(rnd render.Render, db *sql.DB) {
	albums, err := models.GetAlbums(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.HTML(200, "gallery", albums)
}

func Confirm(rnd render.Render) {
	rnd.HTML(200, "confirmation", nil)
}

func SignUp(rnd render.Render) {
	rnd.HTML(200, "signUp", nil)
}
func SignIn(rnd render.Render) {
	rnd.HTML(200, "signIn", nil)
}
func GuestCabinet(rnd render.Render, db *sql.DB, session sessions.Session) {
	userID := session.Get("userID").(string)
	if userID == "" {
		rnd.Redirect("/")
		return
	}
	user, err := models.GetUserByID(userID, db)
	if err != nil {
		log.Print(err)
		return
	}
	if user.Access == 1 {
		profile, err := models.GetProfileByUserID(userID, db)
		if err != nil {
			log.Print(err)
			return
		}
		rnd.HTML(200, "guest", map[string]interface{}{
			"Profile": profile,
			"User":    user})
	} else {
		rnd.Redirect("/")
		return
	}

}

func AdminCabinet(rnd render.Render, db *sql.DB, session sessions.Session) {
	userID := session.Get("userID").(string)
	if userID == "" {
		rnd.Redirect("/")
		return
	}
	user, err := models.GetUserByID(userID, db)
	if err != nil {
		log.Print(err)
		return
	}
	if user.Access == 2 {
		profile, err := models.GetProfileByUserID(userID, db)
		if err != nil {
			log.Print(err)
			return
		}
		feedbacks, err := models.GetAllFeedbacks(db)
		if err != nil {
			log.Print(err)
			return
		}
		rnd.HTML(200, "admin", map[string]interface{}{
			"Profile":   profile,
			"Feedbacks": feedbacks,
			"User":      user})
	} else {
		rnd.Redirect("/")
		return
	}
}

func Cabinet(rnd render.Render, db *sql.DB, session sessions.Session) {
	userID := session.Get("userID").(string)
	if userID == "" {
		rnd.Redirect("/")
		return
	}
	user, err := models.GetUserByID(userID, db)
	if err != nil {
		log.Print(err)
		return
	}
	switch user.Access {
	case 0:
		{
			rnd.Redirect("/")
			return
		}
	case 1:
		{
			rnd.Redirect("/guest")
			return
		}
	case 2:
		{
			rnd.Redirect("/admin")
			return
		}
	}
}
