package database

func (db *appdbimpl) RemovePhoto(p Photo) error {
	_, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, p.Id)
	return err
}
