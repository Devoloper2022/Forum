package repository

import (
	"fmt"
	"forum/internal/models"
)

type Like interface {
	//Comments Likes
	CreateCommentLike(modele models.CommentLike) (models.CommentLike, error)
	UpdateCommentLike(modele models.CommentLike) (models.CommentLike, error)
	UpdateCommentL(id int64, like, dislike int64) error
	GetCommentLike(id int64) (models.CommentLike, error)
	DeleteCommentLike(id int64) error
	//post Likes
	CreatePostLike(modele models.PostLike) (models.PostLike, error)
	UpdatePostLike(modele models.PostLike) (models.PostLike, error)
	UpdatePostL(id int64, like, dislike int64) (models.Post, error)
	GetPostLike(id int64) (models.PostLike, error)
	DeletePostLike(id int64) error
}

// Comment Likes
func (r *Database) CreateCommentLike(modele models.CommentLike) (models.CommentLike, error) {
	query := ("INSERT INTO commentLike (UserID,CommentID,Like,Dislike) VALUES (?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.CommentLike{}, fmt.Errorf("repository : create Comment Like : %w", err)
	}

	row := st.QueryRow(modele.UserID, modele.CommentID, modele.Like, modele.DisLike)
	var newModele models.CommentLike

	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.CommentID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.CommentLike{}, err
	}

	return newModele, nil
}
func (r *Database) UpdateCommentLike(modele models.CommentLike) (models.CommentLike, error) {
	query := ("UPDATE commentLike SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.CommentLike{}, fmt.Errorf("repository : UPDATE Comment Like : %w", err)
	}

	row := st.QueryRow(modele.Like, modele.DisLike, modele.ID)
	var newModele models.CommentLike

	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.CommentID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.CommentLike{}, err
	}

	return newModele, nil
}
func (r *Database) UpdateCommentL(id int64, like, dislike int64) error {
	query := ("UPDATE comments SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Update Comment L : %w", err)
	}

	_, err = st.Exec(like, dislike, id)

	if err != nil {
		return err
	}

	return nil
}
func (r *Database) GetCommentLike(id int64) (models.CommentLike, error) {
	query := ("SELECT * FROM commentLike WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.CommentLike{}, fmt.Errorf("repository : Get Comment Like 1: %w", err)
	}
	defer st.Close()

	row := st.QueryRow(id)

	var newModele models.CommentLike
	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.CommentID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.CommentLike{}, err
	}

	return newModele, nil
}
func (r *Database) DeleteCommentLike(id int64) error {
	query := ("DELETE FROM commentLike WHERE ID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Comment Like 1: %w", err)
	}

	_, err = st.Exec(id)

	if err != nil {
		return err
	}
	return nil
}

//Post Likes

func (r *Database) CreatePostLike(modele models.PostLike) (models.PostLike, error) {
	query := ("INSERT INTO postLike (UserID,PostID,Like,Dislike) VALUES (?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.PostLike{}, fmt.Errorf("repository : create Post Like : %w", err)
	}

	row := st.QueryRow(modele.UserID, modele.PostID, modele.Like, modele.DisLike)
	var newModele models.PostLike

	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.PostID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.PostLike{}, err
	}

	return newModele, nil
}
func (r *Database) UpdatePostLike(modele models.PostLike) (models.PostLike, error) {
	query := ("UPDATE postLike SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.PostLike{}, fmt.Errorf("repository : Update Post Like : %w", err)
	}

	row := st.QueryRow(modele.Like, modele.DisLike, modele.ID)
	var newModele models.PostLike

	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.PostID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.PostLike{}, err
	}

	return newModele, nil
}
func (r *Database) UpdatePostL(id int64, like, dislike int64) (models.Post, error) {
	query := ("UPDATE posts SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.Post{}, fmt.Errorf("repository : Update Post L : %w", err)
	}

	row := st.QueryRow(like, dislike, id)
	var newModele models.Post

	if err = row.Scan(&newModele.ID, &newModele.Title, &newModele.Text, &newModele.Date, &newModele.Like, &newModele.Dislike, &newModele.UserID); err != nil {
		return models.Post{}, err
	}

	return newModele, nil
}
func (r *Database) GetPostLike(id int64) (models.PostLike, error) {
	query := ("SELECT * FROM postLike WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.PostLike{}, fmt.Errorf("repository : Get Post Like 1: %w", err)
	}
	defer st.Close()

	row := st.QueryRow(id)

	var newModele models.PostLike
	if err = row.Scan(&newModele.ID, &newModele.UserID, &newModele.PostID, &newModele.Like, &newModele.DisLike); err != nil {
		return models.PostLike{}, err
	}

	return newModele, nil
}
func (r *Database) DeletePostLike(id int64) error {
	query := ("DELETE FROM postLike WHERE ID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Post Like 1: %w", err)
	}

	_, err = st.Exec(id)

	if err != nil {
		return err
	}
	return nil
}
