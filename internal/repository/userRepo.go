package repository

import (
	m "forum/internal/models"
	"log"
)

func (r *Repo) CreateUser(user m.User) error {
	records := `INSERT INTO users(username,email,password) VALUES(?,?,?)`
	query, err := r.DB.Prepare(records)
	if err != nil {
		log.Println("Repo ==> CreateUser 1" + err.Error())
		return err
	}

	_, err = query.Exec(user.Username, user.Email, user.Password)

	if err != nil {
		log.Println("Repo ==> CreateUser 2" + err.Error())
		return err
	}

	return nil
}

func (r *Repo) getUserById(id int64) (m.User, error) {
	findId := r.DB.QueryRow("SELECT * FROM users WHERE  users_id = ?", id)
	user := m.User{}

	err := findId.Scan(&user.ID, user.Email, user.Password, user.Username)
	if err != nil {
		log.Println("Repo ==> GetUserById")
		return user, err
	}

	return user, nil
}

func (r *Repo) UpdateUser(user m.User) (m.User, error) {
	query := "UPDATE users SET  username =?, email=?, password=? WHERE  ID = ?"
	stmt, err := r.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		log.Println("Repo ==> CreateUser 1" + err.Error())
		return user, err
	}

	_, err = stmt.Exec(user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		log.Println("Repo ==> CreateUser 1" + err.Error())
		return user, err
	}

	return user, nil
}

func DeleteUser(id int64) error {
	query := `DELETE FROM users WHERE id = ?`
}
