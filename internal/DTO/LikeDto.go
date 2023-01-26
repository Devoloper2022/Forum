package dto

import "forum/internal/models"

type PostLikeDto struct {
	ID     int64
	User   UserDto
	PostID int64
	Result bool
}

func (dto *PostLikeDto) GetPostLikeModel() models.PostLike {
	return models.PostLike{
		ID:     dto.ID,
		UserID: dto.User.ID,
		PostID: dto.PostID,
		Result: dto.Result,
	}
}

func GetPostLikeDto(m models.PostLike, u UserDto) PostLikeDto {
	return PostLikeDto{
		ID:     m.ID,
		User:   u,
		PostID: m.PostID,
		Result: m.Result,
	}
}

type CommentLikeDto struct {
	ID        int64
	User      UserDto
	CommentID int64
	Result    bool
}

func (dto *CommentLikeDto) GetUCommentLikeModel() models.CommentLike {
	return models.CommentLike{
		ID:        dto.ID,
		UserID:    dto.User.ID,
		CommentID: dto.CommentID,
		Result:    dto.Result,
	}
}

func GetCommentLikeDto(m models.CommentLike, u UserDto) CommentLikeDto {
	return CommentLikeDto{
		ID:        m.ID,
		User:      u,
		CommentID: m.CommentID,
		Result:    m.Result,
	}
}
