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
	Nickname string `json:"nickname"`
}
type PhotoId struct {
	IdPhoto int64 `json:"photo_id"`
}
type Nickname struct {
	Nickname string `json:"nickname"`
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
	Nickname  string `json:"nickname"`
	Comment   string `json:"comment"`
}
