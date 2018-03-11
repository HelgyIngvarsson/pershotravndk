package models

import (
	"database/sql"
)

type Image struct {
	ImageID string
	Path    string
	Album   string
}

func GetImageByID(imageID string, db *sql.DB) (*Image, error) {
	row := db.QueryRow("Select * from image where image_id =$1", imageID)
	image := new(Image)
	err := row.Scan(&image.ImageID, &image.Path, &image.Album)
	if err != nil {
		return nil, err
	}
	return image, nil
}
