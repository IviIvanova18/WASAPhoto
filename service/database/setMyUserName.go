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
		return ErrUserDoesNotExist
	}
	return nil
	
}

	
	