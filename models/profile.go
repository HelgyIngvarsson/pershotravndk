package models

import (
	"database/sql"
)

type Profile struct {
	ProfileID   string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Sername     string `json:"sername,omitempty"`
	Email       string `json:"email,omitempty"`
	Avatar      *Image `json:"image,omitempty"`
	Description string `json:"desc,omitempty"`
	Mailing     bool   `json:"mailing,omitempty"`
	UserID      string `json:"user_id,omitempty"`
}

func InsertProfile(profile *Profile, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO \"profile\"(name, sername, email, mailing, description, user_id, avatar_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		profile.Name, profile.Sername, profile.Email, profile.Mailing, profile.Description, profile.UserID, profile.Avatar.ImageID)
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

func UpdateProfie(profile *Profile, db *sql.DB) error {
	_, err := db.Exec("UPDATE public.profile SET profile_id=$1, name=$2, sername=$3, email=$4, mailing=$5, description=$6 WHERE user_id=$7;",
		profile.ProfileID, profile.Name, profile.Sername, profile.Email, profile.Mailing, profile.Description, profile.UserID)
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
func GetAdminsProfile(db *sql.DB) ([]*Profile, error) {
	rows, err := db.Query("SELECT \"profile\".* from \"profile\",\"user\" where \"profile\".user_id =\"user\".user_id and \"user\".useraccess=1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := make([]*Profile, 0)
	for rows.Next() {
		profile := new(Profile)
		profile.Avatar = new(Image)
		err := rows.Scan(&profile.ProfileID, &profile.Name, &profile.Sername, &profile.Email,
			&profile.Mailing, &profile.Description, &profile.UserID, &profile.Avatar.ImageID)
		if err != nil {
			return nil, err
		}
		profile.Avatar, err = GetImageByID(profile.Avatar.ImageID, db)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return profiles, nil
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
