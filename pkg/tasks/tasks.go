package tasks

import (
	"context"
	"fmt"

	"github.com/prashantsb/workflow-manager/pkg/preserver"
)

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
	WaitForCompletion bool `yaml:"waitForCompletion"`
	AllowParallel     bool `yaml:"allowParallel"`
	RetryCount        int  `yaml:"retryCount"`
	MaxRetry          int  `yaml:"maxRetry"`
}

type Config struct {
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Version   string     `yaml:"version"`
	Workflows []Workflow `yaml:"workflows"`
}

type Workflow struct {
	WorkflowID string `yaml:"workflowId"`
	Tasks      []Task `yaml:"tasks"`
}

// Task is a concrete implementation of Vertex.
type Task struct {
	TaskID           string     `yaml:"taskId"`
	TaskConfig       TaskParams `yaml:"taskConfig"`
	Fn               RunnerFunc
	parentWorkflowID string
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
	return t.TaskID
}

func (t *Task) Run(ctx context.Context) error {
	t.TaskConfig.RetryCount++
	return t.Fn(ctx)
}

func (t *Task) WorkflowID() string {
	return t.parentWorkflowID

}

func NewTaskConfigFromYaml(in string) (*Config, error) {
	cfg := &Config{}
	fp, err := preserver.NewConfigHandler[Config](in)
	if err != nil {
		return nil, err
	}

	if err := fp.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) GetTasksMap(ver string) (map[string]*Task, error) {
	tasksMap := make(map[string]*Task)

	for _, version := range cfg.Versions {
		if version.Version == ver {
			for _, wf := range version.Workflows {
				for _, task := range wf.Tasks {
					// Set default params if not provided
					if task.TaskConfig == (TaskParams{}) {
						task.TaskConfig = *DefaultTaskParams()
					}
					task.parentWorkflowID = wf.WorkflowID
					Fn, err := getFunctionByID(task.TaskID)
					if err != nil {
						return nil, err
					}
					task.Fn = Fn
					tasksMap[task.TaskID] = &task
				}
			}
		}
	}

	if len(tasksMap) == 0 {
		return nil, fmt.Errorf("no tasks found for version %s", ver)
	}

	return tasksMap, nil
}
