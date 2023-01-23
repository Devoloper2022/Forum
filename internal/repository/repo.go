package repository

import (
	"database/sql"
)

const createPost = "INSERT INTO posts (Title,Content,UserId) VALUES (?,?,?)"

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

type Repository struct {
	Autorization
	Post
	// Comment
	// VotePost
	// VoteComment
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Autorization: NewDatabase(db),
		Post:         NewDatabase(db),
		// 		Comment:      NewCommentRepository(db),
		// 		VotePost:     NewVotePostRepository(db),
		// 		VoteComment:  NewVoteCommentRepository(db),
	}
}
