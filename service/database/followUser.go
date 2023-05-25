package database

func (db *appdbimpl) FollowUser(idFollower uint64, idFollowed uint64) error {
	res, err := db.c.Exec(`INSERT INTO followings (id, idFollower, idFollowed) VALUES (NULL, ?, ?)`,
		idFollower, idFollowed)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
