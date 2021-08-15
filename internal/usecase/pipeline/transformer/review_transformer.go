package transformer

// ReviewTransformer is a pipeline transformer for reviews.
type ReviewTransformer struct{}

// NewReviewTransformer creates a new review transformer.
func NewReviewTransformer() Transformer {
	return &ReviewTransformer{}
}

// Transform transforms a review.
func (t *ReviewTransformer) Transform(data []string) ([]string, error) {
	// TODO: implement the reviews transformer.
	_ = data
	return nil, nil
}
