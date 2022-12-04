package database

func (db *appdbimpl) RemoveLike(l Like) error {
	_, err := db.c.Exec(`DELETE FROM likes WHERE id=?`, l.LikeId)
	return err
}
