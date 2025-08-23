package dag

import (
	"context"
	"io"
)

// Node represents a unit of execution in the DAG.
type Node interface {
	ID() string
	Run(ctx context.Context) error
}

// DAG represents a directed acyclic graph of nodes.
type DAG interface {
	addNode(node Node)
	addEdge(from, to string) error
	nodes() []Node
	dependencies(nodeID string) []string
	isAcyclic() bool
}

// DAGManager orchestrates execution of a DAG.
type DAGManager interface {
	// Load the DAG definition
	Load(dag DAG) error

	// Execute the DAG (with parallelism, joins, waits)
	Execute(ctx context.Context) error

	// Commit section ensures exclusive access.
	Commit(ctx context.Context, iorw io.ReadWriter) error

	// Print the DAG structure for debugging.
	Print(ctx context.Context)
}
