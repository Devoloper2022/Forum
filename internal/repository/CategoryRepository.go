package repository

import (
	"database/sql"
	"errors"
	"forum/internal/models"
)

type Category interface {
	CreatePostCategory(postId int64, categories []int64) error
	GetAllCategories() ([]models.Category, error)
	GetAllCategoriesByPostId(postId int64) ([]models.Category, error)
}

func (r *Database) GetAllCategories() ([]models.Category, error) {
	rows, err := r.db.Query("SELECT * FROM categories")
	if err != nil {
		return []models.Category{}, err
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
	defer st.Close()

	if err != nil {
		return err
	}

	for _, cid := range categories {
		_, err = st.Exec(postId, cid)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (r *Database) CreatePostCategory(postId int64, categories []int64) error {
// 	query := ("INSERT INTO categoriesPost (PostID,CategoryID) VALUES (?,?)")
// 	st, err := r.db.Prepare(query)
// 	if err != nil {
// 		return fmt.Errorf("repository : create PostCategory  checker 1: %w", err)
// 	}
// 	defer st.Close()
// 	for _, cid := range categories {
// 		_, err = st.Exec(postId, cid)
// 		if err != nil {
// 			return fmt.Errorf("repository : create PostCategory checker 3: %w", err)
// 		}
// 	}
// 	return nil
// }

func (r *Database) GetAllCategoriesByPostId(postId int64) ([]models.Category, error) {
	query := ("SELECT categories.ID, categories.Title FROM categories INNER JOIN categoriesPost ON categories.ID = categoriesPost.CategoryID WHERE categoriesPost.PostID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return nil, err
	}

	rows, err := st.Query(postId)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var list []models.Category
	for rows.Next() {
		var cat models.Category
		if err != rows.Scan(&cat.ID, &cat.Title) {
			return nil, err
		}
		list = append(list, cat)
	}

	return list, nil
}
