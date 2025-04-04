package repository

import (
	"context"
	"database/sql"
	"errors"
)

type URLRepository interface {
    Save(ctx context.Context, hash string, url string, ip string) error
    Find(ctx context.Context, hash string) (string, error)
}
type PGURLRepository struct {
	db *sql.DB
}

func NewPGURLRepository(db *sql.DB) (*PGURLRepository, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}
	return &PGURLRepository{db: db}, nil
}

func (r *PGURLRepository) Save(ctx context.Context, shortURL string, originalURL string, ip string) error {
	if r.db == nil {
        return errors.New("database connection is not initialized")
    }

	query := `INSERT INTO short_urls (short_code, long_url, creator_ip) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, shortURL, originalURL, ip)
	if err != nil {
		return err
	}
	return nil
}

func (r *PGURLRepository) Find(ctx context.Context, shortURL string) (string, error) {
    if r.db == nil {
        return "", errors.New("database connection is not initialized")
    }

    query := `SELECT long_url FROM short_urls WHERE short_code = $1`
    var originalURL string
    err := r.db.QueryRowContext(ctx, query, shortURL).Scan(&originalURL)
    if err == sql.ErrNoRows {
        return "", errors.New("short URL not found")
    } else if err != nil {
        return "", errors.New("dont know the error")
    }
    return originalURL, nil
}