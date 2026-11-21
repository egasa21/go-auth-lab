package user

import (
	"context"
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Create(ctx context.Context, u User) error {
	query := `
			INSERT INTO users (username, password_hash) 
			VALUES ($1, $2)
			RETURNING id;`

	return s.db.QueryRowContext(
		ctx,
		query,
		u.Username, u.PasswordHash,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
}
