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
