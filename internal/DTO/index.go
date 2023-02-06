package dto

import "forum/internal/models"

type Index struct {
	List []models.Category
	Post []PostDto
}
