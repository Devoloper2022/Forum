package repository

import (
	"fmt"
	"forum/internal/models"
)

type Post interface {
	CreatePost(post models.Post) (int64, error)
	GetAllPosts() ([]models.Post, error)
	GetPost(postId int64) (models.Post, error)
	GetPostByUserID(userID int64) ([]models.Post, error)
	UpdatePost(post models.Post) error
	GetPostsByMostLikes() ([]models.Post, error)
	GetPostsByLeastLikes() ([]models.Post, error)
	GetPostByCategory(category int64) ([]models.Post, error)
	GetPostsByDate() ([]models.Post, error)
}

// main functions
func (r *Database) CreatePost(post models.Post) (int64, error) {
	query := ("INSERT INTO posts (Title,Text,Like,Dislike,UserId) VALUES (?,?,?,?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer st.Close()

	res, err := st.Exec(post.Title, post.Text, post.Like, post.Dislike, post.UserID)

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(id), nil
} // done

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
} // not checked

func (r *Database) GetAllPosts() ([]models.Post, error) {
	query := ("SELECT *  from posts")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
} // done

func (r *Database) GetPost(postId int64) (models.Post, error) {
	query := ("SELECT * FROM posts WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.Post{}, err
	}
	defer st.Close()

	row := st.QueryRow(postId)
	var post models.Post
	// var date string
	if err = row.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
		return models.Post{}, err
	}
	// newDate, _ := time.Parse("d MMM yyyy HH:mm:ss", date)
	// post.Date = newDate
	return post, nil
}

func (r *Database) GetPostByUserID(userID int64) ([]models.Post, error) {
	query := ("SELECT * FROM posts WHERE UserID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(userID)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err != rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID) {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *Database) GetPostsByMostLikes() ([]models.Post, error) {
	query := ("SELECT *  from posts WHERE Like > 0 ORDER BY Like DESC,Dislike ASC")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *Database) GetPostsByLeastLikes() ([]models.Post, error) {
	query := ("SELECT *  from posts WHERE Dislike > 0 ORDER BY Dislike DESC,Like ASC")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *Database) GetPostsByDate() ([]models.Post, error) {
	query := ("SELECT *  from posts ORDER BY Date  ASC")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *Database) GetPostByCategory(categoryID int64) ([]models.Post, error) {
	query := ("SELECT posts.ID, posts.Title, posts.Text, posts.Date, posts.Like, posts.Dislike, posts.UserID FROM posts INNER JOIN categoriesPost ON posts.ID = categoriesPost.PostID WHERE categoriesPost.categoryID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()

	rows, err := st.Query(categoryID)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Text, &post.Date, &post.Like, &post.Dislike, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
