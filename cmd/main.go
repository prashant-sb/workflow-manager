package main

import (
	"context"
	"time"

	"github.com/prashantsb/workflow-manager/pkg/dag"
	"github.com/prashantsb/workflow-manager/pkg/dagdef"
	"github.com/prashantsb/workflow-manager/pkg/parser"
	"github.com/prashantsb/workflow-manager/pkg/preserver"
)

const (
	persistFile       = "pkg/dagdef/workflow-state.json"
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
	cdag, err := prs.Parse()
	if err != nil {
		panic(err)
	}

	fp, err := preserver.NewConfigHandler[dag.PresistAttrib](persistFile)
	if err != nil {
		panic(err)
	}

	dmgr := dag.NewDagManager(cdag, fp)
	_ = dmgr.Load(cdag)
	dmgr.Print(ctx)
	_ = dmgr.Commit(ctx)
}
