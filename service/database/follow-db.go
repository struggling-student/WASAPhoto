package database

import "database/sql"

func (db *appdbimpl) SetFollow(f Follow) (Follow, error) {
	_, err := db.c.Exec(`INSERT INTO followers (Id, followerId, userId ) VALUES (?, ?, ?)`, f.FollowId, f.FollowedId, f.UserId)
	if err != nil {
		return f, err
	}
	return f, nil
}

func (db *appdbimpl) RemoveFollow(f Follow) error {
	_, err := db.c.Exec(`DELETE FROM followers WHERE id=?`, f.FollowId)
	return err
}

func (db *appdbimpl) GetFollowById(f Follow) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT id, followerId, userId FROM followers WHERE id = ?`, f.FollowId).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeDoesNotExist
		}
	}
	return follow, nil
}
func (db *appdbimpl) GetFollowId(f Follow) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT Id, followerId, userId FROM followers WHERE followerId=? AND userId = ?`, f.FollowedId, f.UserId).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeDoesNotExist
		}
	}
	return follow, nil
}

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

func (db *appdbimpl) GetFollowersCount(id uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE followerId = ?`, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}

func (db *appdbimpl) GetFollowingsCount(id uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE userId = ?`, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}
