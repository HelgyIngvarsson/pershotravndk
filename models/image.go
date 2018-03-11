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
func GetAllImageFromAlbum(albumID string, db *sql.DB) ([]*Image, error) {

	rows, err := db.Query("SELECT * from image where album = $1", albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	images := make([]*Image, 0)
	for rows.Next() {
		image := new(Image)
		err := rows.Scan(&image.ImageID, &image.Path, &image.Album)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return images, nil
}
