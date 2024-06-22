package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connStr string) (*Storage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS urls(
			id SERIAL PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
	`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) (int64, error) {
	var id int64
	err := s.db.QueryRow("INSERT INTO urls(url, alias) VALUES($1, $2) RETURNING id", urlToSave, alias).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
