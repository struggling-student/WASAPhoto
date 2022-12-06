package database

func (db *appdbimpl) RemoveComments(user uint64, banned uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE userId=? AND photoOwner=?`, banned, user)
	return err
}
