package database

import (
	"database/sql"
)

func (db *appdbimpl) GetPhotoById(p Photo) (Photo, error) {
	var photo Photo
	if err := db.c.QueryRow(`SELECT id, userId, photo, date FROM photos WHERE id = ?`, p.Id).Scan(&photo.Id, &photo.UserId, &photo.File, &photo.Date); err != nil {
		if err == sql.ErrNoRows {
			return photo, ErrPhotoDoesNotExist
		}
	}
	return photo, nil
}
