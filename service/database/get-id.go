package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUserById(u User) (User, error) {
	// var user User
	// rows, err := db.c.Query(`SELECT id, username FROM users WHERE username = ?`, u.Username)
	// if err != nil {
	// 	return user, ErrUserDoesNotExist
	// }
	// defer func() { _ = rows.Close() }()
	// for rows.Next() {
	// 	err = rows.Scan(&user.Id, &user.Username)
	// 	return user, err
	// }
	// if rows.Err() != nil {
	// 	return user, err
	// }

	// return user, nil
	var user User
	// Query for a value based on a single row.
	if err := db.c.QueryRow(`SELECT id, username FROM users WHERE username = ?`, u.Username).Scan(&user.Id, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrUserDoesNotExist
		}
	}
	return user, nil
}
