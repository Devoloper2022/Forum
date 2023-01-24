package service

import (
	"forum/internal/repository"
)

type Service struct {
	Autorization
	Post
	CategoryService
	// Comment
	// VotePost
	// VoteComment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization:    NewAuthService(repos.Autorization),
		Post:            NewPostService(repos.Post),
		CategoryService: *NewCategoryService(repos.Category),
		// Comment:      NewCommentService(repos.Comment),
		// VotePost:     NewVotePostService(repos.VotePost),
		// VoteComment:  NewVoteCommentService(repos.VoteComment),
	}
}
