package repository

import (
	"fmt"
	"forum/internal/models"
)

type Comment interface {
	CreateComment(comment models.Comment) error
	GetComment(commentID int64) error
	UpdateComment(comment models.Comment) error
	DeleteComment(commentID int64) error
	GetAllCommentByPostID(postId int64) ([]models.Comment, error)
	GetAllCommentByUserID(userId int64) ([]models.Comment, error)
}

func (r *Database) CreateComment(comment models.Comment) error {
	query := ("INSERT INTO comments (Text,Date,UserID,PostID) VALUES (?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : create Comment : %w", err)
	}

	_, err = st.Exec(comment.Text, comment.Date, comment.UserID, comment.PostID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Database) GetComment(commentID int64) error {
	query := ("SELECT * FROM comments WHERE ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Get Post  checker 1: %w", err)
	}

	_, err = st.Exec(commentID)

	if err != nil {
		return err
	}
	return nil
}

func (r *Database) UpdateComment(comment models.Comment) error {
	query := "UPDATE comments SET  Text =? WHERE  ID = ?"
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("\nrepository : Update Post  checker 1\n: %w", err)
	}
	_, err = st.Query(comment.Text, comment.ID)
	if err != nil {
		return fmt.Errorf("\nrepository : Update Post  checker 1\n: %w", err)
	}
	return nil
}

func (r *Database) DeleteComment(commentID int64) error {
	query := ("DELETE FROM comments WHERE ID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Token  checker 1: %w", err)
	}

	_, err = st.Exec(commentID)

	if err != nil {
		return err
	}
	return nil
}

func (r *Database) GetAllCommentByPostID(postId int64) ([]models.Comment, error) {
	query := ("SELECT * FROM comments WHERE PostID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 1:\n %w", err)
	}
	defer st.Close()
	rows, err := st.Query(postId)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 2\n: %w", err)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 3\n: %w", err)
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err != rows.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.UserID, &comment.PostID) {
			return nil, fmt.Errorf("\n repository : Get Post by UserID  checker 4\n: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *Database) GetAllCommentByUserID(userId int64) ([]models.Comment, error) {
	query := ("SELECT * FROM comments WHERE UserID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 1:\n %w", err)
	}
	defer st.Close()
	rows, err := st.Query(userId)
	if err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 2\n: %w", err)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("\nrepository : Get Post by UserID  checker 3\n: %w", err)
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err != rows.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.UserID, &comment.PostID) {
			return nil, fmt.Errorf("\n repository : Get Post by UserID  checker 4\n: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
