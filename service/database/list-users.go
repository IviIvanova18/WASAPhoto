package database

func (db *appdbimpl) ListUsers() ([]UserLogin, error) {
	var ret []UserLogin

	rows, err := db.c.Query(`SELECT idUser, username, photosCount FROM users`)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var f UserLogin
		var photoCount uint64
		err = rows.Scan(&f.ID, &f.Username, &photoCount)
		if err != nil {
			return nil, err
		}

		ret = append(ret, f)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
