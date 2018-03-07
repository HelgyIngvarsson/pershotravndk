package models

import (
	"database/sql"
)

type User struct {
	UserID       string
	Username     string
	Hashpassword []byte
	Access       int
}

func InsertUser(user *User, db *sql.DB) (string, error) {
	var userID string
	err := db.QueryRow("Insert into \"user\" (username, hashpassword) VALUES ($1,$2) RETURNING user_id", user.Username, user.Hashpassword).Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func SetAccess(user_id string, db *sql.DB) error {
	_, err := db.Exec("UPDATE public.\"user\" SET useraccess=$1 WHERE user_id = $2;",
		1, user_id)
	if err != nil {
		return err
	}
	return nil
}

func GetAccess(userID string, db *sql.DB) (string, error) {
	row := db.QueryRow("select useraccess from \"user\" where user_id =$1", userID)
	var access string
	err := row.Scan(&access)
	if err != nil {
		return "", err
	}
	return access, nil
}

func DeleteUser(userID string, db *sql.DB) error {
	_, err := db.Exec("delete from \"user\" where user_id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(username string, db *sql.DB) (*User, error) {
	row := db.QueryRow("select * from \"user\" where username =$1", username)
	user := new(User)
	err := row.Scan(&user.UserID, &user.Username, &user.Hashpassword, &user.Access)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func GetUserByID(userID string, db *sql.DB) (*User, error) {
	row := db.QueryRow("select * from \"user\" where user_id =$1", userID)
	user := new(User)
	err := row.Scan(&user.UserID, &user.Username, &user.Hashpassword, &user.Access)
	if err != nil {
		return nil, err
	}
	return user, nil
}
