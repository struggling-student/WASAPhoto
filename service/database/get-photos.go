package database

func (db *appdbimpl) GetPhotos(u User) ([]Photo, error) {
	var ret []Photo
	rows, err := db.c.Query(`SELECT id, userId, photo, date FROM photos WHERE userId = ?`, u.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var b Photo
		err = rows.Scan(&b.Id, &b.UserId, &b.File, &b.Date)
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
