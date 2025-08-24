package tasks

import "context"

const (
	// infinite retries
	DefaultMaxRetry = -1
)

// Vertex represents a unit of execution in the DAG.
type Vertex interface {
	ID() string
	Run(ctx context.Context) error
	WorkflowID() string
}

type TaskParams struct {
	WaitForCompletion bool
	AllowParallel     bool
	RetryCount        uint
	MaxRetry          int
}

// Task is a concrete implementation of Vertex.
type Task struct {
	TaskId     string
	WorkflowId string
	TaskConfig *TaskParams
	Fn         func(ctx context.Context) error
}

func NewTask(id, grp string, fn func(ctx context.Context) error) Vertex {
	return &Task{TaskId: id, Fn: fn, WorkflowId: grp}
}

func DefaultTaskParams() *TaskParams {
	return &TaskParams{
		WaitForCompletion: false,
		AllowParallel:     false,
		RetryCount:        0,
		MaxRetry:          DefaultMaxRetry,
	}
}

func (t *Task) ID() string {
	return t.TaskId
}

func (t *Task) Run(ctx context.Context) error {
	t.TaskConfig.RetryCount++
	return t.Fn(ctx)
}

func (t *Task) WorkflowID() string {
	return t.WorkflowId
}

func (t *Task) WithParams(params *TaskParams) *Task {
	t.TaskConfig = params
	return t
}
