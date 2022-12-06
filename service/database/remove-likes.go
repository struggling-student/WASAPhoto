package database

func (db *appdbimpl) RemoveLikes(user uint64, banned uint64) error {
	_, err := db.c.Exec(`DELETE FROM likes WHERE userId=? AND photoOwner=?`, banned, user)
	return err
}
