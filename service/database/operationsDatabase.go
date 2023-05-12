package database

// import (
// 	"database/sql"
// 	"errors"
// )

func (db *appdbimpl) GetIDByUsername(username string) (uint64, error) {
	var id uint64
	err := db.c.QueryRow(`SELECT idUser FROM users WHERE username=?`, username).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db *appdbimpl) GetIDByPhotoID(id uint64) (uint64, error) {
	var idUser uint64
	err := db.c.QueryRow(`SELECT idUser FROM photos WHERE idPhoto=?`, id).Scan(&idUser)

	return idUser, err
}

func (db *appdbimpl) GetUsernameById(id uint64) (string, error) {
	var username string
	err := db.c.QueryRow(`SELECT username FROM users WHERE idUser=?`, id).Scan(&username)

	return username, err
}

func (db *appdbimpl) DeleteLikes(idPhoto uint64) error {
	_, err := db.c.Exec(`DELETE FROM likes WHERE idPhoto=?`, idPhoto)
	return err
}

func (db *appdbimpl) DeleteComments(idPhoto uint64) error {
	_, err := db.c.Exec(`DELETE FROM comments WHERE idPhoto=?`, idPhoto)
	return err
}


func (db *appdbimpl) UpdatePhotoCountUser(idUser uint64, count int64) error {
	query := `UPDATE users SET photosCount = photosCount + ?
			WHERE idUser = ?`

	res, err := db.c.Exec(query, count, idUser)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrUserDoesNotExist
	}
	return err
}

func (db *appdbimpl) UpdateLikesPhoto(photoId uint64, count int64) error {
	query := `UPDATE photos SET likes = likes + ?
				WHERE idPhoto = ?`
	res, err := db.c.Exec(query, count, photoId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrPhotoDoesNotExists
	}
	return err

}
func (db *appdbimpl) IsBanned(user uint64, banned uint64) error {
	var id uint64
	err := db.c.QueryRow(`SELECT id FROM bannedUsers WHERE idUser=? AND idBannedUser=?`, user, banned).Scan(&id)
	return err
}

func (db *appdbimpl) UpdateCommentsPhoto(photoId uint64, count int64) error {
	res, err := db.c.Exec(`UPDATE photos SET comments = comments + ? WHERE idPhoto = ?`, count, photoId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrPhotoDoesNotExists
	}
	return err

}


