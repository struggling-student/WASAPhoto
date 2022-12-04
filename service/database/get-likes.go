package database

func (db *appdbimpl) GetLikes(p Photo) ([]Like, error) {
	var ret []Like
	rows, err := db.c.Query(`SELECT id, photoId, userId FROM likes WHERE photoId = ?`, p.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var l Like
		err = rows.Scan(&l.LikeId, &l.PhotoIdentifier, &l.UserIdentifier)
		if err != nil {
			return nil, err
		}
		ret = append(ret, l)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ret, nil
}
