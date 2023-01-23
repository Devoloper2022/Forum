package repository

import (
	"fmt"
	"forum/internal/models"
)

type Post interface {
	CreatePost(post models.Post) (int, error)
	// GetAllPosts() ([]models.PostInfo, error)
	// GetPost(id int) (models.PostInfo, error)
	// CreatePostCategory(postId int, categories []string) error
	// GetAllCategories() ([]models.Category, error)
	// GetPostsByMostLikes() ([]models.PostInfo, error)
	// GetPostsByLeastLikes() ([]models.PostInfo, error)
	// GetPostByCategory(category string) ([]models.PostInfo, error)
}

func (r *Database) CreatePost(post models.Post) (int, error) {
	com := ("INSERT INTO posts (Title,Text,UserId) VALUES (?,?,?)")
	query, err := r.db.Prepare(com)
	if err != nil {
		return 0, fmt.Errorf("repository : create post : %w", err)
	}
	res, err := query.Exec(post.Title, post.Text, post.UserID)

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
