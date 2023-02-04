package repository

import (
	"forum/internal/models"
)

type LikeComment interface {
	CreateCommentLike(modele models.CommentLike) error
	UpdateCommentLike(modele models.CommentLike) error
	UpdateCommentTable(id int64, like, dislike int64) error
	GetCommentLike(userID, commentID int64) (models.CommentLike, error)
	DeleteCommentLike(id int64) error
}

func (r *Database) CreateCommentLike(modele models.CommentLike) error {
	query := ("INSERT INTO commentLike (UserID,CommentID,Like,Dislike) VALUES (?,?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Exec(modele.UserID, modele.CommentID, modele.Like, modele.DisLike)

	if err != nil {
		return err
	}

	return nil
}

func (r *Database) UpdateCommentLike(modele models.CommentLike) error {
	query := ("UPDATE commentLike SET Like=?,Dislike=? WHERE  ID = ?")
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

func (r *Database) UpdateCommentTable(id int64, like, dislike int64) error {
	query := ("UPDATE comments SET Like=?,Dislike=? WHERE  ID = ?")
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

func (r *Database) GetCommentLike(userID, commentID int64) (models.CommentLike, error) {
	query := ("SELECT * FROM commentLike WHERE commentID = ? AND userID = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return models.CommentLike{}, err
	}

	row := st.QueryRow(commentID, userID)

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
		return err
	}

	_, err = st.Exec(id)

	if err != nil {
		return err
	}
	return nil
}
