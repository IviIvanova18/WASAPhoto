package api

import (
	"git.wasaphoto.ivi/wasaphoto/service/database"
	"time"
	
)

type Identifier struct {
	Id uint64 `json:"id"`
}

// Error Message
type JSONErrorMsg struct{
	Message string
}

//Photo
type Photo struct{
	IDPhoto uint64 		`json:"id"`
	IDUser uint64		`json:"idUser"`
	Username string		`json:"username"`
	DateTime time.Time 	`json:"DateTime"`
	Likes int 			`json:"likes"`
	Comments uint64 	`json:"comments"`
	Path string			`json:"path"`
}




func (p *Photo) FromDatabase(photo database.Photo) {
	p.IDPhoto = photo.IDPhoto
	p.IDUser = photo.IDUser
	p.Username = photo.Username
	p.DateTime = photo.DateTime
	p.Likes = photo.Likes
	p.Comments = photo.Comments
	p.Path = photo.Path
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		IDPhoto:	p.IDPhoto,
		IDUser: 	p.IDUser,
		Username:	p.Username,
		DateTime:  	p.DateTime,
		Likes: 		p.Likes,
		Comments:   p.Comments,
		Path:		p.Path,
	}
}

// User 
type User struct{
	IDUser uint64 		`json:"id"`
	Username string 	`json:"username"`
	PhotosCount int 	`json:"photosCount"`
	FollowersCount int 	`json:"followersCount"`
	FollowingsCount int `json:"followingCount"`
	Photos []string 	`json:"photos"`
	Followers []string	`json:"followers"`
	Followings []string `json:"followings"`
}

func (user *User) ToDatabase() database.User {
	return database.User{
		IDUser:  user.IDUser,
		Username: user.Username,
		PhotosCount: user.PhotosCount,
		FollowersCount: user.FollowersCount,
		FollowingsCount: user.FollowingsCount,
		Followers: user.Followers,
		Followings: user.Followings,
		Photos: user.Photos,
	}
}
func (u *User) FromDatabase(user database.User){
	u.IDUser = user.IDUser
	u.Username = user.Username
	u.PhotosCount = user.PhotosCount
	u.Followers = user.Followers
	u.Followings = user.Followings
	u.Followers = user.Followers
	u.Followings = user.Followings
	u.Photos = user.Photos
}

//Comment 

type Comment struct{
	IDComment uint64 	`json:"id"`
	IDUser uint64 		`json:"idUser"`
	Username string		`json:"username"`
	IDPhoto uint64 		`json:"idPhoto"`
	CommentText string 	`json:"comment"`

}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		IDComment: 	c.IDComment,
		IDUser:    	c.IDUser,
		Username:	c.Username,
		IDPhoto:   	c.IDPhoto,
		CommentText:c.CommentText,
	}
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.IDComment = comment.IDComment
	c.IDUser = comment.IDUser
	c.Username = comment.Username
	c.IDPhoto = comment.IDPhoto
	c.CommentText = comment.CommentText
}

func (c *Comment) isValid() bool {
	length := len([]rune(c.CommentText))
	return 1 <= length && length <= 200
}




//UserLogin

type UserLogin struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

func (u *UserLogin) FromDatabase(user database.UserLogin) {
	u.ID = user.ID
	u.Username = user.Username
}

func (u *UserLogin) ToDatabase() database.UserLogin {
	return database.UserLogin{
		ID:  u.ID,
		Username: u.Username,
	}
}

// The username of the user is valid if it is between 3 and 20 caracters
func (u *UserLogin) isValid() bool {
	length := len([]rune(u.Username))
	return 3 <= length && length <= 20
}

//Like
type Like struct {
	ID       uint64 `json:"id"`
	IDUser uint64 `json:"idUser"`
	IDPhoto uint64 `json:"idPhoto"`

}

func (u *Like) FromDatabase(user database.Like) {
	u.ID = user.ID
	u.IDUser = user.IDUser
	u.IDPhoto = user.IDPhoto

}

func (u *Like) ToDatabase() database.Like {
	return database.Like{
		ID:  u.ID,
		IDUser: u.IDUser,
		IDPhoto: u.IDPhoto,

	}
}
