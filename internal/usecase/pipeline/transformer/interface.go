package transformer

// Transformer defines a pipeline transformer.
type Transformer interface {
	Transform(data []string) ([]string, error)
}
