package dag

import (
	"fmt"

	"github.com/prashantsb/workflow-manager/pkg/preserver"
	"github.com/prashantsb/workflow-manager/pkg/tasks"
)

type Persist = preserver.Preserve[PresistAttrib]
type Vertex = tasks.Vertex
type adjacencyList = map[string][]string
type vertexMap = map[string]Vertex

// DAGOps represents a directed acyclic graph of nodes.
type DAGOps interface {
	AddVertex(v Vertex)
	AddEdge(from, to string) error
	vertices() []Vertex
	dependencies(nodeID string) []string
	IsAcyclic() bool
}

type VertexState struct {
	Status string `json:"status"`
}

// PersistAttrib defines the structure for persisting DAG state.
type PresistAttrib struct {
	Vertices map[string]VertexState `json:"vertices"`
}

// DagAttributes implements DAGOps interface.
type DagAttributes struct {
	verts vertexMap
	edges adjacencyList
}

// DAG implements the DAGManager interface.
type DAG struct {
	persist Persist
	dag     DAGOps
}

// NewDAG creates a new empty DAG.
func NewDAG() DAGOps {
	return &DagAttributes{
		verts: make(vertexMap),
		edges: make(adjacencyList),
	}
}

func (d *DagAttributes) AddVertex(v Vertex) {
	d.verts[v.ID()] = v
	if _, ok := d.edges[v.ID()]; !ok {
		d.edges[v.ID()] = []string{}
	}
}

func (d *DagAttributes) AddEdge(from, to string) error {
	if _, ok := d.verts[from]; !ok {
		return fmt.Errorf("from-vertex %s not found", from)
	}
	if _, ok := d.verts[to]; !ok {
		return fmt.Errorf("to-vertex %s not found", to)
	}
	d.edges[from] = append(d.edges[from], to)
	return nil
}

func (d *DagAttributes) vertices() []Vertex {
	list := make([]Vertex, 0, len(d.verts))
	for _, v := range d.verts {
		list = append(list, v)
	}
	return list
}

func (d *DagAttributes) dependencies(nodeID string) []string {
	return d.edges[nodeID]
}

// cycle detection using DFS
func (d *DagAttributes) IsAcyclic() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var dfs func(string) bool
	dfs = func(node string) bool {
		if recStack[node] {
			return false // cycle found
		}
		if visited[node] {
			return true
		}
		visited[node] = true
		recStack[node] = true

		for _, dep := range d.edges[node] {
			if !dfs(dep) {
				return false
			}
		}
		recStack[node] = false
		return true
	}

	for id := range d.verts {
		if !visited[id] {
			if !dfs(id) {
				return false
			}
		}
	}
	return true
}
