package main

import (
	"context"
	"os"
	"time"

	//wf "github.com/prashantsb/workflow-manager/pkg/workflow"
	"github.com/prashantsb/workflow-manager/pkg/dag"
	"github.com/prashantsb/workflow-manager/pkg/parser"
)

const (
	persistFile = "workflow-state.json"
)

var dotDAG = `digraph G {
	A -> B;
	`

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()

	prs := parser.NewDOTParser(dotDAG)
	if err := prs.Validate(dotDAG); err != nil {
		panic(err)
	}

	cdag, err := prs.Parse(dotDAG)
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
