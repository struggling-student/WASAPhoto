package database

func (db *appdbimpl) GetBans(u User) ([]Ban, error) {
	var ret []Ban
	rows, err := db.c.Query(`SELECT banId, bannedId, userId FROM bans WHERE userId = ?`, u.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var b Ban
		err = rows.Scan(&b.BanId, &b.BannedId, &b.UserId)
		if err != nil {
			return nil, err
		}

		ret = append(ret, b)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
