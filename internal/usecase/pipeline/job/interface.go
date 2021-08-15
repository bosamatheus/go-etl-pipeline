package job

// Job defines a pipeline job.
type Job interface {
	Launch() error
}
