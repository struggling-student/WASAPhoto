package database

import (
	"database/sql"
)

func (db *appdbimpl) GetLikeById(l Like) (Like, error) {
	var like Like
	if err := db.c.QueryRow(`SELECT Id, userId, photoId, photoOwner FROM likes WHERE id = ?`, l.LikeId).Scan(&like.LikeId, &like.UserIdentifier, &like.PhotoIdentifier, &like.PhotoOwner); err != nil {
		if err == sql.ErrNoRows {
			return like, ErrLikeDoesNotExist
		}
	}
	return like, nil
}
