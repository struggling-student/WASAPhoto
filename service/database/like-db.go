package database

import (
	"database/sql"
)

func (db *appdbimpl) SetLike(l Like) (Like, error) {
	_, err := db.c.Exec(`INSERT INTO likes (Id, userId, photoId, photoOwner) VALUES (?, ?, ?, ?)`, l.LikeId, l.UserIdentifier, l.PhotoIdentifier, l.PhotoOwner)
	if err != nil {
		return l, err
	}
	return l, nil
}

func (db *appdbimpl) RemoveLike(l Like) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE id=? AND userId=? AND photoId = ? AND photoOwner = ?`, l.LikeId, l.UserIdentifier, l.PhotoIdentifier, l.PhotoOwner)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrLikeDoesNotExist
	}
	return nil
}

func (db *appdbimpl) RemoveLikes(user uint64, banned uint64) error {
	_, err := db.c.Exec(`DELETE FROM likes WHERE userId=? AND photoOwner=?`, banned, user)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) GetLike(photoid uint64, token uint64) (Like, error) {
	var like Like

	if err := db.c.QueryRow(`SELECT Id, userId, photoId, photoOwner FROM likes WHERE userId = ? AND photoId = ?`, token, photoid).Scan(&like.LikeId, &like.UserIdentifier, &like.PhotoIdentifier, &like.PhotoOwner); err != nil {
		if err == sql.ErrNoRows {
			return like, ErrLikeDoesNotExist
		}
	}
	return like, nil
}

func (db *appdbimpl) GetLikeById(l Like) (Like, error) {
	var like Like
	if err := db.c.QueryRow(`SELECT Id, userId, photoId, photoOwner FROM likes WHERE id = ?`, l.LikeId).Scan(&like.LikeId, &like.UserIdentifier, &like.PhotoIdentifier, &like.PhotoOwner); err != nil {
		if err == sql.ErrNoRows {
			return like, ErrLikeDoesNotExist
		}
	}
	return like, nil
}

func (db *appdbimpl) GetLikesCount(photoid uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM likes WHERE photoId = ?`, photoid).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}
