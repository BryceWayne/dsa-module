package graphs

import (
    "fmt"
    "sort"
)

// DAG (Directed Acyclic Graph) structure with a map
type DAG struct {
    vertices      map[int][]int
    incomingEdges map[int]int // Count of incoming edges for each vertex
}

// NewDAG creates a new instance of a Directed Acyclic Graph.
func NewDAG() *DAG {
    return &DAG{
        vertices:      make(map[int][]int),
        incomingEdges: make(map[int]int),
    }
}

// AddEdge adds a directed edge from `u` to `v` to the DAG.
func (dag *DAG) AddEdge(u, v int) {
    if _, exists := dag.vertices[u]; !exists {
        dag.vertices[u] = []int{}
    }
    if _, exists := dag.vertices[v]; !exists {
        dag.vertices[v] = []int{}
    }

    dag.vertices[u] = append(dag.vertices[u], v)
    dag.incomingEdges[v]++
}

// PrettyPrint prints the DAG in a visually appealing format.
func (dag *DAG) PrettyPrint() {
    keys := make([]int, 0, len(dag.vertices))
    for k := range dag.vertices {
        keys = append(keys, k)
    }
    sort.Ints(keys) // Sort the keys for consistent order

    fmt.Println("DAG:")
    for _, k := range keys {
        fmt.Printf("%d -> ", k)
        for i, v := range dag.vertices[k] {
            if i > 0 {
                fmt.Print(", ")
            }
            fmt.Print(v)
        }
        fmt.Println()
    }
}
