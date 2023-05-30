package database

func (db *appdbimpl) GetStreamFollowing(user UserLogin) ([]Photo, error) {
	var photos []Photo

	query := `
		SELECT photos.idPhoto, photos.idUser, photos.likes, photos.comments, photos.date, users.username, photos.path
		FROM photos
		INNER JOIN followings
		ON photos.idUser = followings.idFollowed
		INNER JOIN users
		ON users.idUser = photos.idUser
		LEFT JOIN bannedUsers
		ON followings.idFollowed = bannedUsers.idUser AND followings.idFollower = bannedUsers.idBannedUser
		WHERE followings.idFollower = ? AND bannedUsers.idUser IS NULL
		ORDER BY photos.date DESC
	`

	rows, err := db.c.Query(query, user.ID)
	if err != nil {
		return []Photo{}, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.PhotoID, &photo.UserID, &photo.Likes, &photo.Comments, &photo.DateTime, &photo.Username, &photo.Path)
		if err != nil {
			return []Photo{}, err
		}

		photos = append(photos, photo)
	}

	if err = rows.Err(); err != nil {
		return []Photo{}, err
	}

	return photos, nil
}
