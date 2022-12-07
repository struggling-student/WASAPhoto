package database

import "database/sql"

func (db *appdbimpl) SetComment(c Comment) (Comment, error) {
	_, err := db.c.Exec(`INSERT INTO comments (Id, userId, photoid, photoOwner, content) VALUES (?, ?, ?, ?, ?)`, c.Id, c.UserId, c.PhotoId, c.PhotoOwner, c.Content)
	if err != nil {
		return c, err
	}
	return c, nil
}

func (db *appdbimpl) RemoveComment(c Comment) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, c.Id)
	return err
}

func (db *appdbimpl) RemoveComments(user uint64, banned uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE userId=? AND photoOwner=?`, banned, user)
	return err
}

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

func (db *appdbimpl) GetCommentById(c Comment) (Comment, error) {
	var comment Comment
	if err := db.c.QueryRow(`SELECT id, userId, photoId, photoOwner, content FROM comments WHERE id = ?`, c.Id).Scan(&comment.Id, &comment.UserId, &comment.PhotoId, &comment.PhotoOwner, &comment.Content); err != nil {
		if err == sql.ErrNoRows {
			return comment, ErrLikeDoesNotExist
		}
	}
	return comment, nil
}

func (db *appdbimpl) GetCommentsCount(photoid uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM comments WHERE photoId = ?`, photoid).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeDoesNotExist
		}
	}
	return count, nil
}
