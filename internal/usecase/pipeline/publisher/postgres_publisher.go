package publisher

import r "github.com/bosamatheus/go-etl-pipeline/internal/infrastructure/repository"

// PostgresPublisher publishes data to a PostgreSQL database.
type PostgresPublisher struct {
	filename   string
	deleteFile bool
	repo       *r.Repository
}

// NewPostgresPublisher creates a new PostgreSQL publisher.
func NewPostgresPublisher(filename string, deleteFile bool, repo *r.Repository) Publisher {
	return &PostgresPublisher{
		filename:   filename,
		deleteFile: deleteFile,
		repo:       repo,
	}
}

// Publish publishes data to a PostgreSQL database.
func (p *PostgresPublisher) Publish() error {
	// TODO: implement the PostgreSQL publisher.
	return nil
}
