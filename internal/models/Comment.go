package models

type Comment struct {
	ID      int64  `json:"id"`
	Text    string `json:"text"`
	Date    string `json:"date"`
	Like    int64  `json:"like"`
	Dislike int64  `json:"dislike"`
	UserID  int64  `json:"userId"`
	PostID  int64  `json:"postId"`
}
