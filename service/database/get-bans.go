package database

func (db *appdbimpl) GetAllBannedUsersDB(uid uint64) ([]Ban, error) {
	var banned []Ban

	query := `
		SELECT users.username
		FROM users
		INNER JOIN bannedUsers
		ON bannedUsers.idBannedUser = users.idUser
		WHERE bannedUsers.idUser = ?
	`

	rows, err := db.c.Query(query, uid)
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var ban Ban
		err = rows.Scan(&ban.BannedUser)
		if err != nil {
			return nil, err
		}

		banned = append(banned, ban)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return banned, nil
}
