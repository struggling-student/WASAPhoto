package database

import (
	"database/sql"
)

func (db *appdbimpl) GetBanById(b Ban) (Ban, error) {
	var ban Ban
	if err := db.c.QueryRow(`SELECT banId, bannedId, userId FROM bans WHERE banId = ?`, b.BanId).Scan(&ban.BanId, &ban.BannedId, &ban.UserId); err != nil {
		if err == sql.ErrNoRows {
			return ban, ErrLikeDoesNotExist
		}
	}
	return ban, nil
}
