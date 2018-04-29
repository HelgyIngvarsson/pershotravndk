package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/martini-contrib/render"
	"golang.org/x/crypto/bcrypt"
	"pershotravndk.com/models"
)

func RespOptions(rnd render.Render, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
}
func LogIn(rnd render.Render, SigningKey []byte, w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type authMessage struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var authUser authMessage
	err := json.NewDecoder(r.Body).Decode(&authUser)
	if err != nil {
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	user, err := models.GetUserByUsername(authUser.Username, db)
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	err = bcrypt.CompareHashAndPassword(user.Hashpassword, []byte(authUser.Password))
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.UserID,
		"exp":    time.Now().Add(time.Hour * 48).Unix()})
	tokenString, _ := token.SignedString(SigningKey)
	rnd.JSON(200, map[string]interface{}{"success": true, "token": tokenString})
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-access-token")
}

func Registration(rnd render.Render, SigningKey []byte, w http.ResponseWriter, r *http.Request, db *sql.DB) {

	type reqData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	var req reqData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	user := new(models.User)
	user.Username = req.Username
	user.Hashpassword, _ = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	userID, err := models.InsertUser(user, db)
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	profile := new(models.Profile)
	profile.Avatar = new(models.Image)
	profile.Avatar.ImageID, err = models.InsertEmptyImage(db)
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	profile.Email = req.Email
	profile.UserID = userID
	err = models.InsertProfile(profile, db)
	if err != nil {
		log.Print(err)
		rnd.JSON(500, map[string]interface{}{"success": false, "token": nil})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 48).Unix()})
	tokenString, _ := token.SignedString(SigningKey)
	rnd.JSON(200, map[string]interface{}{"success": true, "token": tokenString})
}

func VarifyToken(rnd render.Render, SigningKey []byte, w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get("x-access-token")
	if header == "" {
		rnd.Error(401)
		return
	}
	token, _ := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return SigningKey, nil
	})
	var userID string
	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		userID = claims["userID"].(string)
	}
	r.Header.Set("UserID", userID)
}

// func Registration(rnd render.Render, r *http.Request, db *sql.DB) {

// 	user := new(models.User)

// 	user.Username = r.FormValue("username")
// 	user.Hashpassword, _ = bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

// 	userID, err := models.InsertUser(user, db)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	profile := new(models.Profile)
// 	profile.Name = r.FormValue("name")
// 	profile.Sername = r.FormValue("sername")
// 	profile.Email = r.FormValue("email")
// 	profile.Description = r.FormValue("description")
// 	if r.FormValue("mailing") == "on" {
// 		profile.Mailing = true
// 	}
// 	profile.UserID = userID

// 	err = models.InsertProfile(profile, db)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	token := utils.GenerateToken()

// 	err = models.InsertToken(userID, token, db)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	err = utils.SendMessage(
// 		[]string{profile.Email},
// 		"To confirm your registration on pershotravndk.com follow this link https://pershotravndk.herokuapp.com/confirm-email/"+token)
// 	if err != nil {
// 		log.Print(err)
// 	}

// 	rnd.HTML(200, "regSuccess", nil)
// }

// func ConfirmProfile(rnd render.Render, params martini.Params, db *sql.DB) {

// 	token := params["token"]

// 	userID, err := models.GetUserIDFromToken(token, db)

// 	if err != nil {
// 		log.Print(err)
// 		rnd.Redirect("/", 404)
// 		return
// 	}

// 	access, err := models.GetAccess(userID, db)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	if access != "0" {
// 		rnd.Redirect("/")
// 		return
// 	}

// 	err = models.SetAccess(userID, db)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	rnd.Redirect("/confirmation")
// }
