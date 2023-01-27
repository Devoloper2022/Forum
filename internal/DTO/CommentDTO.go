package dto

import "time"

type Comment struct {
	ID     int64          `json:"id"`
	Text   string         `json:"text"`
	Date   time.Time      `json:"date"`
	User   UserDto        `json:"userId"`
	PostID int64          `json:"postId"`
	Likes  CommentLikeDto `json:"likes"`
}
