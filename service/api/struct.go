package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// Struct for user
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

type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         string `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

func (s *PhotoStream) PhotoStreamFromDatabase(photoStream database.PhotoStream) {
	s.Id = photoStream.Id
	s.UserId = photoStream.UserId
	s.File = photoStream.File
	s.Date = photoStream.Date
	s.LikeCount = photoStream.LikeCount
	s.CommentCount = photoStream.CommentCount
}

func (s *PhotoStream) PhotoStreamToDatabase() database.PhotoStream {
	return database.PhotoStream{
		Id:           s.Id,
		UserId:       s.UserId,
		File:         s.File,
		Date:         s.Date,
		LikeCount:    s.LikeCount,
		CommentCount: s.CommentCount,
	}

}

// Struct for followers
// This structure is used only for the api and not for the database

type Follow struct {
	// BanIdentifier is the identifier for the ban action
	FollowId uint64 `json:"followId"`
	// Identifier for the user who is banned
	FollowedId uint64 `json:"followedId"`
	// Identifier for the user who is banning
	UserId uint64 `json:"userId"`
}

// FollowFromDatabase converts a database.Follow to an api.Follow
func (f *Follow) FollowFromDatabase(follow database.Follow) {
	f.FollowId = follow.FollowId
	f.FollowedId = follow.FollowedId
	f.UserId = follow.UserId
}

// FollowToDatabase converts an api.Follow to a database.Follow
func (f *Follow) FollowToDatabase() database.Follow {
	return database.Follow{
		FollowId:   f.FollowId,
		FollowedId: f.FollowedId,
		UserId:     f.UserId,
	}
}

// Struct for bans
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

// BanFromDatabase converts a database.Ban to an api.Ban
func (b *Ban) BanFromDatabase(ban database.Ban) {
	b.BanId = ban.BanId
	b.BannedId = ban.BannedId
	b.UserId = ban.UserId
}

// BanToDatabase converts an api.Ban to a database.Ban
func (b *Ban) BanToDatabase() database.Ban {
	return database.Ban{
		BanId:    b.BanId,
		BannedId: b.BannedId,
		UserId:   b.UserId,
	}
}

// Struct for photos
// This structure is used only for the api and not for the database
type Photos struct {
	// Identifier of the user who has the photos
	Identifier int `json:"identifier"`
	// List of photos
	Photos []Photo `json:"photos"`
}
type Photo struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"userId"`
	File   string `json:"file"`
	Date   string `json:"date"`
}

// PhotoFromDatabase converts a database.Photo to an api.Photo
func (p *Photo) PhotoFromDatabase(photo database.Photo) {
	p.Id = photo.Id
	p.UserId = photo.UserId
	p.File = photo.File
	p.Date = photo.Date
}

// PhotoToDatabase converts an api.Photo to a database.Photo
func (p *Photo) PhotoToDatabase() database.Photo {
	return database.Photo{
		Id:     p.Id,
		UserId: p.UserId,
		File:   p.File,
		Date:   p.Date,
	}
}

// Struct for likes
type Like struct {
	// Identifier for the like that has been added
	LikeId uint64 `json:"likeId"`
	// Identifier for the photo that has the likes
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	// Identifier for the user who liked the photo
	UserIdentifier uint64 `json:"identifier"`
}

// LikeFromDatabase converts a database.Like to an api.Like
func (l *Like) LikeFromDatabase(like database.Like) {
	l.LikeId = like.LikeId
	l.PhotoIdentifier = like.PhotoIdentifier
	l.UserIdentifier = like.UserIdentifier
}

// LikeToDatabase converts an api.Like to a database.Like
func (l *Like) LikeToDatabase() database.Like {
	return database.Like{
		LikeId:          l.LikeId,
		PhotoIdentifier: l.PhotoIdentifier,
		UserIdentifier:  l.UserIdentifier,
	}
}

// Struct for comments
// This structure is used only for the api and not for the database

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

// CommentFromDatabase converts a database.Comment to an api.Comment
func (c *Comment) CommentFromDatabase(comment database.Comment) {
	c.Id = comment.Id
	c.PhotoId = comment.PhotoId
	c.UserId = comment.UserId
	c.Content = comment.Content
}

// CommentToDatabase converts an api.Comment to a database.Comment
func (c *Comment) CommentToDatabase() database.Comment {
	return database.Comment{
		Id:      c.Id,
		PhotoId: c.PhotoId,
		UserId:  c.UserId,
		Content: c.Content,
	}
}
