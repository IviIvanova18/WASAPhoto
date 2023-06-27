package database

import "database/sql"

func (db *appdbimpl) GetIDByPhotoID(id uint64) (uint64, error) {
	var idUser uint64
	err := db.c.QueryRow(`SELECT idUser FROM photos WHERE idPhoto=?`, id).Scan(&idUser)

	return idUser, err
}

func (db *appdbimpl) DeleteLikes(idPhoto uint64) error {

	res, err := db.c.Exec(`DELETE FROM likes WHERE idPhoto=?`, idPhoto)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrUserDoesNotExist
	}
	return err
}

func (db *appdbimpl) DeleteComments(idPhoto uint64) error {

	res, err := db.c.Exec(`DELETE FROM comments WHERE idPhoto=?`, idPhoto)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrUserDoesNotExist
	}
	return err
}

func (db *appdbimpl) UpdatePhotoCountUser(idUser uint64, count int64) error {

	query := `UPDATE users SET photosCount = photosCount + ? WHERE idUser = ?`

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

	query := `UPDATE photos SET likes = likes + ? WHERE idPhoto = ?`

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
func (db *appdbimpl) IsBanned(user uint64, banned uint64) (bool, error) {

	var id sql.NullInt64

	err := db.c.QueryRow(`SELECT id FROM bannedUsers WHERE idUser=? AND idBannedUser=?`, user, banned).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return id.Valid && id.Int64 != 0, nil

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

func (db *appdbimpl) GetFollowersById(id uint64) ([]string, error) {
	var followers []string
	query := `
	SELECT users.username FROM users
	INNER JOIN followings
	ON users.idUser = followings.idFollower
	WHERE followings.idFollowed = ? ;`

	rows, err := db.c.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)

		if err != nil {
			return nil, err
		}
		followers = append(followers, username)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}

func (db *appdbimpl) GetFollowingsById(id uint64) ([]string, error) {
	var followings []string
	query := `
	SELECT users.username 
	FROM users
    INNER JOIN followings
    ON users.idUser = followings.idFollowed
    WHERE followings.idFollower = ?; `
	rows, err := db.c.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)

		if err != nil {
			return nil, err
		}
		followings = append(followings, username)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return followings, nil
}

func (db *appdbimpl) GetPhotosById(id uint64) ([]uint64, []string, error) {
	var paths []string
	var ids []uint64

	rows, err := db.c.Query(`SELECT idPhoto, path FROM photos WHERE idUser=? ORDER BY photos.date DESC`, id)
	if err != nil {
		return nil, nil, err
	}

	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var idPhoto uint64
		var path string

		err = rows.Scan(&idPhoto, &path)

		if err != nil {
			return nil, nil, err
		}
		ids = append(ids, idPhoto)
		paths = append(paths, path)

	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	return ids, paths, nil
}
