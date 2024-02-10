package database

import "time"

type Photo struct {
	Comments []CompleteComment `json:"comments"`
	Likes    []CompleteUser    `json:"likes"`
	Owner    string            `json:"owner"`
	PhotoId  int               `json:"photo_id"`
	Date     time.Time         `json:"date"`
}

type User struct {
	IdUser string `json:"user_id"`
}

type CompleteUser struct {
	IdUser   string `json:"user_id"`
	Username string `json:"Username"`
}
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"`
}
type Username struct {
	Username string `json:"username"`
}

type Comment struct {
	Comment string `json:"comment"`
}

type CommentId struct {
	IdComment int64 `json:"comment_id"`
}

type CompleteComment struct {
	IdComment int64  `json:"comment_id"`
	IdPhoto   int64  `json:"photo_id"`
	IdUser    string `json:"user_id"`
	Username  string `json:"username"`
	Comment   string `json:"comment"`
}
