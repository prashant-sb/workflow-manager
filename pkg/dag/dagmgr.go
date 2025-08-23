package dag

import (
	"context"
	"fmt"
)

// DAGManager orchestrates execution of a DAG.
type DAGManager interface {
	// Load the DAG definition
	Load(dag DAGOps) error

	// Execute the DAG (with parallelism, joins, waits)
	Execute(ctx context.Context) error

	// Commit section ensures exclusive access.
	Commit(ctx context.Context) error

	// Print the DAG structure for debugging.
	Print(ctx context.Context)
}

// Load the DAG definition.
func (d *DAG) Load(dag DAGOps) error {
	if d.dag.isAcyclic() {
		d.dag = dag
		return nil
	}
	return fmt.Errorf("the provided DAG is not acyclic")
}

// Execute the DAG with parallelism, joins, and waits.
func (d *DAG) Execute(ctx context.Context) error {
	// Implementation of DAG execution logic goes here.
	// This would typically involve traversing the DAG, executing nodes,
	// handling parallel execution, and managing dependencies.
	return nil
}

// Commit section ensures exclusive access for the provided function.
func (d *DAG) Commit(ctx context.Context) error {
	// Implementation of commit logic goes here.
	// This would typically involve acquiring a lock or ensuring that
	// the function is executed in a thread-safe manner.
	return nil
}

func (d *DAG) Print(ctx context.Context) {
	// Implementation of printing the DAG structure for debugging.
	// This would typically involve iterating over the nodes and edges
	// and printing their details.
	nodes := d.dag.vertices()
	for _, node := range nodes {
		fmt.Printf("Node ID: %s\n", node.ID())
		deps := d.dag.dependencies(node.ID())
		if len(deps) > 0 {
			fmt.Printf("Dependencies: %v\n", deps)
		} else {
			fmt.Println("No dependencies")
		}
	}
}

// NewDagManager creates a new DAGManager instance.
func NewDagManager(dag DAGOps, pt Persist) DAGManager {
	return &DAG{
		dag:     dag,
		persist: pt,
	}
}
