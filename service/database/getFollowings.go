package database

func (db *appdbimpl) GetFollowingsDB(uid uint64) ([]string, error) {
	var followings []string

	query := `
		SELECT users.username
		FROM users
		INNER JOIN followings
		ON followings.idFollowed = users.idUser
		WHERE followings.idFollower = ?
	`

	rows, err := db.c.Query(query, uid)
	defer func() { _ = rows.Close() }()

	if err != nil {
		return []string{}, err
	}

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			return []string{}, err
		}

		followings = append(followings, username)
	}

	if err = rows.Err(); err != nil {
		return []string{}, err
	}
	return followings, nil
}
