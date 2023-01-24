package service

import (
	dto "forum/internal/DTO"
	"forum/internal/repository"
)

type User interface {
	CreatePost(dto dto.PostDto, categories []string) error
	GetAllPosts() ([]dto.PostDto, error)
	GetAllPostsByUserID(userId int64) ([]dto.PostDto, error)
	GetPost(postId int64) (dto.PostDto, error)
	UpdatePost(post dto.PostDto) error

	// GetPostByFilter(query map[string][]string) ([]models.PostInfo, error)
}

type UserService struct {
	repo repository.Post
}

func NewUserService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}
