package models

import (
	"database/sql"
)

type Profile struct {
	ProfileID   string
	Name        string
	Sername     string
	Email       string
	Avatar      *Image
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

func GetProfileByUserID(userID string, db *sql.DB) (*Profile, error) {
	row := db.QueryRow("Select * from \"profile\" where user_id =$1 ", userID)
	profile := new(Profile)
	profile.Avatar = new(Image)
	err := row.Scan(&profile.ProfileID, &profile.Name, &profile.Sername, &profile.Email,
		&profile.Mailing, &profile.Description, &profile.UserID, &profile.Avatar.ImageID)
	if err != nil {
		return nil, err
	}
	profile.Avatar, err = GetImageByID(profile.Avatar.ImageID, db)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
func GetNameByUserID(userID string, db *sql.DB) (string, error) {
	row := db.QueryRow("Select name,sername from \"profile\" where user_id =$1 ", userID)
	var name, sername string
	err := row.Scan(&name, &sername)
	if err != nil {
		return "Anon", err
	}
	return sername + " " + name, nil
}
