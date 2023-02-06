package repository

import (
	"forum/internal/models"
)

type Comment interface {
	CreateComment(comment models.Comment) error
	GetComment(commentID int64) (models.Comment, error)
	UpdateComment(comment models.Comment) error
	DeleteComment(commentID int64) error
	GetAllCommentByPostID(postId int64) ([]models.Comment, error)
	GetAllCommentByUserID(userId int64) ([]models.Comment, error)
}

func (r *Database) CreateComment(comment models.Comment) error {
	query := ("INSERT INTO comments (Text,Date,Like,Dislike,UserID,PostID) VALUES (?,?,?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Exec(comment.Text, comment.Date, comment.Like, comment.Dislike, comment.UserID, comment.PostID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Database) GetComment(commentID int64) (models.Comment, error) {
	query := ("SELECT * FROM comments WHERE ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.Comment{}, err
	}

	row := st.QueryRow(commentID)
	var comment models.Comment
	if err = row.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.Like, &comment.Dislike, &comment.PostID, &comment.UserID); err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (r *Database) UpdateComment(comment models.Comment) error {
	query := "UPDATE comments SET  Text =?,Like=?,Dislike=? WHERE  ID = ?"
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}
	_, err = st.Query(comment.Text, comment.Like, comment.Dislike, comment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Database) DeleteComment(commentID int64) error {
	query := ("DELETE FROM comments WHERE ID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
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
		return nil, err
	}
	defer st.Close()
	rows, err := st.Query(postId)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err != rows.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.Like, &comment.Dislike, &comment.UserID, &comment.PostID) {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (r *Database) GetAllCommentByUserID(userId int64) ([]models.Comment, error) {
	query := ("SELECT * FROM comments WHERE UserID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer st.Close()
	rows, err := st.Query(userId)
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		if err != rows.Scan(&comment.ID, &comment.Text, &comment.Date, &comment.UserID, &comment.PostID) {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
