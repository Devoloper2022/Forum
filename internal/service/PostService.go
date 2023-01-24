package service

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"time"
)

type Post interface {
	CreatePost(dto dto.PostDto, categories []string) error
	GetAllPosts() ([]dto.PostDto, error)
	GetAllPostsByUserID(userId int64) ([]dto.PostDto, error)
	GetPost(postId int64) (dto.PostDto, error)
	UpdatePost(post dto.PostDto) error

	// GetPostByFilter(query map[string][]string) ([]models.PostInfo, error)
}

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(post dto.PostDto, categories []string) error {
	// t := time.Now().Format("d MMM yyyy HH:mm:ss")
	t := time.Now()
	newPost := models.Post{
		Title:   post.Title,
		Text:    post.Text,
		Date:    t,
		Like:    0,
		Dislike: 0,
		UserID:  post.User.ID,
	}
	pid, err := s.repo.CreatePost(newPost)
	if err != nil {
		return err
	}
	err = s.repo.CreatePostCategory(pid, categories)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostService) UpdatePost(post dto.PostDto) error {
	// check user ouwnership
	upPost := models.Post{
		ID:    post.ID,
		Title: post.Text,
		Text:  post.Text,
	}

	return s.repo.UpdatePost(upPost)
}

func (s *PostService) GetPost(postId int64) (dto.PostDto, error) {
	post, err := s.repo.GetPost(postId)
	if err != nil {
		return dto.PostDto{}, nil
	}

	t := post.Date.Format("d MMM yyyy HH:mm:ss")
	user, err := s.repo.GetUser(post.UserID)
	if err != nil {
		return dto.PostDto{}, nil
	}

	return dto.PostDto{
		ID:    postId,
		Title: post.Title,
		Text:  post.Text,
		Date:  t,
		User:  user,
	}, nil
}

func (s *PostService) GetAllPosts() ([]dto.PostDto, error) {
	list, err := s.repo.GetAllPosts()
	if err != nil {
		return nil, nil
	}

	var listDto []dto.PostDto

	for _, p := range list {
		var post dto.PostDto
		post.ID =
	
	}
}

func (s *PostService) GetAllPostsByUserID(userId int64) ([]dto.PostDto, error) {
}

func ()  {
	
}
