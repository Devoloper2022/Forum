package models

type PostLike struct {
	ID      int64
	UserID  int64
	PostID  int64
	Like    bool
	DisLike bool
}

type CommentLike struct {
	ID        int64
	UserID    int64
	CommentID int64
	Like      bool
	DisLike   bool
}
