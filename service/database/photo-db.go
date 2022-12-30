package database

import (
	"database/sql"
)

func (db *appdbimpl) SetPhoto(p Photo) (Photo, error) {
	_, err := db.c.Exec(`INSERT INTO photos (Id, userId, photo, date) VALUES (?, ?, ?, ?)`, p.Id, p.UserId, p.File, p.Date)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (db *appdbimpl) RemovePhoto(id uint64) error {
	res1, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, id)
	if err != nil {
		return err
	}
	affected, err := res1.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return nil
	}
	res2, err := db.c.Exec(`DELETE FROM likes WHERE photoId=?`, id)
	if err != nil {
		return err
	}
	affected2, err := res2.RowsAffected()
	if err != nil {
	} else if affected2 == 0 {
		return nil
	}
	res3, err := db.c.Exec(`DELETE FROM comments WHERE photoId=?`, id)
	if err != nil {
		return err
	}
	affected3, err := res3.RowsAffected()
	if err != nil {
		return err
	} else if affected3 == 0 {
		return nil
	}
	return nil
}

func (db *appdbimpl) GetPhotos(u User, token uint64) ([]Photo, error) {
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
		if err := db.c.QueryRow(`SELECT COUNT(*) FROM likes WHERE photoId = ?`, b.Id).Scan(&b.LikesCount); err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrLikeDoesNotExist
			}
		}
		if err := db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?`, b.Id).Scan(&b.CommentsCount); err != nil {
			if err == sql.ErrNoRows {
				return nil, ErrLikeDoesNotExist
			}
		}
		if err := db.c.QueryRow(`SELECT EXISTS(SELECT 1 FROM likes WHERE userId = ? AND photoId = ?)`, token, b.Id).Scan(&b.LikeStatus); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			}
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

func (db *appdbimpl) CheckPhoto(p Photo) (Photo, error) {
	var photo Photo
	if err := db.c.QueryRow(`SELECT Id, userId, photo, date FROM photos WHERE Id=?`, p.Id).Scan(&photo.Id, &photo.UserId, &photo.File, &photo.Date); err != nil {
		if err == sql.ErrNoRows {
			return photo, ErrUserDoesNotExist
		}
	}
	return photo, nil
}
