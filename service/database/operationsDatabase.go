package database

// import (
// 	"database/sql"
// 	"errors"
// )

func (db *appdbimpl) GetIDByUsername(username string) (uint64, error) {
	var id uint64
	err := db.c.QueryRow(`SELECT id_user FROM users WHERE username=?`, username).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *appdbimpl) GetUsernameById(id uint64) (string, error) {
	var username string
	err := db.c.QueryRow(`SELECT username FROM users WHERE idUser=?`, id).Scan(&username)

	return username, err
}
