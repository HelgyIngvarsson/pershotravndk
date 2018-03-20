package models

import (
	"database/sql"
)

type Comment struct {
	CommentID string
	UserID    string
	Profile   *Profile
	ArticleID string
	Body      string
	Date      string
}

func InsertComment(comment *Comment, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO public.comment(article_id, user_id, body) VALUES ($1, $2, $3);",
		comment.ArticleID, comment.UserID, comment.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentsByArticleID(articleID string, db *sql.DB) ([]*Comment, error) {
	rows, err := db.Query("Select * from comment where article_id = $1", articleID)
	if err != nil {
		return nil, err
	}
	comments := make([]*Comment, 0)
	for rows.Next() {
		comment := new(Comment)
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.UserID, &comment.Body, &comment.Date)
		if err != nil {
			return nil, err
		}
		comment.Profile = new(Profile)
		comment.Profile, err = GetProfileByUserID(comment.UserID, db)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return comments, nil
}
