package database

import (
	"database/sql"
)

func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", u.Username)
	if err != nil {
		var user User
		if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, u.Username).Scan(&user.Id, &user.Username); err != nil {
			if err == sql.ErrNoRows {
				return user, ErrUserDoesNotExist
			}
		}
		return user, nil
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.Id = uint64(lastInsertID)
	return u, nil
}

func (db *appdbimpl) SetUsername(u User, username string) (User, error) {
	res, err := db.c.Exec(`UPDATE users SET Username=? WHERE Id=? AND Username=?`, u.Username, u.Id, username)
	if err != nil {
		return u, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return u, err
	} else if affected == 0 {
		return u, err
	}
	return u, nil
}

func (db *appdbimpl) GetUserId(username string) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

func (db *appdbimpl) CheckUserByUsername(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, u.Username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

func (db *appdbimpl) CheckUserById(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE id = ?`, u.Id).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

func (db *appdbimpl) CheckUser(u User) (User, error) {
	var user User
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE id = ? AND username = ?`, u.Id, u.Username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}

func (db *appdbimpl) GetMyStream(u User) ([]PhotoStream, error) {
	var ret []PhotoStream
	rows, err := db.c.Query(`SELECT Id, userId, photo, date FROM photos WHERE userId IN (SELECT followerId FROM followers WHERE userId=? AND followerId NOT IN (SELECT userId FROM bans WHERE bannedId=?))`, u.Id, u.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var b PhotoStream
		err = rows.Scan(&b.Id, &b.UserId, &b.File, &b.Date)
		if err != nil {
			return nil, err
		}
		if err := db.c.QueryRow(`SELECT username FROM users WHERE id = ?`, b.UserId).Scan(&b.Username); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		if err := db.c.QueryRow(`SELECT COUNT(*) FROM likes WHERE photoId = ?`, b.Id).Scan(&b.LikeCount); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		if err := db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?`, b.Id).Scan(&b.CommentCount); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		if err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM likes WHERE userId = ? AND photoId = ?)`, u.Id, b.Id).Scan(&b.LikeStatus); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
		}
		ret = append(ret, b)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ret, nil
}

func (db *appdbimpl) GetFollowStatus(r uint64, u uint64) (bool, error) {
	var ret bool
	if err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM followers WHERE userId= ? AND  followerId= ?)`, r, u).Scan(&ret); err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
	}
	return ret, nil
}

func (db *appdbimpl) GetBanStatus(r uint64, u uint64) (bool, error) {
	var ret bool
	if err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM bans WHERE userId=? AND bannedId=?)`, r, u).Scan(&ret); err != nil {
		if err == sql.ErrNoRows {
			return false, err
		}
	}
	return ret, nil
}
