package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/internal/models"
)

type Category interface {
	// CreateCategory(Category models.Category) (int, error)
	// GetAllPosts() ([]models.PostInfo, error)
	// GetPost(id int) (models.PostInfo, error)
	// CreatePostCategory(postId int, categories []string) error
	GetAllCategories() ([]models.Category, error)
	CreatePostCategory(postId int64, categories []int64) error
	// GetPostsByMostLikes() ([]models.PostInfo, error)
	// GetPostsByLeastLikes() ([]models.PostInfo, error)
	// GetPostByCategory(category string) ([]models.PostInfo, error)
}

func (r *Database) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query("SELECT Id,Title FROM categories")
	if err != nil {
		return []models.Category{}, fmt.Errorf("repository : GetAllCategories: %w", err)
	}
	var categoryList []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Title)
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Category{}, err
		} else if err != nil {
			return []models.Category{}, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}

func (r *Database) CreatePostCategory(postId int64, categories []int64) error {
	query := ("INSERT INTO categoriesPost (PostID,CategoryID) VALUES (?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("repository : create PostCategory  checker 1: %w", err)
	}
	defer st.Close()

	for _, cid := range categories {
		_, err = st.Exec(postId, cid)
		if err != nil {
			return fmt.Errorf("repository : create PostCategory checker 3: %w", err)
		}
	}
	return nil
}
