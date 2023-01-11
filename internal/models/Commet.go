package internal

import "time"

type Comment struct {
	ID     int64     `json:"id"`
	Text   string    `json:"text"`
	Data   time.Time `json:"date"`
	UserId int64     `json:"userId"`
	PostId int64     `json:"postId"`
}
