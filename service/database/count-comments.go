package database

import (
	"database/sql"
)

func (db *appdbimpl) GetCommentsCount(photoid uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?`, photoid).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}
