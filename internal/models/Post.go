package models

import "time"

type Post struct {
	ID      int64     `json:"ID"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	Date    time.Time `json:"data"`
	Like    int64     `json:"like"`
	Dislike int64     `json:"dislike"`
	UserID  int64     `json:"userId"`
}
