package database

func (db *appdbimpl) DeletePhoto(id uint64) error {
	res, err := db.c.Exec(`DELETE FROM photos WHERE idPhoto= ? `, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrPhotoDoesNotExists
	}
	return nil
}
