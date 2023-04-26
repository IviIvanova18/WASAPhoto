package database

func (db *appdbimpl) UnlikePhoto(idPhoto uint64, idUser uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE idUser=? AND idPhoto=?`, idUser, idPhoto)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrLikeNotFound
	}
	return err
}
