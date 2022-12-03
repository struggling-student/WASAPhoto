package database

func (db *appdbimpl) SetPhoto(p Photo) (Photo, error) {
	_, err := db.c.Exec(`INSERT INTO photos (Id, userId, photo, date) VALUES (?, ?, ?, ?)`, p.Id, p.UserId, p.File, p.Date)
	if err != nil {
		return p, err
	}
	return p, nil
}
