package database

func (db *appdbimpl) GetUserProfile(username string) (User, error) {
	var user User
	err := db.c.QueryRow(`SELECT idUser, username FROM users WHERE users.username = ?`, username).Scan(&user.UserID, &user.Username)
	if err != nil {
		return user, err
	}

	// Get Followers
	followers, err := db.GetFollowersById(user.UserID)
	if err != nil {
		return user, err
	}
	if followers == nil {
		user.Followers = []string{}
	} else {
		user.Followers = followers
	}

	// Get followings
	followings, err := db.GetFollowingsById(user.UserID)
	if err != nil {
		return user, err
	}
	if followings == nil {
		user.Followings = []string{}
	} else {
		user.Followings = followings
	}

	// Get Photos
	photoIds, paths, err := db.GetPhotosById(user.UserID)
	if err != nil {
		return user, err
	}
	if photoIds == nil {
		user.PhotosId = []uint64{}
	} else if paths == nil {
		user.PhotosPath = []string{}
	} else {
		user.PhotosId = photoIds
		user.PhotosPath = paths
	}
	return user, nil
}
