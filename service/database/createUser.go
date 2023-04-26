package database

func (db *appdbimpl) CreateUser(username string) (UserLogin, error) {
	res, err := db.c.Exec(`INSERT INTO users (id_user, username, photosCount) VALUES (NULL, ?, ?)`,
		username, 0)
	if err != nil {
		return UserLogin{
			ID:  0,
			Username: "",
		}, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return UserLogin{
			ID:  0,
			Username: "",
		}, err
	}

	return UserLogin{
		ID:  uint64(lastID),
		Username: username,
	}, nil
}
