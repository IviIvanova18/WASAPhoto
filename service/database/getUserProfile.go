package database

func (db *appdbimpl) GetUserProfile(user User) (User, error) {
	// var user User
	err := db.c.QueryRow(`SELECT idUser, username, photosCount FROM users WHERE users.username = ?`, user.Username).Scan(&user.IDUser, &user.Username, &user.PhotosCount)
	if err != nil {
		return user, err
	}

	//Get Followers
	followers, err := db.GetFollowersById(user.IDUser)
	if err != nil {
		return user, err
	}
	if followers == nil {
		user.Followers = []string{}
	} else {
		user.Followers = followers
	}

	//Get followings
	followings, err := db.GetFollowingsById(user.IDUser)
	if err != nil {
		return user, err
	}
	if followings == nil {
		user.Followings = []string{}
	} else {
		user.Followings = followings
	}

	//Get Photos
	photos, err := db.GetPhotosById(user.IDUser)
	if err != nil {
		return user, err
	}
	if photos == nil {
		user.Photos = []string{}
	} else {
		user.Photos = photos
	}

	return user, nil
}
