package database

func (db *appdbimpl) BanUser(idUser uint64, idBannedUser uint64) error {
	res, err := db.c.Exec(`INSERT INTO bannedUsers (idUser, idBannedUser) VALUES (?, ?)`,
		idUser, idBannedUser)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}