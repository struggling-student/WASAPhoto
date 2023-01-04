package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrBanDoesNotExist = errors.New("Ban does not exist")
var ErrFollowDoesNotExist = errors.New("Follow does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

type Steam struct {
	Identifier uint64        `json:"identifier"`
	Photos     []PhotoStream `json:"photoStream"`
}

type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	Username     string `json:"username"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
	LikeStatus   bool   `json:"likeStatus"`
}

type Followers struct {
	Id        uint64   `json:"identifier"`
	Followers []Follow `json:"Followers"`
}

type Follow struct {
	FollowId   uint64 `json:"followId"`
	FollowedId uint64 `json:"followedId"`
	UserId     uint64 `json:"userId"`
}

type Bans struct {
	Identifier uint64 `json:"identifier"`
	Username   string `json:"username"`
	Bans       []Ban  `json:"bans"`
}

type Ban struct {
	BanId    uint64 `json:"banId"`
	BannedId uint64 `json:"bannedId"`
	UserId   uint64 `json:"userId"`
}

type Photos struct {
	RequestUser uint64  `json:"requestUser"`
	Identifier  uint64  `json:"identifier"`
	Photos      []Photo `json:"photos"`
}

type Photo struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	File          []byte `json:"file"`
	Date          string `json:"date"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
	LikeStatus    bool   `json:"likeStatus"`
}

type Like struct {
	LikeId          uint64 `json:"likeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

type Comments struct {
	RequestIdentifier uint64    `json:"requestIdentifier"`
	PhotoIdentifier   uint64    `json:"photoIdentifier"`
	PhotoOwner        uint64    `json:"identifier"`
	Comments          []Comment `json:"comments"`
}

type Comment struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	PhotoId       uint64 `json:"photoId"`
	PhotoOwner    uint64 `json:"photoOwner"`
	OwnerUsername string `json:"ownerUsername"`
	Username      string `json:"username"`
	Content       string `json:"content"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CreateUser(User) (User, error)
	SetUsername(User, string) (User, error)
	GetUserId(string) (User, error)
	CheckUserById(User) (User, error)
	CheckUserByUsername(User) (User, error)
	CheckUser(User) (User, error)
	GetMyStream(User) ([]PhotoStream, error)

	SetFollow(Follow) (Follow, error)
	RemoveFollow(uint64, uint64, uint64) error
	GetFollowingId(user1 uint64, user2 uint64) (Follow, error)
	GetFollowers(User, uint64) (Follow, error)
	GetFollowersCount(uint64) (int, error)
	GetFollowingsCount(uint64) (int, error)
	GetFollowStatus(uint64, uint64) (bool, error)

	CreateBan(Ban) (Ban, error)
	RemoveBan(Ban) error
	GetBans(User, uint64) (Ban, error)
	GetBanById(Ban) (Ban, error)
	UpdateBanStatus(int, uint64, uint64) error
	GetBanStatus(uint64, uint64) (bool, error)
	CheckIfBanned(uint64, uint64) (bool, error)

	SetPhoto(Photo) (Photo, error)
	RemovePhoto(uint64) error
	GetPhotos(User, uint64) ([]Photo, error)
	GetPhotosCount(uint64) (int, error)
	CheckPhoto(Photo) (Photo, error)

	SetLike(Like) (Like, error)
	RemoveLike(Like) error
	RemoveLikes(uint64, uint64) error
	GetLike(uint64, uint64) (Like, error)
	GetLikeById(Like) (Like, error)
	GetLikesCount(photoid uint64) (int, error)

	SetComment(Comment) (Comment, error)
	RemoveComment(Comment) error
	RemoveComments(uint64, uint64) error
	GetComments(photoid uint64) ([]Comment, error)
	GetCommentById(Comment) (Comment, error)
	GetCommentsCount(uint64) (int, error)

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
		usersDatabase := `CREATE TABLE users (
			Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE
			);`
		photosDatabase := `CREATE TABLE photos (
			Id INTEGER NOT NULL PRIMARY KEY, 
			userId INTEGER NOT NULL,
			photo BLOB,
			date TEXT ,
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

// Ping checks the connection to the database.
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
