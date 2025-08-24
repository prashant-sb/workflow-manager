package dagdef

import (
	"context"
	"os"

	"github.com/prashantsb/workflow-manager/pkg/tasks"
)

// common DAG definition configuration if any

var TaskRegistry = map[string]*tasks.Task{
	"Net": &tasks.Task{
		TaskId: "Net", WorkflowId: "", Fn: func(ctx context.Context) error {
			// Simulate network setup
			return nil
		}},
	"VM": &tasks.Task{TaskId: "VM", WorkflowId: "", Fn: func(ctx context.Context) error {
		// Simulate VM setup
		return nil
	}},
	"App": &tasks.Task{TaskId: "App", WorkflowId: "", Fn: func(ctx context.Context) error {
		// Simulate App setup
		return nil
	}},
}

func GetDAGFromDefination(in string) (string, error) {
	data, err := os.ReadFile(in)
	if err != nil {
		return in, err
	}

	return string(data), nil
}
