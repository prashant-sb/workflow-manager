package main

import (
	"context"
	"os"
	"time"

	//wf "github.com/prashantsb/workflow-manager/pkg/workflow"
	"github.com/prashantsb/workflow-manager/pkg/dag"
	"github.com/prashantsb/workflow-manager/pkg/dagdef"
	"github.com/prashantsb/workflow-manager/pkg/parser"
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
	if err := prs.Validate(dagstr); err != nil {
		panic(err)
	}

	cdag, err := prs.Parse(dagstr)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(persistFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dmgr := dag.NewDagManager(cdag, f)
	dmgr.Print(ctx)
}
