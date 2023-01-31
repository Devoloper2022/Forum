package service

import (
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
)

type Like interface {
	CreateLikeComment(data models.CommentLike) (models.CommentLike, error)
	LikeComment(data models.CommentLike) (models.CommentLike, error)
	DislikeComment(data models.CommentLike) (models.CommentLike, error)

	CreateLikePost(data models.PostLike) (models.PostLike, error)
	LikePost(data models.PostLike) (models.PostLike, error)
	DislikePost(data models.PostLike) (models.PostLike, error)
}

type LikeService struct {
	repo    repository.Like
	comment repository.Comment
	post    repository.Post
}

func NewLikeService(repo repository.Like, comment repository.Comment, post repository.Post) *LikeService {
	return &LikeService{
		repo:    repo,
		comment: comment,
		post:    post,
	}
}

func (s *LikeService) CreateLikeComment(data models.CommentLike) (models.CommentLike, error) {
	if data.Like == data.DisLike {
		return models.CommentLike{}, dto.ErrLikeDislike
	}

	model, err := s.repo.CreateCommentLike(data)

	if err != nil {
		return models.CommentLike{}, err
	}

	return model, nil
}
func (s *LikeService) LikeComment(data models.CommentLike) (models.CommentLike, error) {

	comment, err := s.comment.GetComment(data.CommentID)

	if err != nil {
		return models.CommentLike{}, err
	}

	if data.DisLike == true {
		return data, dto.ErrDislike
	}
	// comment.Like=comment.Like+1
	// comment.Dislike=comment.Dislike-1
	comment.Like++
	comment.Dislike--

	err = s.repo.UpdateCommentL(comment.ID, comment.Like, comment.Dislike)

	if err != nil {
		return models.CommentLike{}, err
	}

	upModel, err := s.repo.UpdateCommentLike(data)
	if err != nil {
		return models.CommentLike{}, err
	}

	return upModel, nil
}

func (s *LikeService) DislikeComment(data models.CommentLike) (models.CommentLike, error) {

	comment, err := s.comment.GetComment(data.CommentID)

	if err != nil {
		return models.CommentLike{}, err
	}

	if data.Like == true {
		return data, dto.ErrLike
	}
	// comment.Like-=1
	// comment.Dislike+=1
	comment.Like--
	comment.Dislike++

	err = s.repo.UpdateCommentL(comment.ID, comment.Like, comment.Dislike)

	if err != nil {
		return models.CommentLike{}, err
	}

	upModel, err := s.repo.UpdateCommentLike(data)
	if err != nil {
		return models.CommentLike{}, err
	}

	return upModel, nil

}

////post

func (s *LikeService) CreateLikePost(data models.PostLike) (models.PostLike, error) {
	if data.Like == data.DisLike {
		return models.PostLike{}, dto.ErrLikeDislike
	}

	model, err := s.repo.CreatePostLike(data)

	if err != nil {
		return models.PostLike{}, err
	}

	return model, nil

}
func (s *LikeService) LikePost(data models.PostLike) (models.PostLike, error) {
	post, err := s.post.GetPost(data.PostID)

	if err != nil {
		return models.PostLike{}, err
	}

	if data.DisLike == true {
		return data, dto.ErrDislike
	}

	post.Like++
	post.Dislike--

	err = s.repo.UpdateCommentL(post.ID, post.Like, post.Dislike)

	if err != nil {
		return models.PostLike{}, err
	}

	upModel, err := s.repo.UpdatePostLike(data)
	if err != nil {
		return models.PostLike{}, err
	}

	return upModel, nil
}
func (s *LikeService) DislikePost(data models.PostLike) (models.PostLike, error) {
	post, err := s.post.GetPost(data.PostID)

	if err != nil {
		return models.PostLike{}, err
	}

	if data.Like == true {
		return data, dto.ErrDislike
	}

	post.Like--
	post.Dislike++

	err = s.repo.UpdateCommentL(post.ID, post.Like, post.Dislike)

	if err != nil {
		return models.PostLike{}, err
	}

	upModel, err := s.repo.UpdatePostLike(data)
	if err != nil {
		return models.PostLike{}, err
	}

	return upModel, nil
}
