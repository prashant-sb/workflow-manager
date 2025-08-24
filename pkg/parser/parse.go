package parser

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/prashantsb/workflow-manager/pkg/dag"
	"github.com/prashantsb/workflow-manager/pkg/tasks"
)

type DOTParser interface {
	Parse(map[string]*tasks.Task) (dag.DAGOps, error)
	Validate(dag.DagAttributes) error
}

type dotParser struct {
	from *strings.Reader
}

func NewDOTParser(wf string) DOTParser {
	return &dotParser{from: strings.NewReader(wf)}
}

func (p *dotParser) Parse(taskRegistry map[string]*tasks.Task) (dag.DAGOps, error) {
	var currentGroup string
	dag := dag.NewDAG()
	scanner := bufio.NewScanner(p.from)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty or global lines
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "digraph") {
			continue
		}

		// Enter subgraph
		if strings.HasPrefix(line, "subgraph") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				currentGroup = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Exit subgraph
		if line == "}" {
			currentGroup = ""
			continue
		}

		// Example: Net -> VM;
		if strings.Contains(line, "->") {
			parts := strings.Split(line, "->")
			if len(parts) != 2 {
				continue
			}
			from := strings.Trim(strings.TrimSpace(parts[0]), ";")
			to := strings.Trim(strings.TrimSpace(parts[1]), ";")

			// Add or update vertex with group
			if t, ok := taskRegistry[from]; ok {
				t.WorkflowId = currentGroup
				dag.AddVertex(t)
			} else {
				return nil, fmt.Errorf("task %s not found in registry", from)
			}

			if t, ok := taskRegistry[to]; ok {
				t.WorkflowId = currentGroup
				dag.AddVertex(t)
			} else {
				return nil, fmt.Errorf("task %s not found in registry", to)
			}

			// Add edge
			if err := dag.AddEdge(from, to); err != nil {
				return nil, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dag, nil
}

func (p *dotParser) Validate(d dag.DagAttributes) error {
	// Implementation of validation logic goes here.
	// This would typically involve checking the DOT syntax and structure.
	if d.IsAcyclic() {
		return nil
	}
	return fmt.Errorf("the DAG contains cycles")
}
