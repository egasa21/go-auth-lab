package user

import (
	"context"
	"database/sql"
	"log"
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
			RETURNING id, created_at, updated_at;`

	err := s.db.QueryRowContext(
		ctx,
		query,
		u.Username, u.PasswordHash,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		log.Printf("Create user error: %v", err)
		return err
	}

	return nil
}
