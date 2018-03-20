package models

import (
	"database/sql"
)

type Article struct {
	ArticleID string
	Title     string
	Body      string
	UserID    string
	Name      string
	Image     *Image
	Date      string
}

func InsertArticle(article *Article, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO public.article(title, body, user_id, image)VALUES ($1,$2,$3,$4);",
		article.Title, article.Body, article.UserID, article.Image.ImageID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllArticles(db *sql.DB) ([]*Article, error) {
	rows, err := db.Query("select * from article")
	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)
	for rows.Next() {
		article := new(Article)
		article.Image = new(Image)
		err = rows.Scan(&article.ArticleID, &article.Title, &article.Body, &article.UserID, &article.Image.ImageID, &article.Date)
		if err != nil {
			return nil, err
		}
		article.Name, _ = GetNameByUserID(article.UserID, db)
		article.Image, err = GetImageByID(article.Image.ImageID, db)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
		if err = rows.Err(); err != nil {
			return nil, err
		}
	}
	return articles, nil
}

func GetArticleByID(id string, db *sql.DB) (*Article, error) {
	row := db.QueryRow("Select * from article where article_id=$1", id)
	article := new(Article)
	article.Image = new(Image)
	err := row.Scan(&article.ArticleID, &article.Title, &article.Body, &article.UserID, &article.Image.ImageID, &article.Date)
	if err != nil {
		return nil, err
	}
	article.Name, _ = GetNameByUserID(article.UserID, db)
	article.Image, err = GetImageByID(article.Image.ImageID, db)
	if err != nil {
		return nil, err
	}
	return article, nil
}
