package database

func (db *appdbimpl) SetMyUserName(u UserLogin) error {
	res, err := db.c.Exec(`UPDATE users SET username= ? WHERE idUser = ?`, u.Username, u.ID)
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

	
	