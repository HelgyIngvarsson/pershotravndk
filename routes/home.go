package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"pershotravndk.com/models"
)

func GetArticles(rnd render.Render, db *sql.DB) {
	articles, err := models.GetAllArticles(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.JSON(200, map[string]interface{}{"articles": articles})
}

func GetAnonses(rnd render.Render, db *sql.DB, session sessions.Session) {
	anonses, err := models.GetAllAnonses(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.JSON(200, map[string]interface{}{"anonses": anonses})
}

func GetAlbums(rnd render.Render, db *sql.DB) {
	albums, err := models.GetAlbums(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.JSON(200, map[string]interface{}{"albums": albums})
}
func GetAdmins(rnd render.Render, db *sql.DB) {
	admins, err := models.GetAdminsProfile(db)
	if err != nil {
		log.Print(err)
		return
	}
	rnd.JSON(200, map[string]interface{}{"admins": admins})
}

func GetArticle(rnd render.Render, params martini.Params, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	id := params["id"]
	article := new(models.Article)
	article.ArticleID = id
	article, err := models.GetArticleByID(article.ArticleID, db)
	if err != nil {
		log.Print(err)
		rnd.JSON(200, map[string]interface{}{"article": nil})
	}
	comments, err := models.GetCommentsByArticleID(article.ArticleID, db)
	if err != nil {
		log.Print(err)
	}
	rnd.JSON(200, map[string]interface{}{"article": article, "comments": comments})
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

func SendComment(rnd render.Render, r *http.Request, db *sql.DB, session sessions.Session) {

	userID := r.Header.Get("userID")
	comment := new(models.Comment)
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	comment.UserID = userID
	err = models.InsertComment(comment, db)
	if err != nil {
		rnd.JSON(200, map[string]interface{}{"success": false})
		return
	}
	comment.Profile, _ = models.GetProfileByUserID(userID, db)
	rnd.JSON(200, map[string]interface{}{"success": true, "comment": comment})
}
