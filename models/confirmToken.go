package models

import (
	"database/sql"
	"log"
)

func InsertToken(user_id string, token string, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO confirmation_token(user_id, token) VALUES ($1, $2);", user_id, token)
	if err != nil {
		return err
	}
	return nil
}

func GetUserIDFromToken(token string, db *sql.DB) (string, error) {

	row := db.QueryRow("SELECT user_id FROM confirmation_token WHERE token=$1", token)
	var userID string
	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	_, err = db.Exec("UPDATE confirmation_token SET date_used=CURRENT_DATE WHERE token = $1;", token)
	if err != nil {
		log.Print(err)
	}

	return userID, nil
}

func DeleteToken(userID string, db *sql.DB) error {
	_, err := db.Exec("delete from \"confirmation_token\" where user_id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
