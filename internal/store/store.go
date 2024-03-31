package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type Store struct {
	config     *Config
	Connection *pgx.Conn
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open(ctx context.Context) error {
	url := fmt.Sprintf(
		"host=%s dbname=%s sslmode=disable user=%s password=%s",
		s.config.DatabaseHost,
		s.config.DatabaseName,
		s.config.DatabaseUser,
		s.config.DatabasePassword,
	)

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}

	s.Connection = conn

	return nil
}

func (s *Store) Close(ctx context.Context) {
	defer s.Connection.Close(ctx)
}
