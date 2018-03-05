package models

import (
	"database/sql"
)

type Profile struct {
	ProfileID   string
	Name        string
	Sername     string
	Email       string
	Description string
	Mailing     bool
	UserID      string
}

func InsertProfile(profile *Profile, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO \"profile\"(name, sername, email, mailing, description, user_id) VALUES ($1, $2, $3, $4, $5, $6)",
		profile.Name, profile.Sername, profile.Email, profile.Mailing, profile.Description, profile.UserID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteProfile(userID string, db *sql.DB) error {
	_, err := db.Exec("delete from \"profile\" where user_id = $1", userID)
	if err != nil {
		return err
	}
	return nil
}
