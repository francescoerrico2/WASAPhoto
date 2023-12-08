package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

type Profile struct {
	RequestId      uint64 ` json:"requestId"`
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

func (u *User) FromDB(user database.User) {
	u.Id = user.Id
	u.Username = user.Username
}

func (u *User) ToDB() database.User {
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

func (s *PhotoStream) PhotoStreamFromDB(photoStream database.PhotoStream) {
	s.Id = photoStream.Id
	s.UserId = photoStream.UserId
	s.File = photoStream.File
	s.Date = photoStream.Date
	s.LikeCount = photoStream.LikeCount
	s.CommentCount = photoStream.CommentCount
}

func (s *PhotoStream) PhotoStreamToDB() database.PhotoStream {
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

func (f *Follow) FollowFromDB(follow database.Follow) {
	f.FollowId = follow.FollowId
	f.FollowedId = follow.FollowedId
	f.UserId = follow.UserId
}

func (f *Follow) FollowToDB() database.Follow {
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

func (b *Ban) BanFromDB(ban database.Ban) {
	b.BanId = ban.BanId
	b.BannedId = ban.BannedId
	b.UserId = ban.UserId
}

func (b *Ban) BanToDB() database.Ban {
	return database.Ban{
		BanId:    b.BanId,
		BannedId: b.BannedId,
		UserId:   b.UserId,
	}
}

type Photo struct {
	Id           uint64 `json:"id"`
	UserId       uint64 `json:"userId"`
	File         []byte `json:"file"`
	Date         string `json:"date"`
	LikeCount    int    `json:"likeCount"`
	CommentCount int    `json:"commentCount"`
}

func (p *Photo) PhotoFromDB(photo database.Photo) {
	p.Id = photo.Id
	p.UserId = photo.UserId
	p.File = photo.File
	p.Date = photo.Date
	p.LikeCount = photo.LikesCount
	p.CommentCount = photo.CommentsCount
}

func (p *Photo) PhotoToDB() database.Photo {
	return database.Photo{
		Id:            p.Id,
		UserId:        p.UserId,
		File:          p.File,
		Date:          p.Date,
		LikesCount:    p.LikeCount,
		CommentsCount: p.CommentCount,
	}
}

type Like struct {
	LikeId          uint64 `json:"LikeId"`
	UserIdentifier  uint64 `json:"identifier"`
	PhotoIdentifier uint64 `json:"photoIdentifier"`
	PhotoOwner      uint64 `json:"photoOwner"`
}

func (l *Like) LikeFromDB(like database.Like) {
	l.LikeId = like.LikeId
	l.UserIdentifier = like.UserIdentifier
	l.PhotoIdentifier = like.PhotoIdentifier
	l.PhotoOwner = like.PhotoOwner

}

func (l *Like) LikeToDB() database.Like {
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

func (c *Comment) CommentFromDB(comment database.Comment) {
	c.Id = comment.Id
	c.UserId = comment.UserId
	c.PhotoId = comment.PhotoId
	c.Content = comment.Content
}

func (c *Comment) CommentToDB() database.Comment {
	return database.Comment{
		Id:         c.Id,
		UserId:     c.UserId,
		PhotoId:    c.PhotoId,
		PhotoOwner: c.PhotoOwner,
		Content:    c.Content,
	}
}
