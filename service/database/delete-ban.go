package database

func (db *appdbimpl) RemoveBan(b Ban) error {
	_, err := db.c.Exec(`DELETE FROM bans WHERE banId=?`, b.BanId)
	return err
}
