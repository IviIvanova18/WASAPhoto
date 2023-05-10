package database

func (db *appdbimpl) GetAllBannedUsersDB(uid uint64) ([]string, error) {
	var banned []string

	query := `
		SELECT users.username
		FROM users
		INNER JOIN bannedUsers
		ON bannedUsers.idBannedUser = users.idUser
		WHERE bannedUsers.idUser = ?
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

		banned = append(banned, username)
	}

	if err = rows.Err(); err != nil {
		return []string{}, err
	}
	return banned, nil
}
