package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type Profile struct {
	RequestId      uint64 `json:"requestId"`
	Id             uint64 `json:"id"`
	Username       string `json:"username"`
	FollowersCount int    `json:"followersCount"`
	FollowingCount int    `json:"followingCount"`
	PhotoCount     int    `json:"photoCount"`
	FollowStatus   bool   `json:"followStatus"`
	BanStatus      bool   `json:"banStatus"`
	CheckIfBanned  bool   `json:"checkIfBanned"`
}

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
}

func (u *User) FromDatabase(user database.User) {
	u.Id = user.Id
	u.Username = user.Username
}

func (u *User) ToDatabase() database.User {
	return database.User{
		Id:       u.Id,
		Username: u.Username,
	}
}

type PhotoStream struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
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

type Follow struct {
	FollowId   uint64 `json:"followId"`
	FollowedId uint64 `json:"followedId"`
	UserId     uint64 `json:"userId"`
}

func (f *Follow) FollowFromDatabase(follow database.Follow) {
	f.FollowId = follow.FollowId
	f.FollowedId = follow.FollowedId
	f.UserId = follow.UserId
}

func (f *Follow) FollowToDatabase() database.Follow {
	return database.Follow{
		FollowId:   f.FollowId,
		FollowedId: f.FollowedId,
		UserId:     f.UserId,
	}
}

type Ban struct {
	BanId    uint64 `json:"banId"`
	BannedId uint64 `json:"bannedId"`
	UserId   uint64 `json:"userId"`
}

func (b *Ban) BanFromDatabase(ban database.Ban) {
	b.BanId = ban.BanId
	b.BannedId = ban.BannedId
	b.UserId = ban.UserId
}

func (b *Ban) BanToDatabase() database.Ban {
	return database.Ban{
		BanId:    b.BanId,
		BannedId: b.BannedId,
		UserId:   b.UserId,
	}
}

type Photo struct {
	Id            uint64 `json:"id"`
	UserId        uint64 `json:"userId"`
	File          []byte `json:"file"`
	Date          string `json:"date"`
	LikesCount    int    `json:"likesCount"`
	CommentsCount int    `json:"commentsCount"`
}

func (p *Photo) PhotoFromDatabase(photo database.Photo) {
	p.Id = photo.Id
	p.UserId = photo.UserId
	p.File = photo.File
	p.Date = photo.Date
	p.LikesCount = photo.LikesCount
	p.CommentsCount = photo.CommentsCount
}

func (p *Photo) PhotoToDatabase() database.Photo {
	return database.Photo{
		Id:            p.Id,
		UserId:        p.UserId,
		File:          p.File,
		Date:          p.Date,
		LikesCount:    p.LikesCount,
		CommentsCount: p.CommentsCount,
	}
}

type Like struct {
	LikeId          uint64 `json:"likeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

func (l *Like) LikeFromDatabase(like database.Like) {
	l.LikeId = like.LikeId
	l.UserIdentifier = like.UserIdentifier
	l.PhotoIdentifier = like.PhotoIdentifier
	l.PhotoOwner = like.PhotoOwner

}

func (l *Like) LikeToDatabase() database.Like {
	return database.Like{
		LikeId:          l.LikeId,
		UserIdentifier:  l.UserIdentifier,
		PhotoIdentifier: l.PhotoIdentifier,
		PhotoOwner:      l.PhotoOwner,
	}
}

type Comment struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"userId"`
	PhotoId    uint64 `json:"photoId"`
	PhotoOwner uint64 `json:"photoOwner"`
	Content    string `json:"content"`
}

func (c *Comment) CommentFromDatabase(comment database.Comment) {
	c.Id = comment.Id
	c.UserId = comment.UserId
	c.PhotoId = comment.PhotoId
	c.Content = comment.Content
}

func (c *Comment) CommentToDatabase() database.Comment {
	return database.Comment{
		Id:         c.Id,
		UserId:     c.UserId,
		PhotoId:    c.PhotoId,
		PhotoOwner: c.PhotoOwner,
		Content:    c.Content,
	}
}
