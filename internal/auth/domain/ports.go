package domain

import "errors"

// ErrUserNotFound use for different the db's problem at user no exist.
var ErrUserNotFound = errors.New("user not found")

// UserRepository define the port.
// FindByUsername for login.
// Create for register.
type UserRepository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}
