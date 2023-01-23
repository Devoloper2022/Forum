package models

import "time"

type Comment struct {
	ID     int64     `json:"id"`
	Text   string    `json:"text"`
	Date   time.Time `json:"date"`
	UserID int64     `json:"userId"`
	PostID int64     `json:"postId"`
}
