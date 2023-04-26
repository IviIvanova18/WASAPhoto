package database

func (db *appdbimpl) GetPhoto(id uint64) ([]byte, error) {
	var path []byte
	err := db.c.QueryRow(`SELECT path FROM photos WHERE idPhoto=?`, id).Scan(&path)

	if err != nil {
		return make([]byte, 0), err
	}
	return path, nil
}
