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

type Steam struct {
	Identifier uint64        `json:"identifier"`
	Photos     []PhotoStream `json:"photoStream"`
}

type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         string `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
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
	// Identifier for the user who liked the photo
	UserIdentifier uint64 `json:"identifier"`
	// Identifier for the photo that has the likes
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// Identifier for the user who has the photo
	PhotoOwner uint64 `json:"photoOwner"`
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
	// Identifier of the user who has commented
	UserId uint64 `json:"userId"`
	// Identifier for the photo that has the comments
	PhotoId uint64 `json:"photoId"`
	// Identifier for the user who owns the photo
	PhotoOwner uint64 `json:"photoOwner"`
	// Content of the comment
	Content string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// DB functions for users
	// CreateUser creates a new user. Returns the user identifier and an error if the operation failed.
	CreateUser(User) (User, error)
	//
	GetUserById(User) (User, error)
	//
	SetUsername(User) (User, error)
	//
	GetProfile(User) (User, error)
	// Get teh stream of photos for a user, returns a list of photos and an error if the operation failed.
	GetMyStream(User) ([]PhotoStream, error)
	GetCommentsCount(uint64) (int, error)
	GetLikesCount(photoid uint64) (int, error)
	GetFollowersCount(uint64) (int, error)
	GetFollowingsCount(uint64) (int, error)
	GetPhotosCount(uint64) (int, error)

	RemoveComments(uint64, uint64) error
	RemoveLikes(uint64, uint64) error
	GetFollowId(f Follow) (Follow, error)
	// DB functions for bans
	// Bans an user, returns the ban body and an error if the operation failed.
	CreateBan(Ban) (Ban, error)
	// Removes a ban, returns the ban body and an error if the operation failed.
	RemoveBan(Ban) error
	// Check if a ban exists, returns the ban body and an error if the operation failed.
	GetBanById(Ban) (Ban, error)
	// Get the list of bans for a user, returns a list of bans and an error if the operation failed.
	GetBans(User) ([]Ban, error)

	// DB functions for follow
	// Follows a user, returns the follow body and an error if the operation failed.
	SetFollow(Follow) (Follow, error)
	// Removes a follow from the database, returns an error if the operation failed.
	RemoveFollow(Follow) error
	// GetFollowById returns a follow by its id, returns an error if the operation failed.
	GetFollowById(Follow) (Follow, error)
	// GetFollowers returns a list of followers for a user, returns an error if the operation failed.s
	GetFollowers(User) ([]Follow, error)

	// DB functions for photos
	// Insert a photo into the database. Returns the photo with the id, UserId, File and Date filled.
	SetPhoto(Photo) (Photo, error)
	// Remove a photo from the database. Returns an error if the photo cannot be deleted.
	RemovePhoto(Photo) error
	// Checks if a photo exists in the database.
	GetPhotoById(Photo) (Photo, error)
	// Get the photos of a user. Returns an error if the user does not exist.
	GetPhotos(User) ([]Photo, error)

	// DB functions for likes
	// Insert a like into the database. Returns the like with the id, PhotoId, UserId filled.
	SetLike(Like) (Like, error)
	// Checks if a like exists in the database.
	GetLikeById(Like) (Like, error)
	// Remove a like from the database. Returns an error if the like cannot be deleted.
	RemoveLike(Like) error
	// Get all likes for a photo. Returns a list of likes.
	GetLikes(Photo) ([]Like, error)

	// DB functions for comments
	// Insert a comment into the database. Returns the comment with the id, PhotoId, UserId, Content filled.
	SetComment(Comment) (Comment, error)
	// Checks if a comment exists in the database.
	GetCommentById(Comment) (Comment, error)
	// Remove a comment from the database. Returns an error if the comment cannot be deleted.
	RemoveComment(Comment) error
	// GetComments returns all comments for a photo. Returns an error if the operation failed.
	GetComments(Photo) ([]Comment, error)

	// Other DB functions
	// Ping the database to check if it is alive. returns an error if the database is not alive.
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
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
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
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		commentsDatabase := `CREATE TABLE comments (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
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
