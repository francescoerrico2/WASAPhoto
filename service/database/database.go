package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrPhotoDoesntExist = errors.New("photo doesn't exist")
var ErrUserBanned = errors.New("user is banned")

const PhotosPerUserHome = 3

type AppDatabase interface {
	CreateUser(User) error

	ModifyUsername(User, Username) error

	SearchUser(searcher User, userToSearch User) ([]CompleteUser, error)

	CreatePhoto(Photo) (int64, error)

	LikePhoto(PhotoId, User) error

	UnlikePhoto(PhotoId, User) error

	CommentPhoto(PhotoId, User, Comment) (int64, error)

	UncommentPhoto(PhotoId, User, CommentId) error

	FollowUser(a User, b User) error

	UnfollowUser(a User, b User) error

	BanUser(a User, b User) error

	UnbanUser(a User, b User) error

	GetStream(User) ([]Photo, error)

	RemovePhoto(User, PhotoId) error

	GetFollowers(User) ([]User, error)

	GetFollowing(User) ([]User, error)

	GetPhotosList(a User, b User) ([]Photo, error)

	UncommentPhotoAuthor(PhotoId, CommentId) error

	GetUsername(User) (string, error)

	BannedUserCheck(a User, b User) (bool, error)

	CheckUser(a User) (bool, error)

	CheckPhotoExistence(p PhotoId) (bool, error)

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	_, errPramga := db.Exec(`PRAGMA foreign_keys= ON`)
	if errPramga != nil {
		return nil, fmt.Errorf("error setting pragmas: %w", errPramga)
	}

	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		err = createDatabase(db)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{c: db}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

func createDatabase(db *sql.DB) error {
	tables := [6]string{
		`CREATE TABLE IF NOT EXISTS users (
			id_user VARCHAR(16) NOT NULL PRIMARY KEY,
			username VARCHAR(16) NOT NULL
			);`,
		`CREATE TABLE IF NOT EXISTS photos (
			id_photo INTEGER PRIMARY KEY AUTOINCREMENT,
			id_user VARCHAR(16) NOT NULL,
			date DATETIME NOT NULL,
			FOREIGN KEY(id_user) REFERENCES users (id_user) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS  likes (
			id_photo INTEGER NOT NULL,
			id_user VARCHAR(16) NOT NULL,
			PRIMARY KEY (id_photo,id_user),
			FOREIGN KEY(id_photo) REFERENCES photos (id_photo) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS comments (
			id_comment INTEGER PRIMARY KEY AUTOINCREMENT,
			id_photo INTEGER NOT NULL,
			id_user VARCHAR(16) NOT NULL,
			comment VARCHAR(30) NOT NULL,
			FOREIGN KEY(id_photo) REFERENCES photos (id_photo) ON DELETE CASCADE,
			FOREIGN KEY(id_user) REFERENCES users (id_user) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS banned_users (
			banner VARCHAR(16) NOT NULL,
			banned VARCHAR(16) NOT NULL,
			PRIMARY KEY (banner,banned),
			FOREIGN KEY(banner) REFERENCES users (id_user) ON DELETE CASCADE,
			FOREIGN KEY(banned) REFERENCES users (id_user) ON DELETE CASCADE
			);`,
		`CREATE TABLE IF NOT EXISTS followers(
			follower VARCHAR(16) NOT NULL,
			followed VARCHAR(16) NOT NULL,
			PRIMARY KEY (follower,followed),
			FOREIGN KEY(follower) REFERENCES users (id_user) ON DELETE CASCADE,
			FOREIGN KEY(followed) REFERENCES users (id_user) ON DELETE CASCADE
			);`,
	}

	for i := 0; i < len(tables); i++ {

		sqlStmt := tables[i]
		_, err := db.Exec(sqlStmt)

		if err != nil {
			return err
		}
	}
	return nil
}
