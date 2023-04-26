package database

func (db *appdbimpl) LikePhoto(idPhoto uint64, idUser uint64) error {
	_, err := db.c.Exec(`INSERT INTO likes (id, idPhoto, idUser) VALUES (NULL, ?, ?)`,
		idPhoto, idUser)
	return err
}
