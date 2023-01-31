package service

import (
	"database/sql"
	"errors"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"forum/internal/repository"
)

type LikePost interface {
	LikePost(data models.PostLike) error
	DislikePost(data models.PostLike) error
}

type LikePostService struct {
	repo repository.LikePost
	post repository.Post
}

func NewLikePostService(repo repository.LikePost, post repository.Post) *LikePostService {
	return &LikePostService{
		repo: repo,
		post: post,
	}
}

func (s *LikePostService) LikePost(data models.PostLike) error {
	if data.DisLike == data.Like {
		return dto.ErrLikeDislike
	}

	if data.DisLike {
		return dto.ErrDislike
	}

	post, err := s.post.GetPost(data.PostID)
	if err != nil {
		return err
	}

	modelLike, err := s.repo.GetPostLike(data.PostID, data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = s.repo.CreatePostLike(data)

			if err != nil {
				return err
			}

			post.Like++

			err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

			if err != nil {
				return err
			}

			return nil

		} else {
			return err
		}
	}

	if modelLike.Like && modelLike.DisLike == false {
		post.Like--

		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = false

		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else if modelLike.Like == false && modelLike.DisLike == false {
		post.Like++
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = true

		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else if modelLike.Like == false && modelLike.DisLike == true {
		post.Like++
		post.Dislike--
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = true
		modelLike.DisLike = false

		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	} else {
		post.Like--
		post.Dislike++
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.Like = false
		modelLike.DisLike = true

		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}

		return nil
	}

	return nil
}

func (s *LikePostService) DislikePost(data models.PostLike) error {
	if data.DisLike == data.Like {
		return dto.ErrLikeDislike
	}

	if data.Like {
		return dto.ErrLike
	}

	post, err := s.post.GetPost(data.PostID)
	if err != nil {
		return err
	}

	modelLike, err := s.repo.GetPostLike(data.PostID, data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = s.repo.CreatePostLike(data)

			if err != nil {
				return err
			}

			post.Dislike++

			err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

			if err != nil {
				return err
			}

			return nil

		} else {
			return err
		}
	}

	if modelLike.DisLike && modelLike.Like == false {
		if post.Dislike > 0 {
			post.Dislike--
		}

		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = false
		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else if modelLike.DisLike == false && modelLike.Like == false {
		post.Dislike++
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = true
		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else if modelLike.DisLike == false && modelLike.Like == true {

		post.Dislike++
		post.Like--
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = true
		modelLike.Like = false
		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	} else {
		post.Dislike--
		post.Like++
		err = s.repo.UpdatePostTable(post.ID, post.Like, post.Dislike)

		if err != nil {
			return err
		}

		modelLike.DisLike = false
		modelLike.Like = true
		err = s.repo.UpdatePostLike(modelLike)
		if err != nil {
			return nil
		}
		return nil
	}

	return nil
}
