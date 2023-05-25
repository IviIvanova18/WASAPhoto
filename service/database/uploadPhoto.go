package database

func (db *appdbimpl) UploadPhoto(photo Photo) (Photo, error) {
	res, err := db.c.Exec(`INSERT INTO photos (idPhoto, idUser, comments, likes ,date, path) VALUES (NULL, ?, ?, ?, ?, ?)`,
		photo.IDUser, photo.Comments, photo.Likes, photo.DateTime, photo.Path)

	if err != nil {
		return photo, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return photo, err
	} else if affected == 0 {
		return photo, ErrUserDoesNotExist
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return photo, err
	}
	photo.IDPhoto = uint64(lastInsertId)
	return photo, nil

}
