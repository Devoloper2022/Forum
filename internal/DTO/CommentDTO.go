package dto

import "forum/internal/models"

type CommentDto struct {
	ID      int64          `json:"id"`
	Text    string         `json:"text"`
	Date    string         `json:"date"`
	Like    int64          `json:"like"`
	Dislike int64          `json:"dislike"`
	User    UserDto        `json:"userId"`
	PostID  int64          `json:"postId"`
	Likes   CommentLikeDto `json:"likes"`
}

func (dto *CommentDto) GetCommentModel() models.Comment {
	return models.Comment{
		ID:      dto.ID,
		Text:    dto.Text,
		Date:    dto.Date,
		Like:    dto.Like,
		Dislike: dto.Dislike,
		UserID:  dto.User.ID,
		PostID:  dto.PostID,
	}
}

func GetCommentDto(m models.Comment, u UserDto, l CommentLikeDto) CommentDto {
	return CommentDto{
		ID:      m.ID,
		Text:    m.Text,
		Date:    m.Date,
		Like:    m.Like,
		Dislike: m.Dislike,
		User:    u,
		PostID:  m.PostID,
		Likes:   l,
	}
}
