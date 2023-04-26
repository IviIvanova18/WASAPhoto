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

var ErrorUserDoesNotExist = errors.New("Error User does not exists!")
// var ErrImageDoesNotExist = errors.New("image does not exist")
// var ErrCommentNotFound = errors.New("comment does not exist")
// var ErrLikeNotFound = errors.New("like does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	CreateUser(username string) (UserLogin, error)
	GetIDByUsername(username string) (uint64, error)
	SetMyUserName(u UserLogin) error
	SearchUser(user User) (User, error)
	UploadPhoto(photo Photo) (Photo, error)
	GetUsernameById(id uint64) (string, error)

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
			username TEXT NOT NULL UNIQUE,
			comments INTEGER NOT NULL,
			likes INTEGER NOT NULL,
			date DATETIME NOT NULL,
			path VARBINARY NOT NULL,
			FOREIGN KEY (idUser)
				REFERENCES users (idUser)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos: %w", err)
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
	ID  uint64
	Username string
}

type User struct {
	IDUser     uint64
	Username   string
	PhotosCount  uint64
	FollowersCount int 	
	FollowingsCount int 
	Photos []uint64 	
	Followers  []string
	Followings []string
	
}

type Photo struct{
	IDPhoto uint64 	
	IDUser uint64		
	Username string		
	DateTime time.Time 	
	Likes int 			
	Comments []string 	
	Path []byte			
}

type Comment struct {
	IDComment uint64
	IDUser    uint64
	IDPhoto   uint64
	Comment     string
}

