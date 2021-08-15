package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/bosamatheus/go-etl-pipeline/internal/entity"
	"github.com/bosamatheus/go-etl-pipeline/internal/util"
)

// ReviewPostgres is a PostgreSQL repository for reviews.
type ReviewPostgres struct {
	db *sql.DB
}

// NewReviewPostgres creates a new PostgreSQL repository.
func NewReviewPostgres(DB *sql.DB) Repository {
	return &ReviewPostgres{
		db: DB,
	}
}

// Publish publishes a review into the database.
func (r *ReviewPostgres) Publish(e *entity.Review) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO reviews (
			id, 
			, review_id
			, order_id
			, score
			, comment_title
			, comment_message
			, created_at
			, answered_at
			, published_at
		) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Reference.ReviewID,
		e.Reference.OrderID,
		e.Score,
		e.Comment.Title,
		e.Comment.Message,
		e.CreatedAt,
		e.AnsweredAt,
		util.FormatDateISO(time.Now()),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// List lists reviews from the database.
func (r *ReviewPostgres) List() ([]*entity.Review, error) {
	stmt, err := r.db.Prepare(`SELECT id FROM reviews`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var IDs []entity.ID
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ID entity.ID
		err = rows.Scan(&ID)
		if err != nil {
			return nil, err
		}
		IDs = append(IDs, ID)
	}
	if len(IDs) == 0 {
		return nil, fmt.Errorf("not found")
	}
	var reviews []*entity.Review
	for _, ID := range IDs {
		review, err := r.getReview(ID, r.db)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}
	return reviews, nil
}

// getReview gets a single review from the database.
func (r *ReviewPostgres) getReview(ID entity.ID, DB *sql.DB) (*entity.Review, error) {
	stmt, err := DB.Prepare(`
		SELECT
			id, 
			, review_id
			, order_id
			, score
			, comment_title
			, comment_message
			, created_at
			, answered_at
			, published_at
		FROM reviews
		WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	var review entity.Review
	rows, err := stmt.Query(ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&review.ID,
			&review.Reference.ReviewID,
			&review.Reference.OrderID,
			&review.Score,
			&review.Comment.Title,
			&review.Comment.Message,
			&review.CreatedAt,
			&review.AnsweredAt,
			&review.PublishedAt)
		if err != nil {
			continue
		}
	}
	return &review, nil
}
