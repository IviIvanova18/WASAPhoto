package database

func (db *appdbimpl) GetCommentsOfImage(photoId uint64) ([]Comment, error) {
	var comments []Comment

	query := `
		SELECT comments.idUser, comments.commentText, comments.id 
		FROM comments
		WHERE comments.idPhoto = ?
	`

	rows, err := db.c.Query(query, photoId)
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.IDUser, &comment.CommentText, &comment.IDPhoto)
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