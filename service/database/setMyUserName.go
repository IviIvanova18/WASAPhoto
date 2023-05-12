package database

func (db *appdbimpl) SetMyUserName(u UserLogin) error {
	_, err := db.c.Exec(`UPDATE users SET username= ? WHERE idUser = ?`, u.Username, u.ID)

	return err
}