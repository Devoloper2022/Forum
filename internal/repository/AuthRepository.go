package repository

import (
	"fmt"
	"forum/internal/models"
)

type Autorization interface {
	Create(post models.Post) (int, error)
	// GetAllPosts() ([]models.PostInfo, error)
	// GetPost(id int) (models.PostInfo, error)
	// CreatePostCategory(postId int, categories []string) error
	// GetAllCategories() ([]models.Category, error)
	// GetPostsByMostLikes() ([]models.PostInfo, error)
	// GetPostsByLeastLikes() ([]models.PostInfo, error)
	// GetPostByCategory(category string) ([]models.PostInfo, error)
}

func (r *Database) Create(post models.Post) (int, error) {
	res, err := r.db.Exec("INSERT INTO posts (Title,Content,UserId) VALUES (?,?,?)", post.Title, post.Text, post.UserID)
	if err != nil {
		return 0, fmt.Errorf("repository : create post : %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
