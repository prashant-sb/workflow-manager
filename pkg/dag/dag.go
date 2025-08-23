package dag

import (
	"context"
	"fmt"

	"github.com/prashantsb/workflow-manager/pkg/preserver"
)

type Persist = preserver.Preserve[PresistAttrib]
type adjacencyList = map[string][]string
type VertexMap = map[string]Vertex

// Vertex represents a unit of execution in the DAG.
type Vertex interface {
	ID() string
	Run(ctx context.Context) error
	Group() string
}

// DAGOps represents a directed acyclic graph of nodes.
type DAGOps interface {
	AddVertex(v Vertex)
	AddEdge(from, to string) error
	vertices() []Vertex
	dependencies(nodeID string) []string
	isAcyclic() bool
}

type VertexState struct {
	Status string `json:"status"`
}

// PersistAttrib defines the structure for persisting DAG state.
type PresistAttrib struct {
	Vertices map[string]VertexState `json:"vertices"`
}

// Task is a concrete implementation of Vertex.
type Task struct {
	Id       string
	SubGraph string
	Fn       func(ctx context.Context) error
}

// DagAttributes implements DAGOps interface.
type DagAttributes struct {
	verts VertexMap
	edges adjacencyList
}

// DAG implements the DAGManager interface.
type DAG struct {
	persist Persist
	dag     DAGOps
}

func NewTask(id, grp string, fn func(ctx context.Context) error) Vertex {
	return &Task{Id: id, Fn: fn, SubGraph: grp}
}

func (t *Task) ID() string {
	return t.Id
}

func (t *Task) Run(ctx context.Context) error {
	return t.Fn(ctx)
}

func (t *Task) Group() string {
	return t.SubGraph
}

// NewDAG creates a new empty DAG.
func NewDAG() DAGOps {
	return &DagAttributes{
		verts: make(map[string]Vertex),
		edges: make(map[string][]string),
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
func (d *DagAttributes) isAcyclic() bool {
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
