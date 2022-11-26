package database

import "fmt"

func (db *appdbimpl) GetUser(username string) (User, error) {
	//var ret []User
	var user User
	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT id, username FROM users`)
	if err != nil {
		return user, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		//var f User
		err = rows.Scan(&user.Identifier, &user.Username)
		if err != nil {
			return user, err
		}

		//ret = append(ret, f)
	}
	fmt.Printf("%+v\n", user)
	if rows.Err() != nil {
		return user, err
	}

	return user, nil

}
