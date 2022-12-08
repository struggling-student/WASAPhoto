package database

import (
	"database/sql"
)

// TODO decsription
func (db *appdbimpl) CreateUser(u User) (User, error) {
	res, err := db.c.Exec("INSERT INTO users(username) VALUES (?)", u.Username)
	if err != nil {
		return u, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.Id = uint64(lastInsertID)
	return u, nil
}

// TODO description
func (db *appdbimpl) SetUsername(u User, username string) (User, error) {
	res, err := db.c.Exec(`UPDATE users SET Username=? WHERE Id=? AND Username=?`, u.Username, u.Id, username)
	if err != nil {
		return u, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return u, err
	} else if affected == 0 {
		return u, ErrUserDoesNotExist
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
	rows, err := db.c.Query(`SELECT Id, userId, photo, date FROM photos WHERE userId IN (SELECT followerId FROM followers WHERE userId=? AND banStatus = 0) ORDER BY date`, u.Id)
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
		ret = append(ret, b)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}

func (db *appdbimpl) GetProfile(User) (User, error) {
	// panic("implement me")
	return User{}, nil
}
