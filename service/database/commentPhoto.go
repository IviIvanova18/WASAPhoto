package database

func (db *appdbimpl) CommentPhoto(comment Comment) (Comment, error) {
	_, err := db.c.Exec(`INSERT INTO comments (id, idPhoto, idUser, commentText) VALUES (NULL, ?, ?, ?)`,
		comment.PhotoID, comment.UserID, comment.CommentText)

	if err != nil {
		return comment, err
	}
	return comment, nil
}
