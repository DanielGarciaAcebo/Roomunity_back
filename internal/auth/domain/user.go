package domain

import "time"

type User struct {
	ID           int64
	Email        string
	Username     string
	FirstName    string
	LastName     string
	GroupName    string
	Number       int
	PasswordHash string
	CreatedAt    time.Time
}
