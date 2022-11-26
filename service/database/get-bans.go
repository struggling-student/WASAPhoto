package database

func (db *appdbimpl) GetBans(Token int) ([]Ban, error) {
	var ret []Ban
	rows, err := db.c.Query(`SELECT id,username,token FROM WHERE token=?`, Token)
	if err != nil {
		return ret, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var ban Ban
		err = rows.Scan(&ban.BanIdentifier, &ban.Username, &ban.Identifier)
		if err != nil {
			return ret, err
		}
		ret = append(ret, ban)
	}
	if rows.Err() != nil {
		return ret, err
	}

	return ret, nil
}
