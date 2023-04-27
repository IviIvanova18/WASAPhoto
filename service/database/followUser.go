package database

func (db *appdbimpl) FollowUser(idFollowed uint64, idFollower uint64) error {
	res, err := db.c.Exec(`INSERT INTO followings (idFollowed, idFollower) VALUES (?, ?)`,
	idFollowed, idFollower)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}