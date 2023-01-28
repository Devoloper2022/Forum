package service

import (
	"forum/internal/repository"
)

const mailValidation = `[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`

type Service struct {
	Autorization
	Post
	Category
	User
	Comment
	// VotePost
	// VoteComment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization, repos.User),
		Post:         NewPostService(repos.Post, repos.Category, repos.User),
		Category:     NewCategoryService(repos.Category),
		User:         NewUserService(repos.User),
		Comment:      NewCommentService(repos.Comment, repos.User),
		// VotePost:     NewVotePostService(repos.VotePost),
		// VoteComment:  NewVoteCommentService(repos.VoteComment),
	}
}
