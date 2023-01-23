package service

import (
	"forum/internal/models"
	"forum/internal/repository"
)

type Post interface {
	CreatePost(post models.Post) (int, error)
	// GetAllPosts() ([]models.PostInfo, error)
	// GetPost(id int) (models.PostInfo, error)
	// CreatePostCategory(id int, categories []string) error
	// GetAllCategories() ([]models.Category, error)
	// GetPostByFilter(query map[string][]string) ([]models.PostInfo, error)
}

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post models.Post) (int, error) {
	return s.repo.CreatePost(post)
}
