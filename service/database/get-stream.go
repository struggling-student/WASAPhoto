package database

func (db *appdbimpl) GetMyStream(u User) ([]PhotoStream, error) {
	var ret []PhotoStream
	rows, err := db.c.Query(`SELECT Id, userId, photo, date FROM photos WHERE userId IN (SELECT followerId FROM followers WHERE userId=?) ORDER BY date`, u.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var b PhotoStream
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
