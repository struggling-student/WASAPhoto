package database

func (db *appdbimpl) SetUser(username string, identifier int) error {
	_, err := db.c.Exec("INSERT INTO users(id, username) VALUES (?, ?)", identifier, username)
	return err
}
