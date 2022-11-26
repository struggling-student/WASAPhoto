package database

func (db *appdbimpl) SetUsername(username string, newUsername string) (int, error) {
	var identifier int
	_, err := db.c.Exec("UPDATE users SET username=? WHERE username=?", newUsername, username)
	if err != nil {
		return identifier, err
	}
	row, err := db.c.Query(`SELECT id FROM users WHERE username=?`, newUsername)
	for row.Next() {
		err = row.Scan(&identifier)
		if err != nil {
			return identifier, err
		}
	}
	return identifier, nil
}
