package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// List of errors that can be returned by the database.
var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrBanDoesNotExist = errors.New("Ban does not exist")
var ErrFollowDoesNotExist = errors.New("Follow does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// Struct that represents an user in the database.
type User struct {
	// Identifier is the unique identifier for the user
	Id uint64 `json:"id"`
	// Username is the username of the user
	Username string `json:"username"`
}

// Struct that represents the photo stream of an user in the database.
type Steam struct {
	// Identifier of the user stream
	Identifier uint64 `json:"identifier"`
	// Stream of photos
	Photos []PhotoStream `json:"photoStream"`
}

// Struct that represents a photo for the stream in the database.
type PhotoStream struct {
	// Identifier for the photo
	Id uint64 `json:"id"`
	// Identifier for the user who owns the photo
	UserId uint64 `json:"userId"`
	// Username for the user who owns the photo
	Username string `json:"username"`
	// File for the  photo
	File []byte `json:"file"`
	// Date when the photo was uploaded
	Date string `json:"date"`
	// Number of likes for the photo
	LikeCount int `json:"likeCount"`
	// Number of comments for the photo
	CommentCount int `json:"commentCount"`

	LikeStatus bool `json:"likeStatus"`
}

// Struct that represents the followers of an user in the database.
type Followers struct {
	// Identifier for the user that has the followers
	Id uint64 `json:"identifier"`
	// List of followers
	Followers []Follow `json:"Followers"`
}

// Struct that represents the follow of an user in the database.
type Follow struct {
	// Follow is the identifier for the follow action
	FollowId uint64 `json:"followId"`
	// Identifier for the user who is followed
	FollowedId uint64 `json:"followedId"`
	// Identifier for the user who is following
	UserId uint64 `json:"userId"`
	// Ban status for the user who is followed
	// If ban status is 1, the user is banned so it's not considered in the follow list.
	BanStatus int `json:"banStatus"`
}

// Struct that represents the bans of an user in the database.
type Bans struct {
	// Identifier for the user that has the bans
	Identifier uint64 `json:"identifier"`
	// Username for the user that has the bans
	Username string `json:"username"`
	// List of bans
	Bans []Ban `json:"bans"`
}

// Struct that represents the ban of an user in the database.
type Ban struct {
	// BanIdentifier is the identifier for the ban action
	BanId uint64 `json:"banId"`
	// Identifier for the user who is banned
	BannedId uint64 `json:"bannedId"`
	// Identifier for the user who is banning
	UserId uint64 `json:"userId"`
}

// Struct that represents the photos of an user in the database.
type Photos struct {
	// Identifier for the user that has requested the photos
	RequestUser uint64 `json:"requestUser"`
	// Identifier of the user who has the photos
	Identifier uint64 `json:"identifier"`
	// List of photos
	Photos []Photo `json:"photos"`
}

// Struct that represents a photo in the database.
type Photo struct {
	// Identifier for the photo
	Id uint64 `json:"id"`
	// Identifier for the user who owns the photo
	UserId uint64 `json:"userId"`
	// File for the  photo
	File []byte `json:"file"`
	// Date when the photo was uploaded
	Date string `json:"date"`

	LikesCount int `json:"likesCount"`

	CommentsCount int `json:"commentsCount"`
}

// // Struct that represents the likes of a photo in the database.
// type Likes struct {
// 	// Identifier for the user that has requested the likes
// 	RequestIdentifier uint64 `json:"requestIdentifier"`
// 	// Identifier for the photo that has the likes
// 	PhotoIdentifier uint64 `json:"photoIdentifier"`
// 	// Identifier for the owner of the photo
// 	PhotoOwner uint64 `json:"identifier"`
// 	// List of likes under a photo
// 	Likes []Like `json:"likes"`
// }

// Struct that represents a like in the database.
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

// Struct that represents the comments of a photo in the database.
type Comments struct {
	// Identifier for the user that has requested the comments
	RequestIdentifier uint64 `json:"requestIdentifier"`
	// Identifier for the photo that has the likes
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// Identifier for the owner of the photo
	PhotoOwner uint64 `json:"identifier"`
	// List of likes under a photo
	Comments []Comment `json:"comments"`
}

// Struct that represents a comment in the database.
type Comment struct {
	// Identifier of the user who has commented
	Id uint64 `json:"id"`
	// Identifier of the user who has commented
	UserId uint64 `json:"userId"`
	// Identifier for the photo that has the comments
	PhotoId uint64 `json:"photoId"`
	// Identifier for the user who owns the photo
	PhotoOwner uint64 `json:"photoOwner"`

	OwnerUsername string `json:"ownerUsername"`

	Username string `json:"username"`
	// Content of the comment
	Content string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// DB functions for users
	CreateUser(User) (User, error)
	SetUsername(User, string) (User, error)
	GetUserId(string) (User, error)
	CheckUserById(User) (User, error)
	CheckUserByUsername(User) (User, error)
	CheckUser(User) (User, error)
	GetMyStream(User) ([]PhotoStream, error)

	// DB functions for followers
	SetFollow(Follow) (Follow, error)
	RemoveFollow(uint64, uint64, uint64) error
	GetFollowingId(user1 uint64, user2 uint64) (Follow, error)
	GetFollowers(User, uint64) (Follow, error)
	GetFollowersCount(uint64) (int, error)
	GetFollowingsCount(uint64) (int, error)
	GetFollowStatus(uint64, uint64) (bool, error)

	// DB functions for bans
	CreateBan(Ban) (Ban, error)
	RemoveBan(Ban) error
	GetBans(User, uint64) (Ban, error)
	GetBanById(Ban) (Ban, error)
	UpdateBanStatus(int, uint64, uint64) error
	GetBanStatus(uint64, uint64) (bool, error)

	// DB functions for photos
	SetPhoto(Photo) (Photo, error)
	RemovePhoto(uint64) error
	GetPhotos(User) ([]Photo, error)
	GetPhotosCount(uint64) (int, error)
	CheckPhoto(Photo) (Photo, error)

	// DB functions for likes
	SetLike(Like) (Like, error)
	RemoveLike(Like) error
	RemoveLikes(uint64, uint64) error
	GetLike(uint64, uint64) (Like, error)
	GetLikeById(Like) (Like, error)
	GetLikesCount(photoid uint64) (int, error)

	// DB functions for comments
	SetComment(Comment) (Comment, error)
	RemoveComment(Comment) error
	RemoveComments(uint64, uint64) error
	GetComments(photoid uint64) ([]Comment, error)
	GetCommentById(Comment) (Comment, error)
	GetCommentsCount(uint64) (int, error)

	// Other functions
	Ping() error
}

// AppDatabaseImpl is the implementation of the AppDatabase interface
type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	// Check if db is nil
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	// Enable foreign keys
	_, err := db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Create the users table
		usersDatabase := `CREATE TABLE users (
			Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT UNIQUE
			);`
		// Create the users table
		photosDatabase := `CREATE TABLE photos (
			Id INTEGER NOT NULL PRIMARY KEY, 
			userId INTEGER NOT NULL,
			photo BLOB,
			date TEXT,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		// Create the followers table
		likesDatabase := `CREATE TABLE likes (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		// Create the comments table
		commentsDatabase := `CREATE TABLE comments (
			Id INTEGER NOT NULL PRIMARY KEY,
			userId INTEGER NOT NULL,
			photoId INTEGER NOT NULL,
			photoOwner INTEGER NOT NULL,
			content TEXT NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id),
			FOREIGN KEY (photoId) REFERENCES photos(Id)
			);`
		// Create the bans table
		bansDatabase := `CREATE TABLE bans (
			banId INTEGER NOT NULL PRIMARY KEY,
			bannedId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		// Create the followers table
		followersDatabase := `CREATE TABLE followers (
			Id INTEGER NOT NULL PRIMARY KEY,
			followerId INTEGER NOT NULL,
			userId INTEGER NOT NULL,
			banStatus INTEGER NOT NULL,
			FOREIGN KEY (userId) REFERENCES users(Id)
			);`
		// check error
		_, err = db.Exec(usersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// check error
		_, err = db.Exec(photosDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// check error
		_, err = db.Exec(likesDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// check error
		_, err = db.Exec(commentsDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// check error
		_, err = db.Exec(bansDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
		// check error
		_, err = db.Exec(followersDatabase)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// Ping checks the connection to the database.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
