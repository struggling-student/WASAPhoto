package database

import (
	"database/sql"
)

func (db *appdbimpl) GetLikeById(l Like) (Like, error) {
	var like Like
	if err := db.c.QueryRow(`SELECT id, photoId, userId FROM likes WHERE id = ?`, l.LikeId).Scan(&like.LikeId, &like.PhotoIdentifier, &like.UserIdentifier); err != nil {
		if err == sql.ErrNoRows {
			return like, ErrLikeDoesNotExist
		}
	}
	return like, nil
}
