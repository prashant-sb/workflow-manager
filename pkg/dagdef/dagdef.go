package dagdef

import (
	"context"
	"os"

	"github.com/prashantsb/workflow-manager/pkg/dag"
)

// common DAG definition configuration if any

var TaskRegistry = map[string]*dag.Task{
	"Net": &dag.Task{
		Id: "Net", SubGraph: "", Fn: func(ctx context.Context) error {
			// Simulate network setup
			return nil
		}},
	"VM": &dag.Task{Id: "VM", SubGraph: "", Fn: func(ctx context.Context) error {
		// Simulate VM setup
		return nil
	}},
	"App": &dag.Task{Id: "App", SubGraph: "", Fn: func(ctx context.Context) error {
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
