package database

func (db *appdbimpl) GetAllLikesOfPhoto(photoId uint64) ([]Like, error) {
	var likes []Like

	query := `
		SELECT likes.id, likes.idUser, likes.idPhoto
		FROM likes
		WHERE likes.idPhoto = ?
	`

	rows, err := db.c.Query(query, photoId)
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var like Like
		err = rows.Scan(&like.ID, &like.IDUser, &like.IDPhoto)
		if err != nil {
			return nil, err
		}

		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return likes, nil
}
