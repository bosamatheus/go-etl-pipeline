package task

// Task defines a pipeline task.
type Task interface {
	Run() error
}
