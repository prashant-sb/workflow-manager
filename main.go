package main

import (
	"context"
	"time"

	//wf "github.com/prashantsb/workflow-manager/pkg/workflow"
	"github.com/prashantsb/workflow-manager/pkg/dag"
	"github.com/prashantsb/workflow-manager/pkg/dagdef"
	"github.com/prashantsb/workflow-manager/pkg/parser"
	"github.com/prashantsb/workflow-manager/pkg/preserver"
)

const (
	persistFile       = "workflow-state.json"
	dagDefinationFile = "pkg/dagdef/workflow.dot"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	dagstr, err := dagdef.GetDAGFromDefination(dagDefinationFile)
	if err != nil {
		panic(err)
	}

	prs := parser.NewDOTParser(dagstr)
	cdag, err := prs.Parse(dagdef.TaskRegistry)
	if err != nil {
		panic(err)
	}

	fp := preserver.NewFilePreserver[dag.PresistAttrib](persistFile)
	dmgr := dag.NewDagManager(cdag, fp)

	dmgr.Load(cdag)
	dmgr.Print(ctx)
	dmgr.Commit(ctx)
}
