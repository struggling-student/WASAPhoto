package database

func (db *appdbimpl) RemoveComment(c Comment) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, c.Id)
	return err
}
