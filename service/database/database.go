/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var ErrUserDoesNotExist = errors.New("Error User does not exist!")
var ErrPhotoDoesNotExists = errors.New("Error photo does not exist!")
var ErrCommentNotFound = errors.New("Error comment does not exist!")
var ErrLikeNotFound = errors.New("Error like does not exist!")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CreateUser(userLogin UserLogin) (UserLogin, error) //.
	SetMyUserName(user UserLogin) error                //.

	GetIDByPhotoID(id uint64) (uint64, error)     //.
	ListUsers() ([]UserLogin, error)              //.
	GetUserProfile(username string) (User, error) //.

	GetFollowersById(id uint64) ([]string, error)        //.
	GetFollowingsById(id uint64) ([]string, error)       //.
	GetPhotosById(id uint64) ([]uint64, []string, error) //.

	BanUser(idUser uint64, idBannedUser uint64) error   //.
	UnbanUser(idUser uint64, idBannedUser uint64) error //.
	GetAllBannedUsersDB(uid uint64) ([]Ban, error)

	FollowUser(idFollowed uint64, idFollower uint64) error   //.
	UnfollowUser(idFollowed uint64, idFollower uint64) error //.

	UploadPhoto(photo Photo) (Photo, error) //.
	DeletePhoto(id uint64) error            //.
	GetPhoto(id uint64) (Photo, error)
	DeleteLikes(idPhoto uint64) error                      //.
	DeleteComments(idPhoto uint64) error                   //.
	UpdateLikesPhoto(photoId uint64, count int64) error    //.
	UpdatePhotoCountUser(idUser uint64, count int64) error //.

	CommentPhoto(comment Comment) (Comment, error)        //.
	UncommentPhoto(id uint64, idPhoto uint64) error       //.
	GetCommentsOfImage(photoId uint64) ([]Comment, error) //.
	IsBanned(user uint64, banned uint64) error            //.

	LikePhoto(idImage uint64, idUser uint64) error     //.
	UnlikePhoto(idImage uint64, idUser uint64) error   //.
	GetAllLikesOfPhoto(photoId uint64) ([]Like, error) //.

	UpdateCommentsPhoto(photoId uint64, count int64) error //.

	GetStreamFollowing(user UserLogin) ([]Photo, error) //.

	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)

	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
			idUser INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			photosCount INTEGER NOT NULL
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
			idPhoto INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idUser INTEGER NOT NULL,
			comments INTEGER NOT NULL,
			likes INTEGER NOT NULL,
			date DATETIME NOT NULL,
			path TEXT NOT NULL,
			FOREIGN KEY (idUser)
				REFERENCES users (idUser)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos: %w", err)
		}
	}
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idPhoto INTEGER NOT NULL,
			idUser INTEGER NOT NULL,
			FOREIGN KEY (idPhoto) REFERENCES photos(idPhoto),
			FOREIGN KEY (idUser) REFERENCES users(idUser),
			CONSTRAINT unique_user_photo UNIQUE (idUser, idPhoto)
		  );`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure likes: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idPhoto INTEGER NOT NULL,
			idUser INTEGER NOT NULL,
			commentText TEXT NOT NULL,
			FOREIGN KEY (idPhoto) REFERENCES photos(idPhoto),
			FOREIGN KEY (idUser) REFERENCES users(idUser)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure comments: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bannedUsers';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE bannedUsers (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idUser INTEGER NOT NULL,
			idBannedUser INTEGER NOT NULL,
			UNIQUE (idUser, idBannedUser),
			FOREIGN KEY (idUser) REFERENCES users (idUser),
			FOREIGN KEY (idBannedUser) REFERENCES users (idUSer)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure bannedUsers: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followings';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE followings (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			idFollower INTEGER NOT NULL,
			idFollowed INTEGER NOT NULL,
			UNIQUE (idFollower, idFollowed),
			FOREIGN KEY (idFollower) REFERENCES users (idUser),
			FOREIGN KEY (idFollowed) REFERENCES users (idUSer)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure followings: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}

type UserLogin struct {
	ID       uint64
	Username string
}

type User struct {
	UserID          uint64
	Username        string
	PhotosCount     uint64
	FollowersCount  uint64
	FollowingsCount uint64
	PhotosId        []uint64
	PhotosPath      []string
	Followers       []string
	Followings      []string
}

type Photo struct {
	PhotoID  uint64
	UserID   uint64
	Username string
	DateTime time.Time
	Likes    uint64
	Comments uint64
	Path     string
}

type Comment struct {
	IDComment   uint64
	UserID      uint64
	Username    string
	PhotoID     uint64
	CommentText string
}

type Like struct {
	ID      uint64
	UserID  uint64
	PhotoID uint64
}

type Ban struct {
	BannedUser string
}
