package database

func (db *appdbimpl) SetLike(l Like) (Like, error) {
	_, err := db.c.Exec(`INSERT INTO likes (Id, photoId, userId) VALUES (?, ?, ?)`, l.LikeId, l.PhotoIdentifier, l.UserIdentifier)
	if err != nil {
		return l, err
	}
	return l, nil
}
