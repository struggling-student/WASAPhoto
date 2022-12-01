package database

func (db *appdbimpl) CreateBan(b Ban) (Ban, error) {
	_, err := db.c.Exec(`INSERT INTO bans (banId, bannedId, userId ) VALUES (?, ?, ?)`, b.BanId, b.BannedId, b.UserId)
	if err != nil {
		return b, err
	}
	return b, nil
}
