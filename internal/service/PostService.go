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
	CreatePost(dto dto.PostDto, categories []string) (int64, error)
	GetAllPosts() ([]dto.PostDto, error)
	GetAllPostsByUserID(userId int64) ([]dto.PostDto, error)
	GetPost(postId int64) (dto.PostDto, error)
	UpdatePost(post dto.PostDto) error
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

func (s *PostService) CreatePost(post dto.PostDto, categories []string) (int64, error) {
	t := time.Now().Format(time.RFC1123)

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
		return pid, err
	}
	var listID []int64
	for _, cid := range categories {
		num, err := strconv.Atoi(cid)
		if err != nil {
			return pid, err
		}

		listID = append(listID, int64(num))
	}

	if err != nil {
		return pid, fmt.Errorf("service : create PostCategory  checker 2: %w", err)
	}
	err = s.categories.CreatePostCategory(pid, listID)

	if err != nil {
		return pid, err
	}

	return pid, nil
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
		return dto.PostDto{}, err
	}

	user, err := s.user.GetUser(post.UserID)
	if err != nil {
		return dto.PostDto{}, err
	}

	cat, err := s.categories.GetAllCategoriesByPostId(post.ID)
	if err != nil {
		return dto.PostDto{}, err
	}

	return dto.PostDto{
		ID:        postId,
		Title:     post.Title,
		Text:      post.Text,
		Date:      post.Date,
		User:      dto.GetUserDto(user),
		Categorys: cat,
	}, nil
} // fix

func (s *PostService) GetAllPosts() ([]dto.PostDto, error) {
	list, err := s.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}

	var listDto []dto.PostDto

	for _, p := range list {
		dto := dto.GetPostDto(p, dto.UserDto{}, dto.PostLikeDto{}, nil) // FIX
		listDto = append(listDto, dto)
	}

	return listDto, nil
} // checked // fix

func (s *PostService) GetAllPostsByUserID(userId int64) ([]dto.PostDto, error) {
	list, err := s.repo.GetPostByUserID(userId)
	if err != nil {
		return nil, nil
	}

	var listDto []dto.PostDto

	for _, p := range list {
		dto := dto.GetPostDto(p, dto.UserDto{}, dto.PostLikeDto{}, nil) // FIX
		listDto = append(listDto, dto)
	}

	return listDto, nil
}
