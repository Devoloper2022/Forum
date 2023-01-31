package service

import (
	"database/sql"
	"errors"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
)

type Like interface {
	LikeComment(data models.CommentLike) error
	DislikeComment(data models.CommentLike) error
}

type LikeService struct {
	repo    repository.LikeComment
	comment repository.Comment
}

func NewLikeService(repo repository.LikeComment, comment repository.Comment) *LikeService {
	return &LikeService{
		repo:    repo,
		comment: comment,
	}
}

func (s *LikeService) LikeComment(data models.CommentLike) error {
	if data.DisLike == data.Like {
		return dto.ErrLikeDislike
	}

	if data.DisLike {
		return dto.ErrDislike
	}

	comment, err := s.comment.GetComment(data.CommentID)
	if err != nil {
		return err
	}

	modelLike, err := s.repo.GetCommentLike(data.UserID, data.CommentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = s.repo.CreateCommentLike(data)

			if err != nil {
				return err
			}

			comment.Like++

			err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

			if err != nil {
				return err
			}

			return nil

		} else {
			return err
		}
	}

	if modelLike.Like && modelLike.DisLike == false {
		comment.Like--

		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = false

		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else if modelLike.Like == false && modelLike.DisLike == false {
		comment.Like++
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = true

		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else if modelLike.Like == false && modelLike.DisLike == true {
		comment.Like++
		comment.Dislike--
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = true
		modelLike.DisLike = false

		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else {
		comment.Like--
		comment.Dislike++
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = false
		modelLike.DisLike = true

		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	}

	return nil
}

func (s *LikeService) DislikeComment(data models.CommentLike) error {
	if data.DisLike == data.Like {
		return dto.ErrLikeDislike
	}

	if data.Like {
		return dto.ErrLike
	}

	comment, err := s.comment.GetComment(data.CommentID)
	if err != nil {
		return err
	}

	modelLike, err := s.repo.GetCommentLike(data.UserID, data.CommentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = s.repo.CreateCommentLike(data)

			if err != nil {
				return err
			}

			comment.Dislike++

			err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

			if err != nil {
				return err
			}

			return nil

		} else {
			return err
		}
	}

	if modelLike.DisLike && modelLike.Like == false {
		if comment.Dislike > 0 {
			comment.Dislike--
		}

		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = false
		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else if modelLike.DisLike == false && modelLike.Like == false {
		comment.Dislike++
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = true
		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else if modelLike.DisLike == false && modelLike.Like == true {

		comment.Dislike++
		comment.Like--
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = true
		modelLike.Like = false
		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else {
		comment.Dislike--
		comment.Like++
		err = s.repo.UpdateCommentTable(comment.ID, comment.Like, comment.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = false
		modelLike.Like = true
		err = s.repo.UpdateCommentLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	}

	return nil
}
