package models

import "time"

type Session struct {
	ID       int64
	username string
	Token    string
	expiry   time.Time
}

func (s Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}
