package database

func (db *appdbimpl) SetMyUserName(u User) error {
	res, err := db.c.Exec(`UPDATE user SET username=?`,
		u.Username)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrUserDoesNotExist
	}
	return nil
}