package database

func (db *appdbimpl) SetBan(Username string, token int, banIdentifier int) error {
	_, err := db.c.Exec("INSERT INTO bans(id, username,token) VALUES (?, ?, ?)", banIdentifier, Username, token)
	return err
}
