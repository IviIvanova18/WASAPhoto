package database

func (db *appdbimpl) ListUsers() ([]UserLogin, error) {
	var ret []UserLogin

	rows, err := db.c.Query(`SELECT idUser, username, photosCount FROM users`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var user UserLogin
		var photoCount uint64
		err = rows.Scan(&user.ID, &user.Username, &photoCount)
		if err != nil {
			return nil, err
		}

		ret = append(ret, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
