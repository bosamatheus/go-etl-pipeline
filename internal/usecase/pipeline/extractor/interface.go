package extractor

// Extractor defines a pipeline extractor.
type Extractor interface {
	Extract() ([]string, error)
}
