package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// * Struct per commento
type Comments struct {
	Username   string    `json:"username"`
	Identifier int       `json:"identifier"`
	Comments   []Comment `json:"comments"`
}

type Comment struct {
	Content string `json:"comment"`
}

// * Struct per foto
type Photos struct {
	Username   string  `json:"username"`
	Identifier int     `json:"identifier"`
	Photos     []Photo `json:"photos"`
}
type Photo struct {
	Username        string `json:"username"`
	PhotoIdentifier int64  `json:"photoIdentifier"`
	//file            string `json:"file"`
}

// * Struct per bans
type Bans struct {
	Username   string `json:"username"`
	Identifier int    `json:"identifier"`
	Bans       []Ban  `json:"bans"`
}
type Ban struct {
	Identifier    int    `json:"identifier"`
	Username      string `json:"username"`
	BanIdentifier int    `json:"banIdentifier"`
}

// * Struct per followers
type Followers struct {
	Username   string   `json:"username"`
	Identifier int      `json:"identifier"`
	Followers  []Follow `json:"Followers"`
}
type Follow struct {
	Identifier       int    `json:"identifier"`
	Username         string `json:"username"`
	FollowIdentifier int    `json:"followIdentifier"`
}

type User struct {
	Identifier int    `json:"id"`
	Username   string `json:"username"`
}

func (u *User) FromDatabase(user database.User) {
	u.Identifier = user.Identifier
	u.Username = user.Username
}
func (u *User) ToDatabase() database.User {
	return database.User{
		Identifier: u.Identifier,
		Username:   u.Username,
	}
}
