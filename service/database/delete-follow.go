package database

func (db *appdbimpl) RemoveFollow(f Follow) error {
	_, err := db.c.Exec(`DELETE FROM followers WHERE id=?`, f.FollowId)
	return err
}
