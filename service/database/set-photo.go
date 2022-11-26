package database

func (db *appdbimpl) SetPhoto(Username string, identifier uint64, file string) error {
	err := db.c.QueryRow("SELECT username FROM users WHERE username=?").Scan(&Username)
	return err
}
