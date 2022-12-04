package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

type User struct {
	// Identifier is the unique identifier for the user
	Id uint64 `json:"id"`
	// Username is the username of the user
	Username string `json:"username"`
}
type Followers struct {
	// Identifier for the user that has the followers
	Id uint64 `json:"identifier"`
	// List of followers
	Followers []Follow `json:"Followers"`
}
type Follow struct {
	// BanIdentifier is the identifier for the ban action
	FollowId uint64 `json:"followId"`
	// Identifier for the user who is banned
	FollowedId uint64 `json:"followedId"`
	// Identifier for the user who is banning
	UserId uint64 `json:"userId"`
}
type Bans struct {
	// Identifier for the user that has the bans
	Identifier uint64 `json:"identifier"`
	// List of bans
	Bans []Ban `json:"bans"`
}
type Ban struct {
	// BanIdentifier is the identifier for the ban action
	BanId uint64 `json:"banId"`
	// Identifier for the user who is banned
	BannedId uint64 `json:"bannedId"`
	// Identifier for the user who is banning
	UserId uint64 `json:"userId"`
}
type Photos struct {
	// Identifier of the user who has the photos
	Identifier uint64 `json:"identifier"`
	// List of photos
	Photos []Photo `json:"photos"`
}

type Photo struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"userId"`
	File   string `json:"file"`
	Date   string `json:"date"`
}
type Likes struct {
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// Identifier for the user who liked the photo
	UserIdentifier uint64 `json:"identifier"`
	// List of likes under a photo
	Likes []Like `json:"likes"`
}
type Like struct {
	// Identifier for the like that has been added
	LikeId uint64 `json:"likeId"`
	// Identifier for the photo that has the likes
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// Identifier for the user who liked the photo
	UserIdentifier uint64 `json:"identifier"`
}

type Comments struct {
	// Identifier of the user who has commented
	Id uint64 `json:"id"`
	// Identifier for the photo that has the comments
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// List of comments under the photo
	Comments []Comment `json:"comments"`
}
type Comment struct {
	// Identifier of the user who has commented
	Id uint64 `json:"id"`
	// Identifier for the photo that has the comments
	PhotoId uint64 `json:"photoId"`
	// Identifier of the user who has commented
	UserId uint64 `json:"userId"`
	// Content of the comment
	Content string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// DONE
	CreateUser(User) (User, error)
	CreateBan(Ban) (Ban, error)
	GetUserById(User) (User, error)
	SetUsername(User) (User, error) // Review if the username is needed
	RemoveBan(Ban) error
	GetBanById(Ban) (Ban, error)
	GetBans(User) ([]Ban, error)
	//DB functions for follow
	SetFollow(Follow) (Follow, error)
	RemoveFollow(Follow) error
	GetFollowById(Follow) (Follow, error)
	GetFollowers(User) ([]Follow, error)
	// DB functions for photos

	// Insert a photo into the database. Returns the photo with the id, UserId, File and Date filled.
	SetPhoto(Photo) (Photo, error)
	// Remove a photo from the database. Returns an error if the photo cannot be deleted.
	RemovePhoto(Photo) error
	// Checks if a photo exists in the database.
	GetPhotoById(Photo) (Photo, error)

	GetPhotos(User) ([]Photo, error)

	// DB functions for likes
	// Insert a like into the database. Returns the like with the id, PhotoId, UserId filled.
	SetLike(Like) (Like, error)
	// Checks if a like exists in the database.
	GetLikeById(Like) (Like, error)
	// Remove a like from the database. Returns an error if the like cannot be deleted.
	RemoveLike(Like) error
	//
	GetLikes(Photo) ([]Like, error)

	// DB functions for comments
	SetComment(Comment) (Comment, error)
	GetCommentById(Comment) (Comment, error)
	RemoveComment(Comment) error
	GetComments(Photo) ([]Comment, error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	db.Exec("PRAGMA foreign_keys = ON")
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		usersDatabase := `CREATE TABLE users (
			Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT UNIQUE
			);`
		photosDatabase := `CREATE TABLE photos (
			Id INTEGER NOT NULL PRIMARY KEY, 
			userId INTEGER NOT NULL,
			photo TEXT,
			date TEXT,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		likesDatabase := `CREATE TABLE likes (
			Id INTEGER NOT NULL PRIMARY KEY,
			photoId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		commentsDatabase := `CREATE TABLE comments (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		bansDatabase := `CREATE TABLE bans (
			banId INTEGER NOT NULL PRIMARY KEY,
			bannedId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		followersDatabase := `CREATE TABLE followers (
			Id INTEGER NOT NULL PRIMARY KEY,
			followerId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		_, err = db.Exec(photosDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(likesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(bansDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

		_, err = db.Exec(followersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
