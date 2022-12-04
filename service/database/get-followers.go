package database

func (db *appdbimpl) GetFollowers(u User) ([]Follow, error) {
	var ret []Follow
	rows, err := db.c.Query(`SELECT id, followerId, userId FROM followers WHERE userId = ?`, u.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var f Follow
		err = rows.Scan(&f.FollowId, &f.FollowedId, &f.UserId)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ret, nil
}
