package database

import (
	"database/sql"
)

func (db *appdbimpl) GetFollowById(f Follow) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT id, followerId, userId FROM followers WHERE id = ?`, f.FollowId).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeDoesNotExist
		}
	}
	return follow, nil
}
