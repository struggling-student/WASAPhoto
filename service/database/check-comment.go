package database

import (
	"database/sql"
)

func (db *appdbimpl) GetCommentById(c Comment) (Comment, error) {
	var comment Comment
	if err := db.c.QueryRow(`SELECT id, userId, photoId, photoOwner, content FROM comments WHERE id = ?`, c.Id).Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.PhotoOwner, &comment.Content); err != nil {
		if err == sql.ErrNoRows {
			return comment, ErrLikeDoesNotExist
		}
	}
	return comment, nil
}
