package graphs

import (
    "fmt"
)

// Edge represents a weighted edge in the graph.
type Edge struct {
    Node   int
    Weight int
}

// WeightedGraph represents a graph with weighted edges.
type WeightedGraph struct {
    vertices map[int][]Edge
}

// NewWeightedGraph creates a new instance of a weighted graph.
func NewWeightedGraph() *WeightedGraph {
    return &WeightedGraph{
        vertices: make(map[int][]Edge),
    }
}

// AddEdge adds a weighted edge to the graph.
func (wg *WeightedGraph) AddEdge(u, v, weight int) {
    wg.vertices[u] = append(wg.vertices[u], Edge{Node: v, Weight: weight})
    // If the graph is undirected, add an edge in the reverse direction too
    // wg.vertices[v] = append(wg.vertices[v], Edge{Node: u, Weight: weight})
}

// PrettyPrint prints the graph in a visually appealing format.
func (wg *WeightedGraph) PrettyPrint() {
    for k, edges := range wg.vertices {
        fmt.Printf("%d -> ", k)
        for i, e := range edges {
            if i > 0 {
                fmt.Print(", ")
            }
            fmt.Printf("%d (weight: %d)", e.Node, e.Weight)
        }
        fmt.Println()
    }
}
