package database

func (db *appdbimpl) GetComments(p Photo) ([]Comment, error) {
	var ret []Comment
	rows, err := db.c.Query(`SELECT Id, userId, photoId, photoOwner, content FROM comments WHERE photoId = ?`, p.Id)
	if err != nil {
		return ret, ErrUserDoesNotExist
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.Id, &c.UserId, &c.PhotoId, &c.PhotoOwner, &c.Content)
		if err != nil {
			return nil, err
		}
		ret = append(ret, c)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ret, nil
}
