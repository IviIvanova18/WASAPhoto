package database

func (db *appdbimpl) CommentPhoto(comment Comment) (Comment, error) {
	res, err := db.c.Exec(`INSERT INTO comments (id, idPhoto, idUser, commentText) VALUES (NULL, ?, ?, ?)`,
		comment.PhotoID, comment.UserID, comment.CommentText)

	if err != nil {
		return comment, err
	}

	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.IDComment = uint64(lastInsertId)
	return comment, nil
}
