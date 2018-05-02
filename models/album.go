package models

import (
	"database/sql"
	"log"
)

type Album struct {
	AlbumID     string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Images      []*Image `json:"images,omitempty"`
}

func GetAlbums(db *sql.DB) ([]*Album, error) {
	rows, err := db.Query("SELECT * from album where album_id>1")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	albums := make([]*Album, 0)
	for rows.Next() {
		album := new(Album)
		err := rows.Scan(&album.AlbumID, &album.Title, &album.Description)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		album.Images = make([]*Image, 0)
		album.Images, err = GetAllImageFromAlbum(album.AlbumID, db)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		albums = append(albums, album)
		if err = rows.Err(); err != nil {
			log.Print(err)
			return nil, err
		}
	}
	return albums, nil
}
