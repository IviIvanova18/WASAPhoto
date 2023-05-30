package database

import (
	"database/sql"
)

func (db *appdbimpl) CreateUser(userLogin UserLogin) (UserLogin, error) {
	res, err := db.c.Exec(`INSERT INTO users (idUser, username, photosCount) VALUES (NULL, ?, ?)`,
		userLogin.Username, 0)
	if err != nil {
		var user UserLogin
		err := db.c.QueryRow(`SELECT idUser, username FROM users WHERE username = ?`, userLogin.Username).Scan(&user.ID, &user.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				return user, ErrUserDoesNotExist
			}
		}
		return user, nil
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return userLogin, err
	}
	userLogin.ID = uint64(lastInsertID)
	return userLogin, nil
}
