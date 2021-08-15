package publisher

// Publisher defines a pipeline publisher.
type Publisher interface {
	Publish() error
}
