package database

func (db *appdbimpl) SearchUser(user User) (User, error) {
	err := db.c.QueryRow(`SELECT id_user, username, photosCount FROM users WHERE username=?`, user.Username).Scan(&user.IDUser, &user.Username, &user.PhotosCount)
	if err != nil {
		return user, err
	}

	// followers, err := db.GetFollowersIDs(user.IDUser)
	// if err != nil {
	// 	return user, err
	// }

	// followings, err := db.GetFollowingsIDs(user.IDUser)
	// if err != nil {
	// 	return user, err
	// }

	// photoes, err := db.GetPhotoesIDs(user.IDUser)
	// if err != nil {
	// 	return user, err
	// }
	// if followers == nil {
	// 	user.Followers = []string{}
	// } else {
	// 	user.Followers = followers
	// }

	// if followings == nil {
	// 	user.Followings = []string{}
	// } else {
	// 	user.Followings = followings
	// }

	// if photoes == nil {
	// 	user.Photoes = []uint64{}
	// } else {
	// 	user.Photoes = photoes
	// }

	return user, nil
}
