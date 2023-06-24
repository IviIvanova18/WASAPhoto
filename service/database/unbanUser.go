package database

func (db *appdbimpl) UnbanUser(idUser uint64, idBannedUser uint64) error {
	res, err := db.c.Exec(`DELETE FROM bannedUsers WHERE idUser=? AND idBannedUser=?`,
		idUser, idBannedUser)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	} else if rows == 0 {
		return ErrBanDoesNotExist
	}
	return nil
}
