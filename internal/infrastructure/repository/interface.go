package repository

import "github.com/bosamatheus/go-etl-pipeline/internal/entity"

// Repository interface.
type Repository interface {
	Reader
	Writer
}

// Reader interface.
type Reader interface {
	List() ([]*entity.Review, error)
}

// Writer interface.
type Writer interface {
	Publish(e *entity.Review) (entity.ID, error)
}
