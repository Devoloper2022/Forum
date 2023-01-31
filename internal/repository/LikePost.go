package repository

import (
	"fmt"
	"forum/internal/models"
)

type LikePost interface {
	CreatePostLike(modele models.PostLike) error
	UpdatePostLike(modele models.PostLike) error
	UpdatePostTable(id int64, like, dislike int64) error
	GetPostLike(postID, userID int64) (models.PostLike, error)
	DeletePostLike(id int64) error
}

func (r *Database) CreatePostLike(modele models.PostLike) error {
	query := ("INSERT INTO postLike (UserID,PostID,Like,Dislike) VALUES (?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Exec(modele.UserID, modele.PostID, modele.Like, modele.DisLike)

	if err != nil {
		return err
	}

	return nil
}

func (r *Database) UpdatePostLike(modele models.PostLike) error {
	query := ("UPDATE postLike SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Exec(modele.Like, modele.DisLike, modele.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *Database) UpdatePostTable(id int64, like, dislike int64) error {
	query := ("UPDATE posts SET Like=?,Dislike=? WHERE  ID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Exec(like, dislike, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *Database) GetPostLike(postID, userID int64) (models.PostLike, error) {
	query := ("SELECT * FROM postLike WHERE postID = ? AND userID=?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.PostLike{}, fmt.Errorf("repository : Get Post Like 1: %w", err)
	}
	defer st.Close()

	row := st.QueryRow(postID, userID)

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
