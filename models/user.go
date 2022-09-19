package models

type User struct {
	ID            int64
	TelegramID    int64
	Nick          string
	CurrentStatus Status
	Tasks         []Task
}

type Status string

const (
	StartsStatus Status = "start"
	SetTask      Status = "set"
)
