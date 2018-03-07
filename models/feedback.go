package models

import (
	"database/sql"
)

type Feedback struct {
	ID      string
	Message string
	UserID  string
	User    *User
	Date    string
}

func InsertFeedback(feedback *Feedback, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO feedback(message, user_id)VALUES ($1,$2);", feedback.Message, feedback.UserID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllFeedbacks(db *sql.DB) ([]*Feedback, error) {
	rows, err := db.Query("select * from feedback")
	if err != nil {
		return nil, err
	}
	feedbacks := make([]*Feedback, 0)
	for rows.Next() {
		feedback := new(Feedback)
		err = rows.Scan(&feedback.ID, &feedback.Message, &feedback.Date, &feedback.UserID)
		if err != nil {
			return nil, err
		}
		feedback.User, _ = GetUserByID(feedback.UserID, db)
		feedbacks = append(feedbacks, feedback)
		if err = rows.Err(); err != nil {
			return nil, err
		}

	}
	return feedbacks, nil
}
