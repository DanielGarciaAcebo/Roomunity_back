package app

// AuthService holds the business logic for authentication.
type AuthService struct{}

// NewAuthService creates a new AuthService instance.
func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(username, password string) error {
	// No real persistence yet, just pretend it's OK.
	return nil
}

// Login checks if the provided credentials match the hardcoded user.
func (s *AuthService) Login(username, password string) bool {
	return username == "Wild" && password == "123abc"
}
