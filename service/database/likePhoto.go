package database

func (db *appdbimpl) LikePhoto(idPhoto uint64, idUser uint64) error {
	_, err := db.c.Exec(`
		INSERT OR IGNORE INTO likes (idPhoto, idUser)
		VALUES (?, ?)`,
		idPhoto, idUser)
	if err != nil {
		return err
	}

	return err
}
