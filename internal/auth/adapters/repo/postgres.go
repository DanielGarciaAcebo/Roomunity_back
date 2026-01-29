package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"roomunity_back/internal/auth/domain"
)

// PostgresUserRepository implement the port UserRepository to used Postgres
type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	return errors.New("not implemented: Create")
}

// Change this sql func for other, need archive for sql funcions
func (r *PostgresUserRepository) FindByUsername(username string) (*domain.User, error) {
	const query = `
SELECT id, email, username, first_name, last_name, group_name,
       number, password_hash, created_at
FROM public.users
WHERE username = $1;
`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	u := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.GroupName,
		&u.Number,
		&u.PasswordHash,
		&u.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}
