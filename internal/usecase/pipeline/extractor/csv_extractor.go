package extractor

// CSVExtractor extracts data from a CSV file.
type CSVExtractor struct {
	filename string
}

// NewCSVExtractor creates a new CSV extractor.
func NewCSVExtractor(filename string) Extractor {
	return &CSVExtractor{
		filename: filename,
	}
}

// Extract extracts data from a CSV file.
func (e *CSVExtractor) Extract() ([]string, error) {
	// TODO: implement the CSV extractor.
	return nil, nil
}
