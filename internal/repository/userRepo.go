package repository

import (
	"fmt"
	"forum/internal/models"
)

type User interface {
	CreateUser(user models.User) error
	GetUser(userId int64) (models.User, error)
	// UpdateUser(user models.User) error
	// DeleteUser(userID int64) error
}

// main functions
func (r *Database) CreateUser(user models.User) error {
	query := ("INSERT INTO users (Username,Email,Password) VALUES (?,?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("repository : create User : %w", err)
	}
	defer st.Close()

	_, err = st.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *Database) GetUser(userId int64) (models.User, error) {
	query := ("SELECT * FROM users WHERE ID = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.User{}, fmt.Errorf("repository : Get Post  checker 1: %w", err)
	}
	defer st.Close()

	row, err := st.Query(userId)
	defer row.Close()
	var user models.User
	if err = row.Scan(&user.ID, &user.Username); err != nil {
		return models.User{}, err
	}

	return user, nil
}
