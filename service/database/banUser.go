package database

func (db *appdbimpl) BanUser(idUser uint64, idBannedUser uint64) error {
	_, err := db.c.Exec(`INSERT INTO bannedUsers (id, idUser, idBannedUser) VALUES (NULL, ?, ?)`,
		idUser, idBannedUser)
	if err != nil {
		return err
	}

	return nil
}
