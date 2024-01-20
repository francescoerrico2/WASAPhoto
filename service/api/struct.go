package api

import (
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

// eliminare error
const INTERNAL_ERROR_MSG = "internal server error"
const PNG_ERROR_MSG = "file is not a png format"
const JPG_ERROR_MSG = "file is not a jpg format"
const IMG_FORMAT_ERROR_MSG = "images must be jpeg or png"
const INVALID_JSON_ERROR_MSG = "invalid json format"
const INVALID_IDENTIFIER_ERROR_MSG = "identifier must be a string between 3 and 16 characters"

type JSONErrorMsg struct {
	Message string `json:"message"`
}

type Photo struct {
	Comments []database.CompleteComment `json:"comments"`
	Likes    []database.CompleteUser    `json:"likes"`
	Owner    string                     `json:"owner"`
	PhotoId  int                        `json:"photo_id"`
	Date     time.Time                  `json:"date"`
}

type User struct {
	IdUser string `json:"user_id"`
}

type PhotoId struct {
	IdPhoto int64 `json:"photo_id"`
}

type Nickname struct {
	Nickname string `json:"Nickname"`
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

type Profile struct {
	Name      string           `json:"user_id"`
	Nickname  string           `json:"nickname"`
	Followers []database.User  `json:"followers"`
	Following []database.User  `json:"following"`
	Posts     []database.Photo `json:"posts"`
}

func (u User) ToDatabase() database.User {
	return database.User{IdUser: u.IdUser}
}

func (p Photo) ToDatabase() database.Photo {
	return database.Photo{
		Comments: p.Comments,
		Likes:    p.Likes,
		Owner:    p.Owner,
		PhotoId:  p.PhotoId,
		Date:     p.Date,
	}
}

func (p PhotoId) ToDatabase() database.PhotoId {
	return database.PhotoId{
		IdPhoto: p.IdPhoto,
	}
}

func (n Nickname) ToDatabase() database.Nickname {
	return database.Nickname{Nickname: n.Nickname}
}

func (c Comment) ToDatabase() database.Comment {
	return database.Comment{
		Comment: c.Comment,
	}
}

func (c CommentId) ToDatabase() database.CommentId {
	return database.CommentId{
		IdComment: c.IdComment,
	}
}

func (cc CompleteComment) ToDatabase() database.CompleteComment {
	return database.CompleteComment{
		IdComment: cc.IdComment,
		IdPhoto:   cc.IdPhoto,
		IdUser:    cc.IdUser,
		Comment:   cc.Comment,
	}
}
