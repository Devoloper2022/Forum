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
	GetUserByToken(token string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}

// main functions
func (r *Database) CreateUser(user models.User) error {
	query := ("INSERT INTO users (Username,Email,Password) VALUES (?,?,?)")
	st, err := r.db.Prepare(query)
	if err != nil {
		return err
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
		return models.User{}, err
	}
	defer st.Close()

	row := st.QueryRow(userId)

	var user models.User
	if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Database) GetUserByUsername(username string) (models.User, error) {
	query := ("SELECT * FROM users WHERE Username = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	defer st.Close()

	row := st.QueryRow(username)

	var user models.User
	if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Database) GetUserByEmail(email string) (models.User, error) {
	query := ("SELECT * FROM users WHERE Email = ?")
	st, err := r.db.Prepare(query)
	if err != nil {
		return models.User{}, err
	}
	defer st.Close()

	row := st.QueryRow(email)

	var user models.User
	if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Database) GetUserByToken(token string) (models.User, error) {
	query := ("SELECT users.ID, users.Username ,users.Email  , users.Password FROM users  INNER JOIN sessions ON users.ID = sessions.UserID WHERE sessions.Token = ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	fmt.Println()
	if err != nil {
		return models.User{}, err
	}

	row := st.QueryRow(token)
	var user models.User
	if err = row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		return models.User{}, err
	}
	return user, nil
}
