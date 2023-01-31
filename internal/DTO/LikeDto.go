package dto

import "forum/internal/models"

type PostLikeDto struct {
	ID      int64
	User    int64
	Like    bool
	DisLike bool
}

func (dto *PostLikeDto) GetPostLikeModel() models.PostLike {
	return models.PostLike{
		ID:      dto.ID,
		UserID:  dto.User,
		Like:    dto.Like,
		DisLike: dto.DisLike,
	}
}

func GetPostLikeDto(m models.PostLike) PostLikeDto {
	return PostLikeDto{
		ID:      m.ID,
		User:    m.ID,
		Like:    m.Like,
		DisLike: m.DisLike,
	}
}

type CommentLikeDto struct {
	ID      int64
	User    int64
	Like    bool
	DisLike bool
}

func (dto *CommentLikeDto) GetCommentLikeModel() models.CommentLike {
	return models.CommentLike{
		ID:      dto.ID,
		UserID:  dto.User,
		Like:    dto.Like,
		DisLike: dto.DisLike,
	}
}

func GetCommentLikeDto(m models.CommentLike) CommentLikeDto {
	return CommentLikeDto{
		ID:      m.ID,
		User:    m.ID,
		Like:    m.Like,
		DisLike: m.DisLike,
	}
}
