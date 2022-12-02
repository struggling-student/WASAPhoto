package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// Struct for user
//
// This structure is used only for the api and not for the database
type Users struct {
	Users []User `json:"users"`
}
type User struct {
	// Identifier is the unique identifier for the user
	Id uint64 `json:"id"`
	// Username is the username of the user
	Username string `json:"username"`
}

// FromDatabase converts a database.User to an api.User
func (u *User) FromDatabase(user database.User) {
	u.Id = user.Id
	u.Username = user.Username
}

// ToDatabase converts an api.User to a database.User
func (u *User) ToDatabase() database.User {
	return database.User{
		Id:       u.Id,
		Username: u.Username,
	}
}

// Struct per followers
//
//	This structure is used only for the api and not for the database
type Followers struct {
	// Identifier for the user that has the followers
	Id int `json:"identifier"`
	// List of followers
	Followers []Follow `json:"Followers"`
}
type Follow struct {
	// Identifier for the user who is followed
	Identifier int `json:"identifier"`
	// FollowIdentifier is the identifier for the follow action
	FollowIdentifier int `json:"followIdentifier"`
}

// Struct per bans
//
// This structure is used only for the api and not for the database
type Bans struct {
	// Identifier for the user that has the bans
	Identifier int `json:"identifier"`
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

// FromDatabase converts a database.User to an api.User
func (b *Ban) BanFromDatabase(ban database.Ban) {
	b.BanId = ban.BanId
	b.BannedId = ban.BannedId
	b.UserId = ban.UserId
}

// ToDatabase converts an api.User to a database.User
func (b *Ban) BanToDatabase() database.Ban {
	return database.Ban{
		BanId:    b.BanId,
		BannedId: b.BannedId,
		UserId:   b.UserId,
	}
}

// Struct for photos
//
// This structure is used only for the api and not for the database
type Photos struct {
	// Identifier of the user who has the photos
	Identifier int `json:"identifier"`
	// List of photos
	Photos []Photo `json:"photos"`
}
type Photo struct {
	Id     uint64    `json:"id"`
	UserId uint64    `json:"userId"`
	File   [255]byte `json:"file"`
	Date   string    `json:"date"`
}

// Struct for likes
//
// This structure is used only for the api and not for the database
type Likes struct {
	// Identifier of the user who has commented
	Id int `json:"id"`
	// List of likes under a photo
	Likes []Like `json:"likes"`
}
type Like struct {
	// Identifier for the user
	Id int `json:"id"`
	// Identifier for the photo that has the likes
	PhotoIdentifier int `json:"photoIdentifier"`
	// Identifier for the like that has been added
	LikeId int `json:"likeId"`
	// Identifier for the user who liked the photo
	Identifier int `json:"identifier"`
}

// Struct for comments
//
// This structure is used only for the api and not for the database
type Comments struct {
	// Identifier of the user who has commented
	Id int `json:"id"`
	// Identifier for the photo that has the comments
	PhotoIdentifier int `json:"photoIdentifier"`
	// List of comments under the photo
	Comments []Comment `json:"comments"`
}
type Comment struct {
	// Identifier of the user who has commented
	Id int `json:"id"`
	// Identifier for the photo that has the comments
	PhotoId int `json:"photoId"`
	// Identifier of the user who has commented
	CommentId int `json:"commentId"`
	// Content of the comment
	Content string `json:"comment"`
}
