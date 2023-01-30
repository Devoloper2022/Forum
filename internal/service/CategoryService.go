package service

import (
	"forum/internal/models"
	"forum/internal/repository"
)

type Category interface {
	GetAllCategories() ([]models.Category, error)
}

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}
