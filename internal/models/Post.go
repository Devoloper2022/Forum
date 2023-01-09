package models

import "time"

type Post struct {
	ID      int64     `json:"ID"`
	Title   string    `json:"title"`
	Text    string    `json:"text"`
	Data    time.Time `json:"data"`
	UserId  int64     `json:"userId"`
	Likes   int64     `json:"likes"`
	Dislike int64     `json:"dislike"`

	// Categorys []Category `json:"category"`
	// Comments  []Comment  `json:"comments"`
}
