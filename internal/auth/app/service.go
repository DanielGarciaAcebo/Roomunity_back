package app

import (
	"errors"

	"roomunity_back/internal/auth/domain"
)

type AuthService struct {
	users domain.UserRepository
}

func NewAuthService(users domain.UserRepository) *AuthService {
	return &AuthService{users: users}
}

func (s *AuthService) Login(username, password string) (*domain.User, error) {
	u, err := s.users.FindByUsername(username)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, nil
		}
		return nil, err // real error
	}

	if u.PasswordHash != password {
		return nil, nil
	}

	return u, nil
}
