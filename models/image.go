package models

import (
	"database/sql"
)

type Image struct {
	ImageID string `json:"id,omitempty"`
	Path    string `json:"path,omitempty"`
	Album   string `json:"album_id,omitempty"`
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

func InsertEmptyImage(db *sql.DB) (string, error) {
	var imageID string
	err := db.QueryRow("Insert into public.image(path,album) values(default,default) RETURNING image_id").Scan(&imageID)
	if err != nil {
		return "", err
	}
	return imageID, nil
}
func UpdateImage(image *Image, db *sql.DB) error {
	_, err := db.Exec("UPDATE public.image SET path=$1 WHERE image_id=$2;",
		image.Path, image.ImageID)
	if err != nil {
		return err
	}
	return nil
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
