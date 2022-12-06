package database

import (
	"database/sql"
)

func (db *appdbimpl) GetFollowId(f Follow) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT Id, followerId, userId FROM followers WHERE followerId=? AND userId = ?`, f.FollowedId, f.UserId).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeDoesNotExist
		}
	}
	return follow, nil
}
