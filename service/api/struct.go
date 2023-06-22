package api

import (
	"time"

	"git.wasaphoto.ivi/wasaphoto/service/database"
)

// Error Message
type JSONErrorMsg struct {
	Message string
}

// Photo
type Photo struct {
	PhotoID  uint64    `json:"id"`
	UserID   uint64    `json:"userId"`
	Username string    `json:"username"`
	DateTime time.Time `json:"DateTime"`
	Likes    uint64    `json:"likes"`
	Comments uint64    `json:"comments"`
	Path     string    `json:"path"`
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.PhotoID = photo.PhotoID
	p.UserID = photo.UserID
	p.Username = photo.Username
	p.DateTime = photo.DateTime
	p.Likes = photo.Likes
	p.Comments = photo.Comments
	p.Path = photo.Path
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		PhotoID:  p.PhotoID,
		UserID:   p.UserID,
		Username: p.Username,
		DateTime: p.DateTime,
		Likes:    p.Likes,
		Comments: p.Comments,
		Path:     p.Path,
	}
}

// User
type User struct {
	UserID          uint64   `json:"id"`
	Username        string   `json:"username"`
	PhotosCount     uint64   `json:"photosCount"`
	FollowersCount  uint64   `json:"followersCount"`
	FollowingsCount uint64   `json:"followingCount"`
	PhotosId        []uint64 `json:"photosId"`
	PhotosPath      []string `json:"photos"`
	Followers       []string `json:"followers"`
	Followings      []string `json:"followings"`
}

func (user *User) ToDatabase() database.User {
	return database.User{
		UserID:          user.UserID,
		Username:        user.Username,
		PhotosCount:     user.PhotosCount,
		FollowersCount:  user.FollowersCount,
		FollowingsCount: user.FollowingsCount,
		Followers:       user.Followers,
		Followings:      user.Followings,
		PhotosId:        user.PhotosId,
		PhotosPath:      user.PhotosPath,
	}
}
func (u *User) FromDatabase(user database.User) {
	u.UserID = user.UserID
	u.Username = user.Username
	u.PhotosCount = user.PhotosCount
	u.Followers = user.Followers
	u.Followings = user.Followings
	u.Followers = user.Followers
	u.Followings = user.Followings
	u.PhotosId = user.PhotosId
	u.PhotosPath = user.PhotosPath

}

// Comment

type Comment struct {
	IDComment   uint64 `json:"id"`
	UserID      uint64 `json:"userId"`
	Username    string `json:"username"`
	PhotoID     uint64 `json:"photoId"`
	CommentText string `json:"comment"`
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		IDComment:   c.IDComment,
		UserID:      c.UserID,
		Username:    c.Username,
		PhotoID:     c.PhotoID,
		CommentText: c.CommentText,
	}
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.IDComment = comment.IDComment
	c.UserID = comment.UserID
	c.Username = comment.Username
	c.PhotoID = comment.PhotoID
	c.CommentText = comment.CommentText
}

func (c *Comment) isValid() bool {
	length := len([]rune(c.CommentText))
	return 1 <= length && length <= 200
}

// UserLogin

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
		ID:       u.ID,
		Username: u.Username,
	}
}

// The username of the user is valid if it is between 3 and 20 caracters
func (u *UserLogin) isValid() bool {
	length := len([]rune(u.Username))
	return 3 <= length && length <= 20
}

// Like
type Like struct {
	ID      uint64 `json:"id"`
	UserID  uint64 `json:"userId"`
	PhotoID uint64 `json:"photoId"`
}

func (u *Like) FromDatabase(user database.Like) {
	u.ID = user.ID
	u.UserID = user.UserID
	u.PhotoID = user.PhotoID

}

func (u *Like) ToDatabase() database.Like {
	return database.Like{
		ID:      u.ID,
		UserID:  u.UserID,
		PhotoID: u.PhotoID,
	}
}

type Ban struct {
	BannedUser string `json:"bannedUser"`
}

func (b *Ban) FromDatabase(ban database.Ban) {
	b.BannedUser = ban.BannedUser

}

func (b *Ban) ToDatabase() database.Ban {
	return database.Ban{
		BannedUser: b.BannedUser,
	}
}

type FollwedUsers struct {
	Followed []string `json:"followedusers"`
}
