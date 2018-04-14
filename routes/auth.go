package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
	"pershotravndk.com/models"
	"pershotravndk.com/utils"
)

type authMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RespOptions(rnd render.Render, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
}
func LogIn(rnd render.Render, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var authUser authMessage
	err = json.Unmarshal(b, &authUser)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	user, err := models.GetUserByUsername(authUser.Username, db)
	if err != nil {
		log.Print(err)
		//неверный логин
		rnd.JSON(200, map[string]interface{}{"userID": ""})
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Hashpassword, []byte(authUser.Password))
	if err != nil {
		log.Print(err)
		//неверный пароль
		rnd.JSON(200, map[string]interface{}{"userID": ""})
		return
	}
	fmt.Printf("%s", b)
	rnd.JSON(200, map[string]interface{}{"userID": user.UserID})

}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

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
	rnd.Redirect("/confirmation")
}

func Authorization(rnd render.Render, r *http.Request, db *sql.DB, session sessions.Session) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.GetUserByUsername(username, db)
	if err != nil {
		log.Print(err)
		//неверный логин
		rnd.Redirect("/signIn")
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Hashpassword, []byte(password))
	if err != nil {
		//неверный пароль
		rnd.Redirect("/signIn")
		return
	}
	switch user.Access {
	case 0:
		{
			//не подтвержден аккаунт
			rnd.Redirect("/signIn")
			return
		}
	case 1:
		{
			//гость
			session.Set("userID", user.UserID)
			rnd.Redirect("/guest")
			return
		}
	case 2:
		{
			//админ
			session.Set("userID", user.UserID)
			rnd.Redirect("/admin")
			return
		}
	}
}

func SignOut(rnd render.Render, session sessions.Session) {
	session.Set("userID", "")
	rnd.Redirect("/")
}
