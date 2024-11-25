package repository

import "database/sql"

type Storage struct {
	db *sql.DB
}

func New(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Close() {
	s.db.Close()
}
