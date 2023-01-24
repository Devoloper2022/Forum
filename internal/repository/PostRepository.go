package repository

import (
	"fmt"
	dto "forum/internal/DTO"
	"forum/internal/models"
	"strconv"
)

type Post interface {
	CreatePost(post models.Post) (int64, error)
	GetAllPosts() ([]models.Post, error)
	GetPost(postId int64) (models.Post, error)
	GetPostByUserID(userID int64) ([]models.Post, error)
	UpdatePost(post models.Post) error
	// GetPostsByMostLikes() ([]models.PostInfo, error)
	// GetPostsByLeastLikes() ([]models.PostInfo, error)
	// GetPostByCategory(category string) ([]models.PostInfo, error)

	CreatePostCategory(postId int64, categories []string) error
	GetUser(userId int64) (dto.UserDto, error)
}

// main functions
func (r *Database) CreatePost(post models.Post) (int64, error) {
	query := ("INSERT INTO posts (Title,Text,Date,Like,Dislike,UserId) VALUES (?,?,?,?,?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("repository : create post : %w", err)
	}
	defer st.Close()

	res, err := st.Exec(post.Title, post.Text, post.Date, post.Like, post.Dislike, post.UserID)

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}

func (r *Database) UpdatePost(post models.Post) error {
	query := "UPDATE posts SET  Title =?, Text=?, WHERE  ID = ?"
	st, err := r.db.Prepare(query)

	defer st.Close()

	if err != nil {
		return fmt.Errorf("\nrepository : Update Post  checker 1\n: %w", err)
	}

	_, err = st.Query(post.Title, post.Text, post.ID)
	if err != nil {
		return fmt.Errorf("\nrepository : Update Post  checker 1\n: %w", err)
	}

	return nil
}

func (r *Database) GetAllPosts() ([]models.Post, error) {
	query := ("SELECT *  FROM posts")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get All Posts  checker 1\n: %w", err)
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get All Posts  checker 2\n: %w", err)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("\nrepository : Get All Posts  checker 3\n: %w", err)
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err != rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID) {
			return nil, fmt.Errorf("\n repository : Get All Posts  checker 4\n: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *Database) GetPost(postId int64) (models.Post, error) {
	query := ("SELECT * FROM posts WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.Post{}, fmt.Errorf("repository : Get Post  checker 1: %w", err)
	}
	defer st.Close()

	row, err := st.Query(postId)
	defer row.Close()
	var post models.Post
	if err = row.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (r *Database) GetPostByUserID(userID int64) ([]models.Post, error) {
	query := ("SELECT * FROM posts WHERE UserID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 1:\n %w", err)
	}
	defer st.Close()

	rows, err := st.Query(userID)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 2\n: %w", err)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 3\n: %w", err)
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err != rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID) {
			return nil, fmt.Errorf("\n repository : Get Post by UserID  checker 4\n: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// helper functions
func (r *Database) CreatePostCategory(postId int64, categories []string) error {
	query := ("INSERT INTO categoriesPost (PostID,CategoryID) VALUES (?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("repository : create PostCategory  checker 1: %w", err)
	}
	defer st.Close()

	for _, cid := range categories {
		num, err := strconv.Atoi(cid)
		if err != nil {
			return fmt.Errorf("repository : create PostCategory  checker 2: %w", err)
		}
		_, err = st.Exec(postId, num)
		if err != nil {
			return fmt.Errorf("repository : create PostCategory checker 3: %w", err)
		}
	}
	return nil
}

func (r *Database) GetUser(userId int64) (dto.UserDto, error) {
	query := ("SELECT * FROM users WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return dto.UserDto{}, fmt.Errorf("repository : Get Post  checker 1: %w", err)
	}
	defer st.Close()

	row, err := st.Query(userId)
	defer row.Close()
	var user dto.UserDto
	if err = row.Scan(&user.ID, &user.Username); err != nil {
		return dto.UserDto{}, err
	}

	return user, nil
}
