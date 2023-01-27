package service

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
	"time"
)

type Comment interface {
	CreateComment(dto dto.CommentDto) error
	// GetAllCommentsByPostId(postID int64) ([]dto.CommentDto, error)

	// GetComment(commentID int64) (dto.CommentDto, error)
	// GetAllCommentsByUserID(userId int64) ([]dto.CommentDto, error)
}

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
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

// func (s *CommentService) GetAllCommentsByPostId(postID int64) ([]dto.CommentDto, error) {}
// func (s *CommentService) GetComment(commentID int64) (dto.CommentDto, error)            {}
// func (s *CommentService) GetAllCommentsByUserID(userId int64) ([]dto.CommentDto, error) {}
