package service

import (
	dto "forum/internal/DTO"
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

func (s *CategoryService) GetAllCategories() ([]dto.CategoryDto, error) {
	var result []dto.CategoryDto
	var dto dto.CategoryDto
	list, err := s.repo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	for _, c := range list {
		dto.ID = c.ID
		dto.Title = c.Title
		result = append(result, dto)
	}
	return result, nil
}
