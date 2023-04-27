package database

func (db *appdbimpl) UnfollowUser(idFollowed uint64, idFollower uint64) error {
	res, err := db.c.Exec(`DELETE FROM followings WHERE idFollowed=? AND idFollower=?`,
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