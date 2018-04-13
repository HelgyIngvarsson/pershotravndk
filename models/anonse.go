package models

import (
	"database/sql"
)

type Anonse struct {
	AnonseID   string `json:"id,omitempty"`
	Body       string `json:"body,omitempty"`
	PostDate   string `json:"post_date,omitempty"`
	AnonseDate string `json:"anonse_date,omitempty"`
	UserID     string `json:"user_id,omitempty"`
}

func InsertAnonse(anonse *Anonse, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO public.anonse(body, anonse_date, user_id)VALUES ($1,$2,$3);",
		anonse.Body, anonse.AnonseDate, anonse.UserID)
	if err != nil {
		return err
	}
	return nil
}
func GetAllAnonses(db *sql.DB) ([]*Anonse, error) {
	rows, err := db.Query("select * from anonse")
	if err != nil {
		return nil, err
	}
	anonses := make([]*Anonse, 0)
	for rows.Next() {
		anonse := new(Anonse)
		err = rows.Scan(&anonse.AnonseID, &anonse.Body, &anonse.PostDate, &anonse.AnonseDate, &anonse.UserID)
		if err != nil {
			return nil, err
		}
		anonses = append(anonses, anonse)
		if err = rows.Err(); err != nil {
			return nil, err
		}

	}
	return anonses, nil
}
