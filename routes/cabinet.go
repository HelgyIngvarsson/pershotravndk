package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"pershotravndk.com/models"
)

func GetProfile(rnd render.Render, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userID := r.Header.Get("userID")
	profile, err := models.GetProfileByUserID(userID, db)
	if err != nil {
		rnd.Error(404)
		return
	}
	rnd.JSON(200, map[string]interface{}{"profile": profile})
}

func UpdateProfile(rnd render.Render, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userID := r.Header.Get("userID")
	profile := new(models.Profile)
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	profile.UserID = userID
	err = models.UpdateProfie(profile, db)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	rnd.JSON(200, map[string]interface{}{"success": true})
}

func LeaveFeedback(rnd render.Render, r *http.Request, db *sql.DB) {
	userID := r.Header.Get("userID")
	feedback := new(models.Feedback)
	err := json.NewDecoder(r.Body).Decode(&feedback)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	feedback.UserID = userID
	err = models.InsertFeedback(feedback, db)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	rnd.JSON(200, map[string]interface{}{"success": true})
}
func UpdateImage(rnd render.Render, r *http.Request, db *sql.DB) {
	image := new(models.Image)
	err := json.NewDecoder(r.Body).Decode(&image)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	err = models.UpdateImage(image, db)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	rnd.JSON(200, map[string]interface{}{"success": true})
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

func PostArticle(rnd render.Render, r *http.Request, db *sql.DB, session sessions.Session) {

	article := new(models.Article)
	article.Title = r.FormValue("article_title")
	article.Body = r.FormValue("article_body")
	article.UserID = session.Get("userID").(string)
	article.Image = new(models.Image)
	article.Image.ImageID = "1"

	err := models.InsertArticle(article, db)
	if err != nil {
		log.Print(err)
	}
	rnd.Redirect("/admin")
}
