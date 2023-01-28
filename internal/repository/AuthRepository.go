package repository

import (
	"fmt"
	"time"
)

type Autorization interface {
	SaveToken(userID int64, sessionToken string, time time.Time) error
	DeleteToken(token string) error
	DeleteTokenWhenExpireTime() error
}

func (r *Database) SaveToken(userID int64, sessionToken string, time time.Time) error {
	query := ("INSERT INTO sessions (Token,Expiry,UserID) VALUES (?,?,?)")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : create post : %w", err)
	}

	_, err = st.Exec(sessionToken, time, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Database) DeleteToken(token string) error {
	query := ("DELETE FROM session WHERE Token=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Token  checker 1: %w", err)
	}

	_, err = st.Exec(token)

	if err != nil {
		return err
	}
	return nil
}

func (r *Database) DeleteTokenWhenExpireTime() error {
	query := ("DELETE FROM session WHERE ExpireTime <= ?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Token  checker 1: %w", err)
	}

	_, err = st.Exec(time.Now())

	if err != nil {
		return err
	}
	return nil
}
