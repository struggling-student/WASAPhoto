package database

func (db *appdbimpl) SetFollow(f Follow) (Follow, error) {
	_, err := db.c.Exec(`INSERT INTO followers (Id, followerId, userId ) VALUES (?, ?, ?)`, f.FollowId, f.FollowedId, f.UserId)
	if err != nil {
		return f, err
	}
	return f, nil
}
