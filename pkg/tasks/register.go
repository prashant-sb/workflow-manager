package tasks

import (
	"context"
	"fmt"

	"github.com/prashantsb/workflow-manager/pkg/tasks/runners"
)

// RunnerFunc is the signature for all task runners.
type RunnerFunc func(ctx context.Context) error

var runnerRegistry = map[string]RunnerFunc{
	"Net":     runners.NetRunner,
	"Stor":    runners.StorRunner,
	"VM":      runners.VMRunner,
	"InfraR":  runners.InfraRRunner,
	"PreChk":  runners.PreChkRunner,
	"ProvDB":  runners.ProvDBRunner,
	"Schema":  runners.SchemaRunner,
	"DBR":     runners.DBRRunner,
	"ProvMon": runners.ProvMonRunner,
	"Hook":    runners.HookRunner,
	"MonR":    runners.MonRRunner,
	"E2E":     runners.E2ERunner,
	"Deploy":  runners.DeployRunner,
	"Smoke":   runners.SmokeRunner,
	"AppR":    runners.AppRRunner,
	"SecScan": runners.SecScanRunner,
	"Commit":  runners.CommitRunner,
	"Roll":    runners.RollRunner,
	"Notify":  runners.NotifyRunner,
	"Sink":    runners.SinkRunner,
}

func getFunctionByID(id string) (RunnerFunc, error) {
	if fn, exists := runnerRegistry[id]; exists {
		return fn, nil
	}

	return nil, fmt.Errorf("no function registered for task ID: %s", id)
}
