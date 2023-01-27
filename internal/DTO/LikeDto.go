package dto

import "forum/internal/models"

type PostLikeDto struct {
	ID      int64
	User    UserDto
	PostID  int64
	Like    bool
	DisLike bool
}

func (dto *PostLikeDto) GetPostLikeModel() models.PostLike {
	return models.PostLike{
		ID:      dto.ID,
		UserID:  dto.User.ID,
		PostID:  dto.PostID,
		Like:    dto.Like,
		DisLike: dto.DisLike,
	}
}

func GetPostLikeDto(m models.PostLike, u UserDto) PostLikeDto {
	return PostLikeDto{
		ID:      m.ID,
		User:    u,
		PostID:  m.PostID,
		Like:    m.Like,
		DisLike: m.DisLike,
	}
}

type CommentLikeDto struct {
	ID        int64
	User      UserDto
	CommentID int64
	Like      bool
	DisLike   bool
}

func (dto *CommentLikeDto) GetUCommentLikeModel() models.CommentLike {
	return models.CommentLike{
		ID:        dto.ID,
		UserID:    dto.User.ID,
		CommentID: dto.CommentID,
		Like:      dto.Like,
		DisLike:   dto.DisLike,
	}
}

func GetCommentLikeDto(m models.CommentLike, u UserDto) CommentLikeDto {
	return CommentLikeDto{
		ID:        m.ID,
		User:      u,
		CommentID: m.CommentID,
		Like:      m.Like,
		DisLike:   m.DisLike,
	}
}
