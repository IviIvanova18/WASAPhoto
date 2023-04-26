package database

func (db *appdbimpl) CommentPhoto(comment Comment) error {
	res, err := db.c.Exec(`INSERT INTO comments (id, idPhoto, idUser, commentText) VALUES (NULL, ?, ?, ?)`,
		comment.IDPhoto, comment.IDUser, comment.CommentText)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if rows == 0 {
		return ErrCommentNotFound
	}
	return err
}
