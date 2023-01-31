package dto

import (
	"forum/internal/models"
)

type PostDto struct {
	ID        int64             `json:"ID"`
	Title     string            `json:"title"`
	Text      string            `json:"text"`
	Date      string            `json:"data"`
	User      UserDto           `json:"userId"`
	Like      int64             `json:"like"`
	Dislike   int64             `json:"dislike"`
	Likes     models.PostLike   `json:"likes"`
	Categorys []models.Category `json:"categorys"`
}

func GetPostDto(post models.Post, user UserDto, likes models.PostLike, listCat []models.Category) PostDto {
	return PostDto{
		ID:        post.ID,
		Title:     post.Title,
		Text:      post.Text,
		Date:      post.Date,
		Like:      post.Like,
		Dislike:   post.Dislike,
		Likes:     likes,
		Categorys: listCat,
	}
}

func (dto *PostDto) GetPostModel() models.Post {
	// date, _ := time.Parse("d MMM yyyy HH:mm:ss", dto.Date)
	return models.Post{
		ID:      dto.ID,
		Title:   dto.Title,
		Text:    dto.Text,
		Date:    dto.Date,
		Like:    dto.Like,
		Dislike: dto.Dislike,
		UserID:  dto.ID,
	}
}
