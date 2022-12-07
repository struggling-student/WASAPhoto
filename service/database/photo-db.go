package database

import "database/sql"

func (db *appdbimpl) SetPhoto(p Photo) (Photo, error) {
	_, err := db.c.Exec(`INSERT INTO photos (Id, userId, photo, date) VALUES (?, ?, ?, ?)`, p.Id, p.UserId, p.File, p.Date)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (db *appdbimpl) RemovePhoto(id uint64) error {
	res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, id)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPhotoDoesNotExist
	}
	return nil
}

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

func (db *appdbimpl) GetPhotosCount(id uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM photos WHERE userId = ?`, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}
