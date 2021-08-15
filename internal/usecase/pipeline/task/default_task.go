package task

import (
	e "github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/extractor"
	l "github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/loader"
	t "github.com/bosamatheus/go-etl-pipeline/internal/usecase/pipeline/transformer"
)

// DefaultTask defines a default pipeline task expecting an extractor, a transformer, and a loader.
type DefaultTask struct {
	extractor   e.Extractor
	transformer t.Transformer
	loader      l.Loader
}

// NewDefaultTask creates a new default task.
func NewDefaultTask(e e.Extractor, t t.Transformer, l l.Loader) Task {
	return &DefaultTask{
		extractor:   e,
		transformer: t,
		loader:      l,
	}
}

// Run runs the default task.
func (t *DefaultTask) Run() error {
	raw, err := t.extractor.Extract()
	if err != nil {
		return err
	}
	transformed, err := t.transformer.Transform(raw)
	if err != nil {
		return err
	}
	err = t.loader.Load(transformed)
	if err != nil {
		return err
	}
	return nil
}
