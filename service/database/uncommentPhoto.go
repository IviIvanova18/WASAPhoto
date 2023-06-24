package database

func (db *appdbimpl) UncommentPhoto(id uint64, idPhoto uint64) error {

	res, err := db.c.Exec(`DELETE FROM comments WHERE id=? AND idPhoto=?`, id, idPhoto)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrCommentNotFound
	}
	return err
}
