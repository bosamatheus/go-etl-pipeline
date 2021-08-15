package loader

// CSVLoader loads data into a CSV file.
type CSVLoader struct {
	filename string
}

// NewCSVLoader creates a new CSV loader.
func NewCSVLoader(filename string) Loader {
	return &CSVLoader{
		filename: filename,
	}
}

// Load loads data into a temporary CSV file.
func (l *CSVLoader) Load(data []string) error {
	// TODO: implement the CSV loader.
	_ = data
	return nil
}
