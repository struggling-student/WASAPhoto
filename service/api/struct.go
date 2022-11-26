package api

type User struct {
	Username   string `json:"username"`
	Identifier int    `json:"identifier"`
}

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
