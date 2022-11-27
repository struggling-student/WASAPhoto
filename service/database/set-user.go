package database

func (db *appdbimpl) CreateUser(u User) (User, error) {
	_, err := db.c.Exec("INSERT INTO users(identifier, username) VALUES (?, ?)", u.Identifier, u.Username)
	if err != nil {
		return u, err
	}
	return u, nil
}
