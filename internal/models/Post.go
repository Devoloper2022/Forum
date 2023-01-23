package models

import "time"

type Post struct {
	ID        int64      `json:"ID"`
	Title     string     `json:"title"`
	Text      string     `json:"text"`
	Date      time.Time  `json:"data"`
	UserID    int64      `json:"userId"`
	Like      int64      `json:"like"`
	Dislike   int64      `json:"dislike"`
	Likes     []Post     `json:"likes"`
	Categorys []Category `json:"categorys"`
	// Comments  []Comment  `json:"comments"`
}
