package models

import "time"

type Task struct {
	ID         int64
	UserID     int64
	Title      string
	Body       string
	Expiration time.Time
	CreatedAt  time.Time
}
