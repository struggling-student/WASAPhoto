package database

import "database/sql"

func (db *appdbimpl) CreateBan(b Ban) (Ban, error) {
	_, err := db.c.Exec(`INSERT INTO bans (banId, bannedId, userId ) VALUES (?, ?, ?)`, b.BanId, b.BannedId, b.UserId)
	if err != nil {
		return b, err
	}
	return b, nil
}

func (db *appdbimpl) RemoveBan(b Ban) error {
	res, err := db.c.Exec(`DELETE FROM bans WHERE banId=? AND userId=? AND bannedId = ?`, b.BanId, b.UserId, b.BannedId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanDoesNotExist
	}
	return nil
}

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

func (db *appdbimpl) GetBanById(b Ban) (Ban, error) {
	var ban Ban
	if err := db.c.QueryRow(`SELECT banId, bannedId, userId FROM bans WHERE banId = ?`, b.BanId).Scan(&ban.BanId, &ban.BannedId, &ban.UserId); err != nil {
		if err == sql.ErrNoRows {
			return ban, ErrLikeDoesNotExist
		}
	}
	return ban, nil
}

func (db *appdbimpl) UpdateBanStatus(status int, followerId uint64, userId uint64) error {
	res, err := db.c.Exec(`UPDATE followers SET banStatus=? WHERE followerId=? AND userId=?`, status, followerId, userId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanDoesNotExist
	}
	return nil
}
