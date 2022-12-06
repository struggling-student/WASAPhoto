package database

func (db *appdbimpl) SetLike(l Like) (Like, error) {
	_, err := db.c.Exec(`INSERT INTO likes (Id, userId, photoId, photoOwner) VALUES (?, ?, ?, ?)`, l.LikeId, l.UserIdentifier, l.PhotoIdentifier, l.PhotoOwner)
	if err != nil {
		return l, err
	}
	return l, nil
}
