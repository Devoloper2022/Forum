package dto

type PostLikeDto struct {
	ID     int64
	UserID int64
	PostID int64
	Result bool
}

type CommentLikeDto struct {
	ID        int64
	UserID    int64
	CommentID int64
	Result    bool
}
