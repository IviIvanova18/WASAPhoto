package database

func (db *appdbimpl) SearchUser(user User) (User, error) {
	err := db.c.QueryRow(`SELECT idUser, username, photosCount FROM users WHERE username=?`, user.Username).Scan(&user.IDUser, &user.Username, &user.PhotosCount)
	if err != nil {
		return user, err
	}

	followers, err := db.GetFollowersById(user.IDUser)
	if err != nil {
		return user, err
	}

	followings, err := db.GetFollowingsById(user.IDUser)
	if err != nil {
		return user, err
	}

	photos, err := db.GetPhotosById(user.IDUser)
	if err != nil {
		return user, err
	}
	if followers == nil {
		user.Followers = []string{}
	} else {
		user.Followers = followers
	}

	if followings == nil {
		user.Followings = []string{}
	} else {
		user.Followings = followings
	}

	if photos == nil {
		user.Photos = []string{}
	} else {
		user.Photos = photos
	}

	return user, nil
}
