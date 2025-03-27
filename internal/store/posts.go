package store

import (
	"context"
	"database/sql"
)

type PostStrore struct {
	db *sql.DB
}

func (s *PostStrore) Create(ctx context.Context) error{
	return nil
}