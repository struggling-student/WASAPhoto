package database

func (db *appdbimpl) RemoveBan(banIdentifier int) error {
	_, err := db.c.Exec(`DELETE FROM bans WHERE id=?`, banIdentifier)
	return err
}
