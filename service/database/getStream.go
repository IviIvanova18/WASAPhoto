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
		WHERE followings.idFollower = ?
		ORDER BY photos.date DESC
	`

	rows, err := db.c.Query(query, user.ID)
	if err != nil {
		return []Photo{}, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.IDPhoto, &photo.IDUser, &photo.Likes, &photo.Comments, &photo.DateTime, &photo.Username, &photo.Path)
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
