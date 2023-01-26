package service

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"strconv"
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
	repo       repository.Post
	categories repository.Category
	user       repository.User
}

func NewPostService(repo repository.Post, categories repository.Category, user repository.User) *PostService {
	return &PostService{
		repo:       repo,
		categories: categories,
		user:       user,
	}
}

func (s *PostService) CreatePost(post dto.PostDto, categories []string) error {
	t := time.Now().Format("d MMM yyyy HH:mm:ss")

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
	var listID []int64
	for _, cid := range categories {
		num, err := strconv.Atoi(cid)
		if err != nil {
			return err
		}

		listID = append(listID, int64(num))
	}

	if err != nil {
		return fmt.Errorf("repository : create PostCategory  checker 2: %w", err)
	}
	err = s.categories.CreatePostCategory(pid, listID)

	if err != nil {
		return err
	}

	return nil
} // done

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

	// t := post.Date.Format("d MMM yyyy HH:mm:ss")
	user, err := s.user.GetUser(post.UserID)
	fmt.Println(user) // fix

	if err != nil {
		return dto.PostDto{}, nil
	}

	return dto.PostDto{
		ID:    postId,
		Title: post.Title,
		Text:  post.Text,
		Date:  post.Date,
		User:  dto.UserDto{},
	}, nil
} // fix

func (s *PostService) GetAllPosts() ([]dto.PostDto, error) {
	list, err := s.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Level service result: %s\n", list)
	var listDto []dto.PostDto

	for _, p := range list {
		dto := dto.GetPostDto(p, dto.UserDto{}, nil, nil) // FIX
		listDto = append(listDto, dto)
	}

	return listDto, nil
}

func (s *PostService) GetAllPostsByUserID(userId int64) ([]dto.PostDto, error) {
	list, err := s.repo.GetPostByUserID(userId)
	if err != nil {
		return nil, nil
	}

	var listDto []dto.PostDto

	for _, p := range list {
		dto := dto.GetPostDto(p, dto.UserDto{}, nil, nil) // FIX
		listDto = append(listDto, dto)
	}

	return listDto, nil
}
