package service

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"time"
)

type Comment interface {
	CreateComment(dto dto.CommentDto) error
	GetAllCommentsByPostId(postID int64) ([]dto.CommentDto, error)

	// GetComment(commentID int64) (dto.CommentDto, error)
	// GetAllCommentsByUserID(userId int64) ([]dto.CommentDto, error)
}

type CommentService struct {
	repo repository.Comment
	user repository.User
	// like repository.Like
}

func NewCommentService(repo repository.Comment, user repository.User) *CommentService {
	return &CommentService{
		repo: repo,
		user: user,
		// like: like,
	}
}

func (s *CommentService) CreateComment(dto dto.CommentDto) error {
	t := time.Now().Format(time.RFC1123)

	newComment := models.Comment{
		Text:    dto.Text,
		Date:    t,
		Like:    0,
		Dislike: 0,
		PostID:  dto.PostID,
		UserID:  dto.User.ID,
	}
	err := s.repo.CreateComment(newComment)
	if err != nil {
		return err
	}

	return nil
} // done

func (s *CommentService) GetAllCommentsByPostId(postID int64) ([]dto.CommentDto, error) {
	list, err := s.repo.GetAllCommentByPostID(postID)
	if err != nil {
		return nil, err
	}

	var dtoList []dto.CommentDto

	for _, m := range list {
		user, err := s.user.GetUser(m.UserID)
		if err != nil {
			return nil, err
		}

		dtoUser := dto.GetUserDto(user)
		dtoComment := dto.GetCommentDto(m, dtoUser, models.CommentLike{})
		dtoList = append(dtoList, dtoComment)
	}

	return dtoList, nil
}

// func (s *CommentService) GetComment(commentID int64) (dto.CommentDto, error)            {}
// func (s *CommentService) GetAllCommentsByUserID(userId int64) ([]dto.CommentDto, error) {}
