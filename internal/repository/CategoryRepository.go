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

// func (r *Database) getUserById(id int64) (m.User, error) {
// 	findId := r.db.QueryRow("SELECT * FROM users WHERE  users_id = ?", id)
// 	user := m.User{}
// 	err := findId.Scan(&user.ID, user.Email, user.Password, user.Username)
// 	if err != nil {
// 		log.Println("Repo ==> GetUserById")
// 		return user, err
// 	}
// 	return user, nil
// }
