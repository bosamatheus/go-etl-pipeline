package loader

// Loader defines a pipeline loader.
type Loader interface {
	Load(data []string) error
}
