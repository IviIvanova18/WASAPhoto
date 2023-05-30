package database

func (db *appdbimpl) GetPhoto(id uint64) (Photo, error) {
	var photo Photo
	err := db.c.QueryRow(`SELECT idPhoto, IdUser, comments, likes, path 
	FROM photos WHERE idPhoto=?`, id).Scan(&photo.PhotoID, &photo.UserID, &photo.Comments, &photo.Likes, &photo.Path)

	if err != nil {
		return photo, err
	}
	return photo, nil
}
