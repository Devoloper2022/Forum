package repository

import (
	"fmt"
	"time"
)

type Autorization interface {
	SaveToken(userID int64, sessionToken string, time time.Time) error
	GetToken(id int64) (string, time.Time, error)
	UpdateToken(tokenName, newToken string, expireTime time.Time) error
	DeleteToken(token string) error
	DeleteTokenWhenExpireTime() error

	GetTokens(id int64) error
	DeleteTokenByUserID(userID int64) error
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

func (r *Database) UpdateToken(tokenName, newToken string, expireTime time.Time) error {
	query := ("UPDATE sessions SET Token=?, Expiry=? WHERE Token=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : UPDATE Comment Like : %w", err)
	}

	_, err = st.Exec(newToken, expireTime, tokenName)
	if err != nil {
		return err
	}

	return nil
}

func (r *Database) GetToken(id int64) (string, time.Time, error) {
	query := ("SELECT Token,Expiry FROM sessions WHERE UserID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return "", time.Time{}, err
	}

	row := st.QueryRow(id)

	var token string
	var expireTime time.Time
	if err = row.Scan(&token, &expireTime); err != nil {
		return "", time.Time{}, err
	}

	return token, expireTime, nil
}

func (r *Database) GetTokens(id int64) error {
	query := ("SELECT Token,Expiry FROM sessions WHERE UserID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return err
	}

	_, err = st.Query(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Database) DeleteToken(token string) error {
	query := ("DELETE FROM sessions WHERE Token=?")
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

func (r *Database) DeleteTokenByUserID(userID int64) error {
	query := ("DELETE FROM sessions WHERE UserID=?")
	st, err := r.db.Prepare(query)
	defer st.Close()

	if err != nil {
		return fmt.Errorf("repository : Delete Token  checker 1: %w", err)
	}

	_, err = st.Exec(userID)

	if err != nil {
		return err
	}
	return nil
}

func (r *Database) DeleteTokenWhenExpireTime() error {
	query := ("DELETE FROM sessions WHERE Expiry <= ?")
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
