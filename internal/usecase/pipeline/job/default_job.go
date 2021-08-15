package job

import (
	p "github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/publisher"
	t "github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/task"
)

// DefaultJob is a default pipeline job expecting a task and an optional publisher.
type DefaultJob struct {
	task    t.Task
	publish p.Publisher
}

// NewDefaultJob creates a new default job.
func NewDefaultJob(t t.Task, p p.Publisher) Job {
	return &DefaultJob{
		task:    t,
		publish: p,
	}
}

// Launch launches the default job.
func (j *DefaultJob) Launch() error {
	err := j.task.Run()
	if j.publish != nil {
		err = j.publish.Publish()
	}
	return err
}
