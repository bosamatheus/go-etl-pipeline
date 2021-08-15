package entity

import "time"

// Review represents a review entity.
type Review struct {
	ID          ID
	Reference   reviewReference
	Score       int
	Comment     reviewComment
	CreatedAt   time.Time
	AnsweredAt  time.Time
	PublishedAt time.Time
}

// reviewReference represents a review reference.
type reviewReference struct {
	ReviewID string
	OrderID  int64
}

// reviewComment represents a review comment.
type reviewComment struct {
	Title   string
	Message string
}
