package database

func (db *appdbimpl) GetCommentsOfImage(photoId uint64) ([]Comment, error) {
	var comments []Comment

	query := `
		SELECT users.username, comments.idUser, comments.commentText, comments.id, comments.idPhoto
		FROM comments
		INNER JOIN users
		ON comments.idUser = users.idUser
		WHERE comments.idPhoto = ?
	`

	rows, err := db.c.Query(query, photoId)
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.Username, &comment.IDUser, &comment.CommentText, &comment.IDComment,&comment.IDPhoto)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}