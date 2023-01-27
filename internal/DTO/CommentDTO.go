package dto

type CommentDto struct {
	ID     int64          `json:"id"`
	Text   string         `json:"text"`
	Date   string         `json:"date"`
	User   UserDto        `json:"userId"`
	PostID int64          `json:"postId"`
	Likes  CommentLikeDto `json:"likes"`
}
