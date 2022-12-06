package database

func (db *appdbimpl) SetComment(c Comment) (Comment, error) {
	_, err := db.c.Exec(`INSERT INTO comments (Id, userId, photoid, photoOwner, content) VALUES (?, ?, ?, ?, ?)`, c.Id, c.UserId, c.PhotoId, c.PhotoOwner, c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}
